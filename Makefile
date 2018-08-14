VERSION=0.37
GOCMD=`which go`
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
XGO=GOPATH=$(GOPATH) $(GOPATH)/bin/xgo -go 1.10.x --dest=build
BUILDVERSION=-ldflags="-X main.VERSION=$(VERSION)"
BINARY_NAME=statup
DOCKER=/usr/local/bin/docker
PATH:=/usr/local/bin:$(GOPATH)/bin:$(PATH)

all: deps compile install clean

release: deps build-all compress

install: build
	mv $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)
	$(GOPATH)/bin/$(BINARY_NAME) version

build:
	$(GOBUILD) -ldflags="-X main.VERSION=$(VERSION)" -o $(BINARY_NAME) -v ./cmd

run: build
	./$(BINARY_NAME)

compile:
	cd source && $(GOPATH)/bin/rice embed-go
	$(GOPATH)/bin/wt compile source/scss/base.scss -b source/css

test: compile test-env
	$(GOTEST) ./... -p 1 -ldflags="-X main.VERSION=$(VERSION)" -coverprofile=coverage.out -v

test-all: compile test-env databases
	$(GOTEST) ./... -p 1 -ldflags="-X main.VERSION=$(VERSION)" -coverprofile=coverage.out -v

coverage:
	$(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=travis -repotoken $(COVERALLS)

build-all: clean compile
	mkdir build
	$(XGO) $(BUILDVERSION) --targets=darwin/amd64 ./cmd
	$(XGO) $(BUILDVERSION) --targets=darwin/386 ./cmd
	$(XGO) $(BUILDVERSION) --targets=linux/amd64 ./cmd
	$(XGO) $(BUILDVERSION) --targets=linux/386 ./cmd
	$(XGO) $(BUILDVERSION) --targets=windows-6.0/amd64 ./cmd
	$(XGO) $(BUILDVERSION) --targets=linux/arm-7 ./cmd
	$(XGO) $(BUILDVERSION) --targets=linux/arm64 ./cmd
	$(XGO) --targets=linux/amd64 -ldflags="-X main.VERSION=$VERSION -linkmode external -extldflags -static" -out alpine ./cmd

docker:
	$(DOCKER) build -t hunterlong/statup:latest .

docker-dev:
	$(DOCKER) build -t hunterlong/statup:dev -f Dockerfile-dev .

docker-run: docker
	$(DOCKER) run -t -p 8080:8080 hunterlong/statup:latest

docker-dev-run: docker-dev
	$(DOCKER) run -t -p 8080:8080 hunterlong/statup:dev

databases:
	$(DOCKER) run --name statup_postgres -p 5432:5432 -e POSTGRES_PASSWORD=password123 -e POSTGRES_USER=root -e POSTGRES_DB=root -d postgres
	$(DOCKER) run --name statup_mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password123 -e MYSQL_DATABASE=root -d mysql
	sleep 30

deps:
	$(GOGET) github.com/wellington/wellington/wt
	$(GOGET) github.com/stretchr/testify/assert
	$(GOGET) golang.org/x/tools/cmd/cover
	$(GOGET) github.com/mattn/goveralls
	$(GOINSTALL) github.com/mattn/goveralls
	$(GOGET) github.com/rendon/testcli
	$(GOGET) github.com/karalabe/xgo
	$(GOGET) github.com/GeertJohan/go.rice
	$(GOGET) github.com/GeertJohan/go.rice/rice

clean:
	rm -rf build
	rm -f statup
	rm -rf logs
	rm -rf cmd/logs
	rm -rf cmd/plugins
	rm -rf cmd/statup.db
	rm -rf cmd/config.yml
	rm -rf cmd/.sass-cache
	rm -rf core/logs
	rm -rf core/.sass-cache
	rm -rf core/config.yml
	rm -f core/statup.db
	rm -rf handlers/config.yml
	rm -rf handlers/statup.db
	rm -rf source/logs
	rm -rf utils/logs
	rm -rf .sass-cache
	rm -f coverage.out

tag:
	git tag "v$(VERSION)" --force

test-env:
	export GO_ENV=test
	export DB_HOST=localhost
	export DB_USER=root
	export DB_PASS=password123
	export DB_DATABASE=root
	export NAME=Demo
	export CMD_FILE=$(GOPATH)/src/github.com/hunterlong/statup/cmd.sh
	export STATUP_DIR=$(GOPATH)/src/github.com/hunterlong/statup

compress:
	mv build/alpine-linux-amd64 build/$(BINARY_NAME)
	tar -czvf build/$(BINARY_NAME)-linux-alpine.tar.gz build/$(BINARY_NAME) && rm -f build/$(BINARY_NAME)
	mv build/cmd-darwin-10.6-amd64 build/$(BINARY_NAME)
	tar -czvf build/$(BINARY_NAME)-osx-x64.tar.gz build/$(BINARY_NAME) && rm -f build/$(BINARY_NAME)
	mv build/cmd-darwin-10.6-386 build/$(BINARY_NAME)
	tar -czvf build/$(BINARY_NAME)-osx-x32.tar.gz build/$(BINARY_NAME) && rm -f build/$(BINARY_NAME)
	mv build/cmd-linux-amd64 build/$(BINARY_NAME)
	tar -czvf build/$(BINARY_NAME)-linux-x64.tar.gz build/$(BINARY_NAME) && rm -f build/$(BINARY_NAME)
	mv build/cmd-linux-386 build/$(BINARY_NAME)
	tar -czvf build/$(BINARY_NAME)-linux-x32.tar.gz build/$(BINARY_NAME) && rm -f build/$(BINARY_NAME)
	mv build/cmd-windows-6.0-amd64.exe build/$(BINARY_NAME).exe
	zip build/$(BINARY_NAME)-windows-x64.zip build/$(BINARY_NAME).exe  && rm -f build/$(BINARY_NAME).exe
	mv build/cmd-linux-arm-7 build/$(BINARY_NAME)
	tar -czvf build/$(BINARY_NAME)-linux-arm7.tar.gz build/$(BINARY_NAME) && rm -f build/$(BINARY_NAME)
	mv build/cmd-linux-arm64 build/$(BINARY_NAME)
	tar -czvf build/$(BINARY_NAME)-linux-arm64.tar.gz build/$(BINARY_NAME) && rm -f build/$(BINARY_NAME)
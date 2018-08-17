VERSION=0.44
BINARY_NAME=statup
GOPATH:=$(GOPATH)
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
XGO=GOPATH=$(GOPATH) $(GOPATH)/bin/xgo -go 1.10.x --dest=build
BUILDVERSION=-ldflags "-X main.VERSION=$(VERSION) -X main.COMMIT=$(TRAVIS_COMMIT)"
RICE=$(GOPATH)/bin/rice
PATH:=/usr/local/bin:$(GOPATH)/bin:$(PATH)
PUBLISH_BODY='{ "request": { "branch": "master", "config": { "env": { "VERSION": "$(VERSION)", "COMMIT": "$(TRAVIS_COMMIT)" } } } }'
DOCKER_TEST='{ "request": { "branch": "master", "config": { "script": "make docker-run-test", "services": ["docker"], "before_script": [], "after_deploy": [], "after_success": ["make publish-dev", "sleep 240", "make travis-crypress"], "deploy": [], "before_deploy": [], "env": { "VERSION": "$(VERSION)" } } } }'
TEST_DIR=$(GOPATH)/src/github.com/hunterlong/statup

all: deps compile install clean

release: deps build-all compress

install: clean build
	mv $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)
	$(GOPATH)/bin/$(BINARY_NAME) version

build: compile
	$(GOBUILD) $(BUILDVERSION) -o $(BINARY_NAME) -v ./cmd

run: build
	./$(BINARY_NAME) --ip localhost --port 8080

compile:
	cd source && $(GOPATH)/bin/rice embed-go
	sass source/scss/base.scss source/css/base.css

test: clean compile install
	STATUP_DIR=$(TEST_DIR) GO_ENV=test go test -v -p=1 $(BUILDVERSION) -coverprofile=coverage.out ./...
	gocov convert coverage.out > coverage.json

test-all: compile databases
	$(GOTEST) ./... -p=1 $(BUILDVERSION) -coverprofile=coverage.out -v

coverage:
	$(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=travis -repotoken $(COVERALLS)

docs:
	godoc2md github.com/hunterlong/statup > servers/docs/README.md
	gocov-html coverage.json > servers/docs/COVERAGE.html
	revive -formatter stylish > servers/docs/LINT.md

build-all: clean compile
	mkdir build
	$(XGO) $(BUILDVERSION) --targets=darwin/amd64 ./cmd
	$(XGO) $(BUILDVERSION) --targets=darwin/386 ./cmd
	$(XGO) $(BUILDVERSION) --targets=linux/amd64 ./cmd
	$(XGO) $(BUILDVERSION) --targets=linux/386 ./cmd
	$(XGO) $(BUILDVERSION) --targets=windows-6.0/amd64 ./cmd
	$(XGO) $(BUILDVERSION) --targets=linux/arm-7 ./cmd
	$(XGO) $(BUILDVERSION) --targets=linux/arm64 ./cmd
	$(XGO) --targets=linux/amd64 -ldflags="-X main.VERSION=$(VERSION) -X main.COMMIT=$(TRAVIS_COMMIT) -linkmode external -extldflags -static" -out alpine ./cmd

build-alpine: clean compile
	mkdir build
	$(XGO) --targets=linux/amd64 -ldflags="-X main.VERSION=$(VERSION) -X main.COMMIT=$(TRAVIS_COMMIT) -linkmode external -extldflags -static" -out alpine ./cmd

docker:
	docker build -t hunterlong/statup:latest .

docker-dev:
	docker build -t hunterlong/statup:dev -f ./cmd/Dockerfile .

docker-test:
	docker build -t hunterlong/statup:test -f ./.travis/Dockerfile .

docker-run: docker
	docker run -t -p 8080:8080 hunterlong/statup:latest

docker-run-dev: docker-dev
	docker run -t -p 8080:8080 hunterlong/statup:dev

docker-run-test: docker-test
	docker run -t -p 8080:8080 hunterlong/statup:test

databases:
	docker run --name statup_postgres -p 5432:5432 -e POSTGRES_PASSWORD=password123 -e POSTGRES_USER=root -e POSTGRES_DB=root -d postgres
	docker run --name statup_mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password123 -e MYSQL_DATABASE=root -d mysql
	sleep 30

deps:
	$(GOGET) github.com/stretchr/testify/assert
	$(GOGET) golang.org/x/tools/cmd/cover
	$(GOGET) github.com/mattn/goveralls
	$(GOINSTALL) github.com/mattn/goveralls
	$(GOGET) github.com/rendon/testcli
	$(GOGET) github.com/karalabe/xgo
	$(GOGET) github.com/GeertJohan/go.rice
	$(GOGET) github.com/GeertJohan/go.rice/rice
	$(GOINSTALL) github.com/GeertJohan/go.rice/rice
	$(GOCMD) get github.com/davecheney/godoc2md
	$(GOCMD) install github.com/davecheney/godoc2md
	$(GOCMD) get github.com/axw/gocov/gocov
	$(GOCMD) get -u gopkg.in/matm/v1/gocov-html
	$(GOCMD) install gopkg.in/matm/v1/gocov-html
	$(GOCMD) get -u github.com/mgechev/revive
	$(GOGET) -d ./...

clean:
	rm -rf ./{logs,assets,plugins,statup.db,config.yml,.sass-cache,config.yml,statup,build}
	rm -rf cmd/{logs,assets,plugins,statup.db,config.yml,.sass-cache}
	rm -rf core/{logs,assets,plugins,statup.db,config.yml,.sass-cache}
	rm -rf handlers/{logs,assets,plugins,statup.db,config.yml,.sass-cache}
	rm -rf notifiers/{logs,assets,plugins,statup.db,config.yml,.sass-cache}
	rm -rf source/{logs,assets,plugins,statup.db,config.yml,.sass-cache}
	rm -rf types/{logs,assets,plugins,statup.db,config.yml,.sass-cache}
	rm -rf utils/{logs,assets,plugins,statup.db,config.yml,.sass-cache}
	rm -rf .sass-cache
	rm -f coverage.out
	rm -f coverage.json

tag:
	git tag "v$(VERSION)" --force

test-env:
	export GO_ENV=test
	export DB_HOST=localhost
	export DB_USER=root
	export DB_PASS=password123
	export DB_DATABASE=root
	export NAME=Demo
	export STATUP_DIR=$(GOPATH)/src/github.com/hunterlong/statup

compress:
	cd build && mv alpine-linux-amd64 $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-linux-alpine.tar.gz $(BINARY_NAME) && rm -f $(BINARY_NAME)
	cd build && mv cmd-darwin-10.6-amd64 $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-osx-x64.tar.gz $(BINARY_NAME) && rm -f $(BINARY_NAME)
	cd build && mv cmd-darwin-10.6-386 $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-osx-x32.tar.gz $(BINARY_NAME) && rm -f $(BINARY_NAME)
	cd build && mv cmd-linux-amd64 $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-linux-x64.tar.gz $(BINARY_NAME) && rm -f $(BINARY_NAME)
	cd build && mv cmd-linux-386 $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-linux-x32.tar.gz $(BINARY_NAME) && rm -f $(BINARY_NAME)
	cd build && mv cmd-windows-6.0-amd64.exe $(BINARY_NAME).exe
	cd build && zip $(BINARY_NAME)-windows-x64.zip $(BINARY_NAME).exe  && rm -f $(BINARY_NAME).exe
	cd build && mv cmd-linux-arm-7 $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-linux-arm7.tar.gz $(BINARY_NAME) && rm -f $(BINARY_NAME)
	cd build && mv cmd-linux-arm64 $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-linux-arm64.tar.gz $(BINARY_NAME) && rm -f $(BINARY_NAME)

publish-dev:
	curl -H "Content-Type: application/json" --data '{"docker_tag": "dev"}' -X POST $(DOCKER)

publish-test:
	curl -H "Content-Type: application/json" --data '{"docker_tag": "test"}' -X POST $(DOCKER)

publish-latest:
	curl -H "Content-Type: application/json" --data '{"docker_tag": "latest"}' -X POST $(DOCKER)

publish-homebrew:
	curl -s -X POST -H "Content-Type: application/json" -H "Accept: application/json" -H "Travis-API-Version: 3" -H "Authorization: token $(TRAVIS_API)" -d $(PUBLISH_BODY) https://api.travis-ci.com/repo/hunterlong%2Fhomebrew-statup/requests

travis-crypress:
	curl -s -X POST -H "Content-Type: application/json" -H "Accept: application/json" -H "Travis-API-Version: 3" -H "Authorization: token $(TRAVIS_API)" -d $(PUBLISH_BODY) https://api.travis-ci.com/repo/hunterlong%2Fstatup-testing/requests

travis-docker-test:
	curl -s -X POST -H "Content-Type: application/json" -H "Accept: application/json" -H "Travis-API-Version: 3" -H "Authorization: token $(TRAVIS_API)" -d $(DOCKER_TEST) https://api.travis-ci.com/repo/hunterlong%2Fstatup/requests

.PHONY: build build-all build-alpine
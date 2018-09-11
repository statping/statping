VERSION=0.57
BINARY_NAME=statup
GOPATH:=$(GOPATH)
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
XGO=GOPATH=$(GOPATH) xgo -go 1.10.x --dest=build
BUILDVERSION=-ldflags "-X main.VERSION=$(VERSION) -X main.COMMIT=$(TRAVIS_COMMIT)"
RICE=$(GOPATH)/bin/rice
PATH:=/usr/local/bin:$(GOPATH)/bin:$(PATH)
PUBLISH_BODY='{ "request": { "branch": "master", "config": { "env": { "VERSION": "$(VERSION)", "COMMIT": "$(TRAVIS_COMMIT)" } } } }'
TRAVIS_BUILD_CMD='{ "request": { "branch": "master", "message": "Compile master for Statup v$(VERSION)", "config": { "os": [ "linux" ], "language": "go", "go": [ "1.10.x" ], "go_import_path": "github.com/hunterlong/statup", "install": true, "sudo": "required", "services": [ "docker" ], "env": { "VERSION": "$(VERSION)" }, "matrix": { "allow_failures": [ { "go": "master" } ], "fast_finish": true }, "before_deploy": [ "git config --local user.name \"hunterlong\"", "git config --local user.email \"info@socialeck.com\"", "make tag" ], "deploy": [ { "provider": "releases", "api_key": "$(GH_TOKEN)", "file": [ "build/statup-osx-x64.tar.gz", "build/statup-osx-x32.tar.gz", "build/statup-linux-x64.tar.gz", "build/statup-linux-x32.tar.gz", "build/statup-linux-arm64.tar.gz", "build/statup-linux-arm7.tar.gz", "build/statup-linux-alpine.tar.gz", "build/statup-windows-x64.zip" ], "skip_cleanup": true } ], "notifications": { "email": false }, "before_script": ["gem install sass"], "script": [ "travis_wait 30 docker pull karalabe/xgo-latest", "make release" ], "after_success": [], "after_deploy": [ "make publish-dev" ] } } }'
TEST_DIR=$(GOPATH)/src/github.com/hunterlong/statup

# build and compile Statup and then test
all: dev-deps compile install test-all docker-build-all

# build all arch's and release Statup
release: dev-deps build-all

# test all versions of Statup, golang testing and then cypress UI testing
test-all: dev-deps test

# test all versions of Statup, golang testing and then cypress UI testing
test-ui: dev-deps docker-build-dev cypress-test

# testing to be ran on travis ci
travis-test: dev-deps cypress-install test coverage

# build and compile all arch's for Statup
build-all: build-mac build-linux build-windows build-alpine compress

# build all docker tags
docker-build-all: docker-build-base docker-build-dev docker-build-latest

# push all docker tags built
docker-publish-all: docker-push-base docker-push-dev docker-push-latest

# build Statup for local arch
build: compile
	$(GOBUILD) $(BUILDVERSION) -o $(BINARY_NAME) -v ./cmd

# build Statup plugins
build-plugin:
	$(GOBUILD) $(BUILDVERSION) -buildmode=plugin -o $(BINARY_NAME) -v ./dev/plugin

# build Statup debug app
build-debug: compile
	$(GOBUILD) $(BUILDVERSION) -tags debug -o $(BINARY_NAME) -v ./cmd

# install Statup for local arch and move binary to gopath/src/bin/statup
install: build
	mv $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)
	$(GOPATH)/bin/$(BINARY_NAME) version

# run Statup from local arch
run: build
	./$(BINARY_NAME) --ip 0.0.0.0 --port 8080

# compile assets using SASS and Rice. compiles scss -> css, and run rice embed-go
compile:
	cd source && $(GOPATH)/bin/rice embed-go
	sass source/scss/base.scss source/css/base.css
	rm -rf .sass-cache

# benchmark testing
benchmark:
	cd handlers && go test -v -run=^$ -bench=. -benchtime=5s -memprofile=prof.mem -cpuprofile=prof.cpu

# view benchmark testing using pprof
benchmark-view:
	go tool pprof handlers/handlers.test handlers/prof.cpu > top20

# test Statup golang tetsing files
test: clean compile install
	STATUP_DIR=$(TEST_DIR) go test -v -p=1 $(BUILDVERSION) -coverprofile=coverage.out ./...
	gocov convert coverage.out > coverage.json

# report coverage to Coveralls
coverage:
	$(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=travis -repotoken $(COVERALLS)

# generate documentation for Statup functions
docs:
	godoc2md github.com/hunterlong/statup > servers/docs/README.md
	gocov-html coverage.json > servers/docs/COVERAGE.html
	revive -formatter stylish > servers/docs/LINT.md

#
#    Build binary for Statup
#

# build Statup for Mac, 64 and 32 bit
build-mac: compile
	mkdir build
	$(XGO) $(BUILDVERSION) --targets=darwin/amd64 ./cmd
	$(XGO) $(BUILDVERSION) --targets=darwin/386 ./cmd

# build Statup for Linux 64, 32 bit, arm6/arm7
build-linux: compile
	$(XGO) $(BUILDVERSION) --targets=linux/amd64 ./cmd
	$(XGO) $(BUILDVERSION) --targets=linux/386 ./cmd
	$(XGO) $(BUILDVERSION) --targets=linux/arm-7 ./cmd
	$(XGO) $(BUILDVERSION) --targets=linux/arm64 ./cmd

# build for windows 64 bit only
build-windows: compile
	$(XGO) $(BUILDVERSION) --targets=windows-6.0/amd64 ./cmd

# build Alpine linux binary (used in docker images)
build-alpine: compile
	$(XGO) --targets=linux/amd64 -ldflags="-X main.VERSION=$(VERSION) -X main.COMMIT=$(TRAVIS_COMMIT) -linkmode external -extldflags -static" -out alpine ./cmd

#
#    Docker Makefile commands
#

# build :latest docker tag
docker-build-latest: docker-build-base
	docker build -t hunterlong/statup:latest --no-cache -f Dockerfile .

# build :dev docker tag
docker-build-dev: docker-build-base
	docker build -t hunterlong/statup:dev --no-cache -f dev/Dockerfile-dev .

# build :base and base-v{VERSION} docker tag
docker-build-base: clean
	wget -q https://assets.statup.io/sass && chmod +x sass
	$(XGO) --targets=linux/amd64 -ldflags="-X main.VERSION=$(VERSION) -linkmode external -extldflags -static" -out alpine ./cmd
	docker build -t hunterlong/statup:base --no-cache -f dev/Dockerfile-base .
	docker tag hunterlong/statup:base hunterlong/statup:base-v$(VERSION)

# build Cypress UI testing :cypress docker tag
docker-build-cypress: clean
	GOPATH=$(GOPATH) xgo -out statup -go 1.10.x -ldflags "-X main.VERSION=$(VERSION) -X main.COMMIT=$(TRAVIS_COMMIT)" --targets=linux/amd64 ./cmd
	docker build -t hunterlong/statup:cypress -f dev/Dockerfile-cypress .
	rm -f statup

# run hunterlong/statup:latest docker image
docker-run: docker-build-latest
	docker run -it -p 8080:8080 hunterlong/statup:latest

# run hunterlong/statup:dev docker image
docker-run-dev: docker-build-dev
	docker run -t -p 8080:8080 hunterlong/statup:dev

# run Cypress UI testing, hunterlong/statup:cypress docker image
docker-run-cypress: docker-build-cypress
	docker run -t hunterlong/statup:cypress

# push the :dev tag to Docker hub
docker-push-dev:
	docker push hunterlong/statup:dev

# push the :cypress tag to Docker hub
docker-push-cypress:
	docker push hunterlong/statup:cypress

# push the :latest tag to Docker hub
docker-push-latest: docker
	docker push hunterlong/statup:latest

# push the :base and :base-v{VERSION} tag to Docker hub
docker-push-base: docker-build-base
	docker tag hunterlong/statup:base hunterlong/statup:base-v$(VERSION)
	docker push hunterlong/statup:base
	docker push hunterlong/statup:base-v$(VERSION)

# create Postgres, and MySQL instance using Docker (used for testing)
databases:
	docker run --name statup_postgres -p 5432:5432 -e POSTGRES_PASSWORD=password123 -e POSTGRES_USER=root -e POSTGRES_DB=root -d postgres
	docker run --name statup_mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password123 -e MYSQL_DATABASE=root -d mysql
	sleep 30


#
#    Download and Install dependencies
#

# run dep to install all required golang dependecies
dep:
	dep ensure -vendor-only

# install all required golang dependecies
dev-deps: dep
	$(GOGET) -u github.com/jinzhu/gorm/...
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
	$(GOCMD) get gopkg.in/matm/v1/gocov-html
	$(GOCMD) install gopkg.in/matm/v1/gocov-html
	$(GOCMD) get github.com/mgechev/revive

# remove files for a clean compile/build
clean:
	rm -rf ./{logs,assets,plugins,statup.db,config.yml,.sass-cache,config.yml,statup,build,.sass-cache,statup.db,index.html,vendor}
	rm -rf cmd/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf core/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf handlers/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf notifiers/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf source/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf types/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf utils/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf dev/test/cypress/videos
	rm -f coverage.* sass
	find . -name "*.out" -type f -delete
	find . -name "*.cpu" -type f -delete
	find . -name "*.mem" -type f -delete
	find . -name "*.test" -type f -delete

# tag version using git
tag:
	git tag "v$(VERSION)" --force

# compress built binaries into tar.gz and zip formats
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

# push the :dev docker tag using curl
publish-dev:
	curl -H "Content-Type: application/json" --data '{"docker_tag": "dev"}' -X POST $(DOCKER)

# push the :latest docker tag using curl
publish-latest:
	curl -H "Content-Type: application/json" --data '{"docker_tag": "latest"}' -X POST $(DOCKER)

# update the homebrew application to latest for mac
publish-homebrew:
	curl -s -X POST -H "Content-Type: application/json" -H "Accept: application/json" -H "Travis-API-Version: 3" -H "Authorization: token $(TRAVIS_API)" -d $(PUBLISH_BODY) https://api.travis-ci.com/repo/hunterlong%2Fhomebrew-statup/requests

# install NPM reuqirements for cypress testing
cypress-install:
	cd dev/test && npm install

# run Cypress UI testing
cypress-test: clean cypress-install
	cd dev/test && npm test

# build Statup using a travis ci trigger
travis-build:
	curl -s -X POST -H "Content-Type: application/json" -H "Accept: application/json" -H "Travis-API-Version: 3" -H "Authorization: token $(TRAVIS_API)" -d $(TRAVIS_BUILD_CMD) https://api.travis-ci.com/repo/hunterlong%2Fstatup/requests

# install xgo and pull the xgo docker image
xgo-install: clean
	go get github.com/karalabe/xgo
	docker pull karalabe/xgo-latest

.PHONY: all build build-all build-alpine test-all test

VERSION=$(shell cat version.txt)
SIGN_KEY=B76D61FAA6DB759466E83D9964B9C6AAE2D55278
BINARY_NAME=statping
GOPATH:=$(GOPATH)
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
XGO=GOPATH=$(GOPATH) xgo -go 1.11 --dest=build
BUILDVERSION=-ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=$(TRAVIS_COMMIT)"
RICE=$(GOPATH)/bin/rice
PATH:=/usr/local/bin:$(GOPATH)/bin:$(PATH)
PUBLISH_BODY='{ "request": { "branch": "master", "message": "Homebrew update version v${VERSION}", "config": { "env": { "VERSION": "${VERSION}", "COMMIT": "$(TRAVIS_COMMIT)" } } } }'
TRAVIS_BUILD_CMD='{ "request": { "branch": "master", "message": "Compile master for Statping v${VERSION}", "config": { "os": [ "linux" ], "language": "go", "go": [ "1.10.x" ], "go_import_path": "github.com/hunterlong/statping", "install": true, "sudo": "required", "services": [ "docker" ], "env": { "VERSION": "${VERSION}" }, "matrix": { "allow_failures": [ { "go": "master" } ], "fast_finish": true }, "before_deploy": [ "git config --local user.name \"hunterlong\"", "git config --local user.email \"info@socialeck.com\"", "git tag v$(VERSION) --force"], "deploy": [ { "provider": "releases", "api_key": "$(GH_TOKEN)", "file_glob": true, "file": "build/*", "skip_cleanup": true } ], "notifications": { "email": false }, "before_script": ["gem install sass"], "script": [ "wget -O statping.gpg $(SIGN_URL)", "gpg --import statping.gpg", "travis_wait 30 docker pull karalabe/xgo-latest", "make release" ], "after_success": [], "after_deploy": [ "make publish-homebrew" ] } } }'
TEST_DIR=$(GOPATH)/src/github.com/hunterlong/statping
PATH:=$(PATH)

# build all arch's and release Statping
release: dev-deps build-all

# build and push the images to docker hub
docker: docker-build-all docker-publish-all

# test all versions of Statping, golang testing and then cypress UI testing
test-all: dev-deps test

# test all versions of Statping, golang testing and then cypress UI testing
test-ui: dev-deps docker-build-dev cypress-test

# testing to be ran on travis ci
travis-test: dev-deps cypress-install test coverage

# build and compile all arch's for Statping
build-all: build-mac build-linux build-windows build-alpine compress

# build all docker tags
docker-build-all: docker-build-latest

# push all docker tags built
docker-publish-all: docker-push-latest

snapcraft: clean snapcraft-build snapcraft-release

# build Statping for local arch
build: compile
	$(GOBUILD) $(BUILDVERSION) -o $(BINARY_NAME) -v ./cmd

# build Statping plugins
build-plugin:
	$(GOBUILD) $(BUILDVERSION) -buildmode=plugin -o ./dev/plugin/example.so -v ./dev/plugin

test-plugin: clean
	mkdir plugins
	$(GOBUILD) $(BUILDVERSION) -buildmode=plugin -o ./dev/plugin/example.so -v ./dev/plugin
	mv ./dev/plugin/example.so ./plugins/example.so
	STATPING_DIR=$(TEST_DIR) go test -v -p=1 $(BUILDVERSION) -coverprofile=coverage.out ./plugin

# build Statping debug app
build-debug: compile
	$(GOBUILD) $(BUILDVERSION) -tags debug -o $(BINARY_NAME) -v ./cmd

# install Statping for local arch and move binary to gopath/src/bin/statping
install: build
	mv $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)
	$(GOPATH)/bin/$(BINARY_NAME) version

# run Statping from local arch
run: build
	./$(BINARY_NAME) --ip 0.0.0.0 --port 8080

# compile assets using SASS and Rice. compiles scss -> css, and run rice embed-go
compile: generate
	sass source/scss/base.scss source/css/base.css
	cd source && $(GOPATH)/bin/rice embed-go
	rm -rf .sass-cache

# benchmark testing
benchmark:
	cd handlers && go test -v -run=^$ -bench=. -benchtime=5s -memprofile=prof.mem -cpuprofile=prof.cpu

# view benchmark testing using pprof
benchmark-view:
	go tool pprof handlers/handlers.test handlers/prof.cpu > top20

# test Statping golang tetsing files
test: clean compile install build-plugin
	STATPING_DIR=$(TEST_DIR) go test -v -p=1 $(BUILDVERSION) -coverprofile=coverage.out ./...
	gocov convert coverage.out > coverage.json

test-api:
	DB_CONN=sqlite DB_HOST=localhost DB_DATABASE=sqlite DB_PASS=none DB_USER=none statping &
	sleep 300 && newman run source/tmpl/postman.json -e dev/postman_environment.json --delay-request 500

# report coverage to Coveralls
coverage:
	$(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=travis -repotoken $(COVERALLS)

# generate documentation for Statping functions
docs:
	rm -f dev/README.md
	printf "# Statping Dev Documentation\n" > dev/README.md
	printf "This readme is automatically generated from the Golang documentation. [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/hunterlong/statping)\n\n" > dev/README.md
	godocdown github.com/hunterlong/statping >> dev/README.md
	godocdown github.com/hunterlong/statping/cmd >> dev/README.md
	godocdown github.com/hunterlong/statping/core >> dev/README.md
	godocdown github.com/hunterlong/statping/handlers >> dev/README.md
	godocdown github.com/hunterlong/statping/notifiers >> dev/README.md
	godocdown github.com/hunterlong/statping/plugin >> dev/README.md
	godocdown github.com/hunterlong/statping/source >> dev/README.md
	godocdown github.com/hunterlong/statping/types >> dev/README.md
	godocdown github.com/hunterlong/statping/utils >> dev/README.md
	gocov-html coverage.json > dev/COVERAGE.html
	revive -formatter stylish > dev/LINT.md

#
#    Build binary for Statping
#

# build Statping for Mac, 64 and 32 bit
build-mac: compile
	mkdir build
	$(XGO) $(BUILDVERSION) --targets=darwin/amd64,darwin/386 ./cmd

# build Statping for Linux 64, 32 bit, arm6/arm7
build-linux: compile
	$(XGO) $(BUILDVERSION) --targets=linux/amd64,linux/386,linux/arm-7,linux/arm-6,linux/arm64 ./cmd

# build for windows 64 bit only
build-windows: compile
	$(XGO) $(BUILDVERSION) --targets=windows-6.0/amd64 ./cmd

# build Alpine linux binary (used in docker images)
build-alpine: compile
	$(XGO) --targets=linux/amd64 -ldflags="-X main.VERSION=${VERSION} -X main.COMMIT=$(TRAVIS_COMMIT) -linkmode external -extldflags -static" -out alpine ./cmd

#
#    Docker Makefile commands
#

docker-test:
	docker-compose -f docker-compose.test.yml -p statping build
	docker-compose -f docker-compose.test.yml -p statping up -d
	docker logs -f statping_sut_1
	docker wait statping_sut_1

# build :latest docker tag
docker-build-latest:
	docker build --build-arg VERSION=${VERSION} -t hunterlong/statping:latest --no-cache -f Dockerfile .
	docker tag hunterlong/statping:latest hunterlong/statping:v${VERSION}

# build :dev docker tag
docker-build-dev:
	docker build --build-arg VERSION=${VERSION} -t hunterlong/statping:latest --no-cache -f Dockerfile .
	docker tag hunterlong/statping:dev hunterlong/statping:dev-v${VERSION}

# build Cypress UI testing :cypress docker tag
docker-build-cypress: clean
	GOPATH=$(GOPATH) xgo -out statping -go 1.10.x -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=$(TRAVIS_COMMIT)" --targets=linux/amd64 ./cmd
	docker build -t hunterlong/statping:cypress -f dev/Dockerfile-cypress .
	rm -f statping

# run hunterlong/statping:latest docker image
docker-run: docker-build-latest
	docker run -it -p 8080:8080 hunterlong/statping:latest

# run hunterlong/statping:dev docker image
docker-run-dev: docker-build-dev
	docker run -t -p 8080:8080 hunterlong/statping:dev

# run Cypress UI testing, hunterlong/statping:cypress docker image
docker-run-cypress: docker-build-cypress
	docker run -t hunterlong/statping:cypress

# push the :base and :base-v{VERSION} tag to Docker hub
docker-push-base:
	docker tag hunterlong/statping:base hunterlong/statping:base-v${VERSION}
	docker push hunterlong/statping:base
	docker push hunterlong/statping:base-v${VERSION}

# push the :dev tag to Docker hub
docker-push-dev:
	docker push hunterlong/statping:dev
	docker push hunterlong/statping:dev-v${VERSION}

# push the :cypress tag to Docker hub
docker-push-cypress:
	docker push hunterlong/statping:cypress

# push the :latest tag to Docker hub
docker-push-latest:
	docker tag hunterlong/statping hunterlong/statping:dev
	docker push hunterlong/statping:latest
	docker push hunterlong/statping:dev
	docker push hunterlong/statping:v${VERSION}

docker-run-mssql:
	docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=PaSsW0rD123' -p 1433:1433 -d microsoft/mssql-server-linux

# create Postgres, and MySQL instance using Docker (used for testing)
databases:
	docker run --name statping_postgres -p 5432:5432 -e POSTGRES_PASSWORD=password123 -e POSTGRES_USER=root -e POSTGRES_DB=root -d postgres
	docker run --name statping_mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password123 -e MYSQL_DATABASE=root -d mysql
	sleep 30


#
#    Download and Install dependencies
#
# run dep to install all required golang dependecies
dep:
	dep ensure -vendor-only

# install all required golang dependecies
dev-deps:
	$(GOGET) github.com/stretchr/testify/assert
	$(GOGET) golang.org/x/tools/cmd/cover
	$(GOGET) github.com/mattn/goveralls
	$(GOINSTALL) github.com/mattn/goveralls
	$(GOGET) github.com/rendon/testcli
	$(GOGET) github.com/robertkrimen/godocdown/godocdown
	$(GOGET) github.com/karalabe/xgo
	$(GOGET) github.com/GeertJohan/go.rice
	$(GOGET) github.com/GeertJohan/go.rice/rice
	$(GOINSTALL) github.com/GeertJohan/go.rice/rice
	$(GOCMD) get github.com/axw/gocov/gocov
	$(GOCMD) get gopkg.in/matm/v1/gocov-html
	$(GOCMD) install gopkg.in/matm/v1/gocov-html
	$(GOCMD) get github.com/mgechev/revive
	$(GOCMD) get github.com/fatih/structs
	$(GOGET) github.com/ararog/timeago
	$(GOGET) gopkg.in/natefinch/lumberjack.v2
	$(GOGET) golang.org/x/crypto/bcrypt

# remove files for a clean compile/build
clean:
	rm -rf ./{logs,assets,plugins,statup.db,config.yml,.sass-cache,config.yml,statping,build,.sass-cache,statup.db,index.html,vendor}
	rm -rf cmd/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log,*.html,*.json}
	rm -rf core/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf handlers/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf notifiers/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf source/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf types/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf utils/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log}
	rm -rf dev/{logs,assets,plugins,statup.db,config.yml,.sass-cache,*.log,test/app,plugin/*.so}
	rm -rf {parts,prime,snap,stage}
	rm -rf dev/test/cypress/videos
	rm -f coverage.* sass
	rm -f source/rice-box.go
	rm -f *.db-journal
	rm -rf *.snap
	find . -name "*.out" -type f -delete
	find . -name "*.cpu" -type f -delete
	find . -name "*.mem" -type f -delete
	rm -rf build

# tag version using git
tag:
	git tag v${VERSION} --force

generate:
	cd source && go generate

# compress built binaries into tar.gz and zip formats
compress:
	cd build && mv alpine-linux-amd64 $(BINARY_NAME)
	cd build && gpg --default-key $(SIGN_KEY) --batch --detach-sign --output statping.asc --armor $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-linux-alpine.tar.gz $(BINARY_NAME) statping.asc && rm -f $(BINARY_NAME) statping.asc
	cd build && mv cmd-darwin-10.6-amd64 $(BINARY_NAME)
	cd build && gpg --default-key $(SIGN_KEY) --batch --detach-sign --output statping.asc --armor $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-osx-x64.tar.gz $(BINARY_NAME) statping.asc && rm -f $(BINARY_NAME) statping.asc
	cd build && mv cmd-darwin-10.6-386 $(BINARY_NAME)
	cd build && gpg --default-key $(SIGN_KEY) --batch --detach-sign --output statping.asc --armor $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-osx-x32.tar.gz $(BINARY_NAME) statping.asc && rm -f $(BINARY_NAME) statping.asc
	cd build && mv cmd-linux-amd64 $(BINARY_NAME)
	cd build && gpg --default-key $(SIGN_KEY) --batch --detach-sign --output statping.asc --armor $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-linux-x64.tar.gz $(BINARY_NAME) statping.asc && rm -f $(BINARY_NAME) statping.asc
	cd build && mv cmd-linux-386 $(BINARY_NAME)
	cd build && gpg --default-key $(SIGN_KEY) --batch --detach-sign --output statping.asc --armor $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-linux-x32.tar.gz $(BINARY_NAME) statping.asc && rm -f $(BINARY_NAME) statping.asc
	cd build && mv cmd-windows-6.0-amd64.exe $(BINARY_NAME).exe
	cd build && gpg --default-key $(SIGN_KEY) --batch --detach-sign --output statping.asc --armor $(BINARY_NAME).exe
	cd build && zip $(BINARY_NAME)-windows-x64.zip $(BINARY_NAME).exe statping.asc && rm -f $(BINARY_NAME).exe statping.asc
	cd build && mv cmd-linux-arm-7 $(BINARY_NAME)
	cd build && gpg --default-key $(SIGN_KEY) --batch --detach-sign --output statping.asc --armor $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-linux-arm7.tar.gz $(BINARY_NAME) statping.asc && rm -f $(BINARY_NAME) statping.asc
	cd build && mv cmd-linux-arm-6 $(BINARY_NAME)
	cd build && gpg --default-key $(SIGN_KEY) --batch --detach-sign --output statping.asc --armor $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-linux-arm6.tar.gz $(BINARY_NAME) statping.asc && rm -f $(BINARY_NAME) statping.asc
	cd build && mv cmd-linux-arm64 $(BINARY_NAME)
	cd build && gpg --default-key $(SIGN_KEY) --batch --detach-sign --output statping.asc --armor $(BINARY_NAME)
	cd build && tar -czvf $(BINARY_NAME)-linux-arm64.tar.gz $(BINARY_NAME) statping.asc && rm -f $(BINARY_NAME) statping.asc

# push the :dev docker tag using curl
publish-dev:
	curl -H "Content-Type: application/json" --data '{"docker_tag": "dev"}' -X POST $(DOCKER)

# update the homebrew application to latest for mac
publish-homebrew:
	curl -s -X POST -H "Content-Type: application/json" -H "Accept: application/json" -H "Travis-API-Version: 3" -H "Authorization: token $(TRAVIS_API)" -d $(PUBLISH_BODY) https://api.travis-ci.com/repo/hunterlong%2Fhomebrew-statping/requests

# install NPM reuqirements for cypress testing
cypress-install:
	cd dev/test && npm install

# run Cypress UI testing
cypress-test: clean cypress-install
	cd dev/test && npm test

# build Statping using a travis ci trigger
travis-build:
	curl -s -X POST -H "Content-Type: application/json" -H "Accept: application/json" -H "Travis-API-Version: 3" -H "Authorization: token $(TRAVIS_API)" -d $(TRAVIS_BUILD_CMD) https://api.travis-ci.com/repo/hunterlong%2Fstatping/requests
	curl -H "Content-Type: application/json" --data '{"docker_tag": "latest"}' -X POST $(DOCKER)

snapcraft-build: build-all
	PWD=$(shell pwd)
	cp build/$(BINARY_NAME)-linux-x64.tar.gz build/$(BINARY_NAME)-linux.tar.gz
	snapcraft clean statping -s pull
	docker run --rm -v ${PWD}:/build -w /build --env VERSION=${VERSION} snapcore/snapcraft bash -c "apt update && snapcraft --target-arch=amd64"
	cp build/$(BINARY_NAME)-linux-x32.tar.gz build/$(BINARY_NAME)-linux.tar.gz
	snapcraft clean statping -s pull
	docker run --rm -v ${PWD}:/build -w /build --env VERSION=${VERSION} snapcore/snapcraft bash -c "apt update && snapcraft --target-arch=i386"
	cp build/$(BINARY_NAME)-linux-arm64.tar.gz build/$(BINARY_NAME)-linux.tar.gz
	snapcraft clean statping -s pull
	docker run --rm -v ${PWD}:/build -w /build --env VERSION=${VERSION} snapcore/snapcraft bash -c "apt update && snapcraft --target-arch=arm64"
	cp build/$(BINARY_NAME)-linux-arm7.tar.gz build/$(BINARY_NAME)-linux.tar.gz
	snapcraft clean statping -s pull
	docker run --rm -v ${PWD}:/build -w /build --env VERSION=${VERSION} snapcore/snapcraft bash -c "apt update && snapcraft --target-arch=armhf"
	rm -f build/$(BINARY_NAME)-linux.tar.gz

snapcraft-release:
	snapcraft push statping_${VERSION}_arm64.snap --release stable
	snapcraft push statping_${VERSION}_i386.snap --release stable
	snapcraft push statping_${VERSION}_armhf.snap --release stable

sign-all:
	gpg --default-key $SIGN_KEY --detach-sign --armor statpinger

valid-sign:
	gpg --verify statping.asc

# install xgo and pull the xgo docker image
xgo-install: clean
	go get github.com/karalabe/xgo
	docker pull karalabe/xgo-latest

heroku:
	git push heroku master
	heroku container:push web
	heroku container:release web

.PHONY: all build build-all build-alpine test-all test test-api docker

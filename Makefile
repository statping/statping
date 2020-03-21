VERSION=$(shell cat version.txt)
SIGN_KEY=B76D61FAA6DB759466E83D9964B9C6AAE2D55278
BINARY_NAME=statping
GOBUILD=go build -a
GOVERSION=1.14.0
XGO=xgo -go $(GOVERSION) --dest=build
BUILDVERSION=-ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=$(TRAVIS_COMMIT)"
TRVIS_SECRET=lRqWSt5BoekFK6+padJF+b77YkGdispPXEUKNuD7/Hxb7yJMoI8T/n8xZrTHtCZPdjtpy7wIlJCezNoYEZB3l2GnD6Y1QEZEXF7MIxP7hwsB/uSc5/lgdGW0ZLvTBfv6lwI/GjQIklPBW/4xcKJtj4s1YBP7xvqyIb/lDN7TiOqAKF4gqRVVfsxvlkm7j4TiPCXtz17hYQfU8kKBbd+vd3PuZgdWqs//5RwKk3Ld8QR8zoo9xXQVC5NthiyVbHznzczBsHy2cRZZoWxyi7eJM1HrDw8Jn/ivJONIHNv3RgFVn2rAoKu1X8F6FyuvPO0D2hWC62mdO/e0kt4X0mn9/6xlLSKwrHir67UgNVQe3tvlH0xNKh+yNZqR5x9t0V54vNks6Pgbhas5EfLHoWn5cF4kbJzqkXeHjt1msrsqpA3HKbmtwwjJr4Slotfiu22mAhqLSOV+xWV+IxrcNnrEq/Pa+JAzU12Uyxs8swaLJGPRAlWnJwzL9HK5aOpN0sGTuSEsTwj0WxeMMRx25YEq3+LZOgwOy3fvezmeDnKuBZa6MVCoMMpx1CRxMqAOlTGZXHjj+ZPmqDUUBpzAsFSzIdVRgcnDlLy7YRiz3tVWa1G5S07l/VcBN7ZgvCwOWZ0QgOH0MxkoDfhrfoMhNO6MBFDTRKCEl4TroPEhcInmXU8=
PUBLISH_BODY='{ "request": { "branch": "master", "message": "Homebrew update version v${VERSION}", "config": { "env": { "VERSION": "${VERSION}", "COMMIT": "$(TRAVIS_COMMIT)" } } } }'
TRAVIS_BUILD_CMD='{ "request": { "branch": "master", "message": "Compile master for Statping v${VERSION}", "config": { "merge_mode": "replace", "language": "go", "install": true, "sudo": "required", "services": ["docker"], "env": { "secure": "${TRVIS_SECRET}" }, "before_deploy": ["git config --local user.name \"hunterlong\"", "git config --local user.email \"info@socialeck.com\"", "git tag v$(VERSION) --force"], "deploy": [{ "provider": "releases", "api_key": "$$TAG_TOKEN", "file_glob": true, "file": "build/*", "skip_cleanup": true, "on": { "branch": "master" } }], "before_script": ["rm -rf ~/.nvm && git clone https://github.com/creationix/nvm.git ~/.nvm && (cd ~/.nvm && git checkout `git describe --abbrev=0 --tags`) && source ~/.nvm/nvm.sh && nvm install stable", "nvm install 10.17.0", "nvm use 10.17.0 --default", "npm install -g sass", "npm install -g cross-env"], "script": ["travis_wait 30 docker pull crazymax/xgo:${GOVERSION}", "make release"], "after_success": [], "after_deploy": ["make publish-homebrew"] } } }'
TEST_DIR=$(GOPATH)/src/github.com/statping/statping
PATH:=/usr/local/bin:$(GOPATH)/bin:$(PATH)

all: clean yarn-install compile docker-base docker-vue build-all compress

up:
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml up -d --remove-orphans
	make print_details

down:
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml down --volumes --remove-orphans

lite: clean
	docker build -t statping/statping:dev -f dev/Dockerfile.dev .
	docker-compose -f dev/docker-compose.lite.yml down
	docker-compose -f dev/docker-compose.lite.yml up --remove-orphans

reup: down clean compose-build-full up

test: clean
	go test -v -p=4 -ldflags="-X main.VERSION=testing" -coverprofile=coverage.out ./...

# build all arch's and release Statping
release: test-deps
	wget -O statping.gpg $(SIGN_URL)
	gpg --import statping.gpg
	make build-all

test-ci: clean compile test-deps
	SASS=`which sass` go test -v -covermode=count -coverprofile=coverage.out -p=1 ./...
	goveralls -coverprofile=coverage.out -service=travis-ci -repotoken ${COVERALLS}

test-api:
	DB_CONN=sqlite DB_HOST=localhost DB_DATABASE=sqlite DB_PASS=none DB_USER=none statping &
	sleep 5000 && newman run source/tmpl/postman.json -e dev/postman_environment.json --delay-request 500

test-deps:
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls
	go get github.com/GeertJohan/go.rice/rice

yarn-serve:
	cd frontend && yarn serve

yarn-install:
	cd frontend && rm -rf node_modules && yarn

go-run:
	go run ./cmd

start:
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml start

stop:
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml stop

logs:
	docker logs statping --follow

db-up:
	docker-compose -f dev/docker-compose.db.yml up -d --remove-orphans

db-down:
	docker-compose -f dev/docker-compose.full.yml down --remove-orphans

console:
	docker exec -t -i statping /bin/sh

compose-build-full: docker-base
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml build --parallel --build-arg VERSION=${VERSION}

docker-base:
	docker build -t statping/statping:base -f Dockerfile.base --build-arg VERSION=${VERSION} .

docker-latest: docker-base
	docker build -t statping/statping:latest --build-arg VERSION=${VERSION} .

docker-vue:
	docker build -t statping/statping:vue --build-arg VERSION=${VERSION} .

docker-test:
	docker-compose -f docker-compose.test.yml up --remove-orphans

push-base: clean compile docker-base
	docker push statping/statping:base

push-vue: clean compile docker-base docker-vue
	docker push statping/statping:base
	docker push statping/statping:vue

modd:
	modd -f ./dev/modd.conf

top:
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml top

frontend-build:
	rm -rf source/dist && rm -rf frontend/dist
	cd frontend && yarn && yarn build
	cp -r frontend/dist source/ && cp -r frontend/src/assets/scss source/dist/
	cp -r source/tmpl/*.* source/dist/

frontend-copy:
	cp -r source/tmpl/*.* source/dist/

yarn:
	rm -rf source/dist && rm -rf frontend/dist
	cd frontend && yarn

# compile assets using SASS and Rice. compiles scss -> css, and run rice embed-go
compile: frontend-build
	rm -f source/rice-box.go
	cd source && rice embed-go

embed:
	cd source && rice embed-go

build:
	$(GOBUILD) $(BUILDVERSION) -o $(BINARY_NAME) ./cmd

install: build
	mv $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)

generate:
	cd source && go generate

# remove files for a clean compile/build
clean:
	rm -rf ./{logs,assets,plugins,*.db,config.yml,.sass-cache,config.yml,statping,build,.sass-cache,index.html,vendor}
	rm -rf cmd/{logs,assets,plugins,*.db,config.yml,.sass-cache,*.log,*.html,*.json}
	rm -rf core/{logs,assets,plugins,*.db,config.yml,.sass-cache,*.log}
	rm -rf types/notifications/{logs,assets,plugins,*.db,config.yml,.sass-cache,*.log}
	rm -rf handlers/{logs,assets,plugins,*.db,config.yml,.sass-cache,*.log}
	rm -rf notifiers/{logs,assets,plugins,*.db,config.yml,.sass-cache,*.log}
	rm -rf source/{logs,assets,plugins,*.db,config.yml,.sass-cache,*.log}
	rm -rf types/{logs,assets,plugins,*.db,config.yml,.sass-cache,*.log}
	rm -rf utils/{logs,assets,plugins,*.db,config.yml,.sass-cache,*.log}
	rm -rf dev/{logs,assets,plugins,*.db,config.yml,.sass-cache,*.log,test/app,plugin/*.so}
	rm -rf {parts,prime,snap,stage}
	rm -rf dev/test/cypress/videos
	rm -f coverage.* sass
	rm -f source/rice-box.go
	rm -rf **/*.db-journal
	rm -rf *.snap
	find . -name "*.out" -type f -delete
	find . -name "*.cpu" -type f -delete
	find . -name "*.mem" -type f -delete
	rm -rf {build,tmp}

print_details:
	@echo \==== Statping Development Instance ====
	@echo \Statping Vue Frontend:     http://localhost:8888
	@echo \Statping Backend API:      http://localhost:8585
	@echo \==== Statping Instances ====
	@echo \Statping SQLite:     http://localhost:4000
	@echo \Statping MySQL:      http://localhost:4005
	@echo \Statping Postgres:   http://localhost:4010
	@echo \==== Databases ====
	@echo \PHPMyAdmin:          http://localhost:6000  \(MySQL database management\)
	@echo \SQLite Web:          http://localhost:6050  \(SQLite database management\)
	@echo \PGAdmin:             http://localhost:7000  \(Postgres database management \| email: admin@admin.com password: admin\)
	@echo \Prometheus:          http://localhost:7050  \(Prometheus Web UI\)
	@echo \==== Monitoring and IDE ====
	@echo \Grafana:             http://localhost:3000  \(username: admin, password: admin\)

build-all: xgo-install compile build-mac build-linux build-windows build-linux build-alpine compress

coverage: test-deps
	$(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=travis -repotoken $(COVERALLS)

# build Statping using a travis ci trigger
travis-build: travis_s3_creds
	curl -s -X POST -H "Content-Type: application/json" -H "Accept: application/json" -H "Travis-API-Version: 3" -H "Authorization: token $(TRAVIS_API)" -d $(TRAVIS_BUILD_CMD) https://api.travis-ci.com/repo/statping%2Fstatping/requests
	curl -H "Content-Type: application/json" --data '{"docker_tag": "latest"}' -X POST $(DOCKER)

download-key:
	wget -O statping.gpg $(SIGN_URL)
	gpg --import statping.gpg

# build Statping for Mac, 64 and 32 bit
build-mac:
	mkdir build
	$(XGO) $(BUILDVERSION) --targets=darwin/amd64,darwin/386 ./cmd

# build Statping for Linux 64, 32 bit, arm6/arm7
build-linux:
	$(XGO) $(BUILDVERSION) --targets=linux/amd64,linux/386,linux/arm-7,linux/arm-6,linux/arm64 ./cmd

# build for windows 64 bit only
build-windows:
	$(XGO) $(BUILDVERSION) --targets=windows-6.0/amd64 ./cmd

# build Alpine linux binary (used in docker images)
build-alpine:
	$(XGO) --targets=linux/amd64 -ldflags="-X main.VERSION=${VERSION} -X main.COMMIT=$(TRAVIS_COMMIT) -linkmode external -extldflags -static" -out alpine ./cmd

# build :latest docker tag
docker-build-latest:
	docker build --build-arg VERSION=${VERSION} -t statping/statping:latest --no-cache -f Dockerfile .
	docker tag statping/statping:latest statping/statping:v${VERSION}

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
	curl -s -X POST -H "Content-Type: application/json" -H "Accept: application/json" -H "Travis-API-Version: 3" -H "Authorization: token $(TRAVIS_API)" -d $(PUBLISH_BODY) https://api.travis-ci.com/repo/statping%2Fhomebrew-statping/requests

upload_to_s3:
	aws s3 cp ./source/css $(ASSETS_BKT) --recursive --exclude "*" --include "*.css"
	aws s3 cp ./source/js $(ASSETS_BKT) --recursive --exclude "*" --include "*.js"
	aws s3 cp ./source/font $(ASSETS_BKT) --recursive --exclude "*" --include "*.eot" --include "*.svg" --include "*.woff" --include "*.woff2" --include "*.ttf" --include "*.css"
	aws s3 cp ./source/scss $(ASSETS_BKT) --recursive --exclude "*" --include "*.scss"
	aws s3 cp ./install.sh $(ASSETS_BKT)

travis_s3_creds:
	mkdir -p ~/.aws
	echo "[default]\naws_access_key_id = ${AWS_ACCESS_KEY_ID}\naws_secret_access_key = ${AWS_SECRET_ACCESS_KEY}" > ~/.aws/credentials

sign-all:
	gpg --default-key $SIGN_KEY --detach-sign --armor statpinger

valid-sign:
	gpg --verify statping.asc

# install xgo and pull the xgo docker image
xgo-install: clean
	go get github.com/crazy-max/xgo
	docker pull crazymax/xgo:${GOVERSION}


.PHONY: all build build-all build-alpine test-all test test-api docker frontend up down print_details lite
.SILENT: travis_s3_creds

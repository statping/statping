VERSION=$(shell cat version.txt)
SIGN_KEY=B76D61FAA6DB759466E83D9964B9C6AAE2D55278
BINARY_NAME=statping
GOBUILD=go build -a
GOVERSION=1.13.5
XGO=xgo -go $(GOVERSION) --dest=build
BUILDVERSION=-ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=$(TRAVIS_COMMIT)"
TRVIS_SECRET=lRqWSt5BoekFK6+padJF+b77YkGdispPXEUKNuD7/Hxb7yJMoI8T/n8xZrTHtCZPdjtpy7wIlJCezNoYEZB3l2GnD6Y1QEZEXF7MIxP7hwsB/uSc5/lgdGW0ZLvTBfv6lwI/GjQIklPBW/4xcKJtj4s1YBP7xvqyIb/lDN7TiOqAKF4gqRVVfsxvlkm7j4TiPCXtz17hYQfU8kKBbd+vd3PuZgdWqs//5RwKk3Ld8QR8zoo9xXQVC5NthiyVbHznzczBsHy2cRZZoWxyi7eJM1HrDw8Jn/ivJONIHNv3RgFVn2rAoKu1X8F6FyuvPO0D2hWC62mdO/e0kt4X0mn9/6xlLSKwrHir67UgNVQe3tvlH0xNKh+yNZqR5x9t0V54vNks6Pgbhas5EfLHoWn5cF4kbJzqkXeHjt1msrsqpA3HKbmtwwjJr4Slotfiu22mAhqLSOV+xWV+IxrcNnrEq/Pa+JAzU12Uyxs8swaLJGPRAlWnJwzL9HK5aOpN0sGTuSEsTwj0WxeMMRx25YEq3+LZOgwOy3fvezmeDnKuBZa6MVCoMMpx1CRxMqAOlTGZXHjj+ZPmqDUUBpzAsFSzIdVRgcnDlLy7YRiz3tVWa1G5S07l/VcBN7ZgvCwOWZ0QgOH0MxkoDfhrfoMhNO6MBFDTRKCEl4TroPEhcInmXU8=
PUBLISH_BODY='{ "request": { "branch": "master", "message": "Homebrew update version v${VERSION}", "config": { "env": { "VERSION": "${VERSION}", "COMMIT": "$(TRAVIS_COMMIT)" } } } }'
TRAVIS_BUILD_CMD='{ "request": { "branch": "master", "message": "Compile master for Statping v${VERSION}", "config": { "os": [ "linux" ], "language": "go", "go": [ "${GOVERSION}" ], "go_import_path": "github.com/hunterlong/statping", "install": true, "sudo": "required", "services": [ "docker" ], "env": { "VERSION": "${VERSION}", "secure": "${TRVIS_SECRET}" }, "matrix": { "allow_failures": [ { "go": "master" } ], "fast_finish": true }, "before_deploy": [ "git config --local user.name \"hunterlong\"", "git config --local user.email \"info@socialeck.com\"", "git tag v$(VERSION) --force"], "deploy": [ { "provider": "releases", "api_key": "$$TAG_TOKEN", "file_glob": true, "file": "build/*", "skip_cleanup": true, "on": {"branch": "master"} } ], "notifications": { "email": false }, "before_script": ["gem install sass"], "script": [ "travis_wait 30 docker pull crazymax/xgo:$(GOVERSION)", "make release" ], "after_success": [], "after_deploy": [ "make publish-homebrew" ] } } }'
TEST_DIR=$(GOPATH)/src/github.com/hunterlong/statping
PATH:=/usr/local/bin:$(GOPATH)/bin:$(PATH)

up:
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml up -d --remove-orphans
	make print_details

down:
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml down --volumes --remove-orphans

lite: clean
	docker build -t hunterlong/statping:dev -f dev/Dockerfile.dev .
	docker-compose -f dev/docker-compose.lite.yml down
	docker-compose -f dev/docker-compose.lite.yml up --remove-orphans

reup: down clean compose-build-full up

start:
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml start

stop:
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml stop

logs:
	docker logs statping --follow

console:
	docker exec -t -i statping /bin/sh

compose-build-full: docker-base
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml build --parallel --build-arg VERSION=${VERSION}

docker-base:
	docker build -t hunterlong/statping:base -f Dockerfile.base --build-arg VERSION=${VERSION} .

docker-latest: docker-base
	docker build -t hunterlong/statping:latest --build-arg VERSION=${VERSION} .

docker-vue:
	docker build -t hunterlong/statping:vue --build-arg VERSION=${VERSION} .

modd:
	modd -f ./dev/modd.conf

top:
	docker-compose -f docker-compose.yml -f dev/docker-compose.full.yml top

frontend-build:
	cd frontend && rm -rf dist && yarn build
	rm -rf source/dist && cp -r frontend/dist source/ && cp -r frontend/src/assets/scss source/dist/
	cp -r source/tmpl/*.* source/dist/

# compile assets using SASS and Rice. compiles scss -> css, and run rice embed-go
compile: generate
	cd source && rice embed-go

build:
	$(GOBUILD) $(BUILDVERSION) -o $(BINARY_NAME) ./cmd

install:
	mv $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)

generate:
	cd source && go generate
	cd handlers/graphql && go generate

# remove files for a clean compile/build
clean:
	rm -rf ./{logs,assets,plugins,*.db,config.yml,.sass-cache,config.yml,statping,build,.sass-cache,index.html,vendor}
	rm -rf cmd/{logs,assets,plugins,*.db,config.yml,.sass-cache,*.log,*.html,*.json}
	rm -rf core/{logs,assets,plugins,*.db,config.yml,.sass-cache,*.log}
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
	rm -rf {build,tmp,docker}

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

.PHONY: all build build-all build-alpine test-all test test-api docker frontend up down print_details lite
.SILENT: travis_s3_creds

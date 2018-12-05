#!/usr/bin/env bash

RELEASE_MASTER='{ "request": { "branch": "master", "config": { "script": "make docker-run-test", "services": ["docker"], "before_script": [], "after_deploy": [], "after_success": ["make publish-homebrew", "make publish-latest"], "deploy": [], "before_deploy": [], "env": { "VERSION": "$(VERSION)" } } } }'

curl -s -X POST -H "Content-Type: application/json" -H "Accept: application/json" -H "Travis-API-Version: 3" -H "Authorization: token $(TRAVIS_API)" -d $(RELEASE_MASTER) https://api.travis-ci.com/repo/hunterlong%2Fstatping/requests
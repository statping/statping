#!/usr/bin/env bash

# update homebrew to newest version by building on travis
body='{ "request": { "branch": "master", "config": { "env": { "VERSION": "'$VERSION'" } } } }'

curl -s -X POST \
 -H "Content-Type: application/json" \
 -H "Accept: application/json" \
 -H "Travis-API-Version: 3" \
 -H "Authorization: token $TRAVIS_API" \
 -d "$body" \
 https://api.travis-ci.com/repo/hunterlong%2Fstatup-testing/requests

#git clone https://$GH_USER:$GH_TOKEN@github.com/hunterlong/homebrew-statup.git
#cd homebrew-statup
#
#./build.sh
#cd ../

curl -H "Content-Type: application/json" --data '{"source_type": "Branch", "source_name": "'"$TRAVIS_BRANCH"'"}' -X POST $DOCKER > /dev/null
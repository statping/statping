#!/usr/bin/env bash

# update homebrew and cypress testing to newest version by building on travis
body='{ "request": { "branch": "master", "config": { "env": { "VERSION": "'$VERSION'" } } } }'

curl -s -X POST \
 -H "Content-Type: application/json" \
 -H "Accept: application/json" \
 -H "Travis-API-Version: 3" \
 -H "Authorization: token $TRAVIS_API" \
 -d "$body" \
 https://api.travis-ci.com/repo/hunterlong%2Fstatup-testing/requests

# notify Docker hub to built this branch
if [ "$TRAVIS_BRANCH" == "master" ]
then
     curl -s -X POST \
     -H "Content-Type: application/json" \
     -H "Accept: application/json" \
     -H "Travis-API-Version: 3" \
     -H "Authorization: token $TRAVIS_API" \
     -d "$body" \
     https://api.travis-ci.com/repo/hunterlong%2Fhomebrew-statup/requests

    curl -H "Content-Type: application/json" --data '{"docker_tag": "dev"}' -X POST $DOCKER
else
    curl -H "Content-Type: application/json" --data '{"source_type": "Branch", "source_name": "'"$TRAVIS_BRANCH"'"}' -X POST $DOCKER > /dev/null
fi

#!/usr/bin/env bash

# update homebrew to newest version by building on travis
body='{
  "request": {
    "branch":"master",
    "config": {
      "merge_mode": "merge",
      "env": {
        "VERSION": "'"$VERSION"'"
      }
    }
  }
}'
curl -s -X POST \
 -H "Content-Type: application/json" \
 -H "Accept: application/json" \
 -H "Travis-API-Version: 3" \
 -H "Authorization: token $TRAVIS_API" \
 -d "$body" \
 https://api.travis-ci.com/repo/hunterlong%2Fhomebrew-statup/requests

#if [ "$TRAVIS_BRANCH" == "master" ]
#then
#    curl -X POST $DOCKER > /dev/null
#else
#    curl -H "Content-Type: application/json" --data '{"source_type": "Tag", "source_name": "v'"$VERSION"'"}' -X POST $DOCKER > /dev/null
#fi
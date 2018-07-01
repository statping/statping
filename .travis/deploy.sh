#!/usr/bin/env bash

# update homebrew to newest version by building on travis
#body='{
#  "request": {
#    "branch":"master",
#    "config": {
#      "merge_mode": "merge",
#      "env": {
#        "VERSION": "'"$VERSION"'"
#      }
#    }
#  }
#}'
#curl -s -X POST \
# -H "Content-Type: application/json" \
# -H "Accept: application/json" \
# -H "Travis-API-Version: 3" \
# -H "Authorization: token $TRAVIS_API" \
# -d "$body" \
# https://api.travis-ci.com/repo/hunterlong%2Fhomebrew-statup/requests


git clone https://$GH_USER:$GH_TOKEN@github.com/hunterlong/homebrew-statup.git
cd homebrew-statup

./build.sh
cd ../

docker login -u $DOCKER_USER -p $DOCKER_PASS $DOCKER_URL
docker build -t $DOCKER_URL/hunterlong/statup .
docker push $DOCKER_URL/hunterlong/statup

if [ "$TRAVIS_BRANCH" != "master" ]
then
    curl -X POST $DOCKER > /dev/null
fi
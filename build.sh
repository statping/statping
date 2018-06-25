#!/usr/bin/env bash

APP="statup"
REPO="hunterlong/statup"

printf "UPDATE core SET version='$VERSION';\n" >> sql/upgrade.sql

rice embed-go

mkdir build
xgo -go 1.10.x --targets=darwin/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=darwin/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=linux/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=linux/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=windows-6.0/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=linux/arm-7 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo -go 1.10.x --targets=linux/arm64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./

CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.VERSION=$VERSION" -a -o build/statup-linux-alpine .

cd build
ls
cd ../

mv build/$APP-darwin-10.6-amd64 build/$APP
tar -czvf build/$APP-osx-x64.tar.gz build/$APP && rm -f build/$APP

mv build/$APP-darwin-10.6-386 build/$APP
tar -czvf build/$APP-osx-x32.tar.gz build/$APP && rm -f build/$APP

mv build/$APP-linux-amd64 build/$APP
tar -czvf build/$APP-linux-x64.tar.gz build/$APP && rm -f build/$APP

mv build/$APP-linux-386 build/$APP
tar -czvf build/$APP-linux-x32.tar.gz build/$APP && rm -f build/$APP

mv build/$APP-windows-6.0-amd64.exe build/$APP.exe
zip $APP-windows-x64.zip build/$APP.exe  && rm -f build/$APP.exe

mv build/$APP-linux-arm-7 build/$APP
tar -czvf build/$APP-linux-arm7.tar.gz build/$APP && rm -f build/$APP

mv build/$APP-linux-arm64 build/$APP
tar -czvf build/$APP-linux-arm64.tar.gz build/$APP && rm -f build/$APP

mv build/$APP-linux-alpine build/$APP
tar -czvf build/$APP-linux-alpine.tar.gz build/$APP && rm -f build/$APP

body='{
 "request": {
 "message": "Updating Homebrew Formula to v'"$VERSION"'",
 "branch":"master",
 "config": {
   "env": {
     ["VERSION='"$VERSION"'"]
   }
  }
}}'

curl -s -X POST \
 -H "Content-Type: application/json" \
 -H "Accept: application/json" \
 -H "Travis-API-Version: 3" \
 -H "Authorization: $TRAVIS_API" \
 -d "$body" \
 https://api.travis-ci.com/repo/hunterlong%2Fhomebrew-statup/requests
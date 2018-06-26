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

CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.VERSION=$VERSION" -a -o build/$APP-linux-alpine .

cd build
ls

mv $APP-linux-alpine $APP
tar -czvf $APP-linux-alpine.tar.gz $APP && rm -f $APP

mv $APP-darwin-10.6-amd64 $APP
tar -czvf $APP-osx-x64.tar.gz $APP && rm -f $APP

mv $APP-darwin-10.6-386 $APP
tar -czvf $APP-osx-x32.tar.gz $APP && rm -f $APP

mv $APP-linux-amd64 $APP
tar -czvf $APP-linux-x64.tar.gz $APP && rm -f $APP

mv $APP-linux-386 $APP
tar -czvf $APP-linux-x32.tar.gz $APP && rm -f $APP

mv $APP-windows-6.0-amd64.exe $APP.exe
zip $APP-windows-x64.zip $APP.exe  && rm -f $APP.exe

mv $APP-linux-arm-7 $APP
tar -czvf $APP-linux-arm7.tar.gz $APP && rm -f $APP

mv $APP-linux-arm64 $APP
tar -czvf $APP-linux-arm64.tar.gz $APP && rm -f $APP
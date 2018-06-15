#!/usr/bin/env bash

mkdir build

APP="statup"

rice embed-go

xgo --targets=darwin/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo --targets=darwin/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./

xgo --targets=linux/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./
xgo --targets=linux/386 --dest=build -ldflags="-X main.VERSION=$VERSION" ./

xgo --targets=windows/amd64 --dest=build -ldflags="-X main.VERSION=$VERSION" ./


cd build
ls
cd ../

mv build/$APP-darwin-10.6-amd64 build/$APP-osx-x64
mv build/$APP-darwin-10.6-386 build/$APP-osx-x32
mv build/$APP-linux-amd64 build/$APP-linux-x64
mv build/$APP-linux-386 build/$APP-linux-x32
mv build/$APP-windows-4.0-amd64.exe build/$APP-windows-x64.exe

#!/usr/bin/env bash

cd .travis
cp Dockerfile.dev ../
cd ../
docker build -t hunterlong/statup:dev -f Dockerfile.dev .

rm -rf Dockerfile.dev
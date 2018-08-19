#!/usr/bin/env bash

DIR=`pwd`
DOCKER=`which docker`

$DOCKER build -t hunterlong/statup:dev -f ../Dockerfile ../../
$DOCKER run -it -d -p 8080:8080 -v $DIR/app:/app --name statup_dev hunterlong/statup:dev

./node_modules/.bin/start-server-and-test start http://localhost:8080/robots.txt cy:run

$DOCKER stop statup_dev || true && $DOCKER rm -f statup_dev || true

sudo rm -rf $DIR/app
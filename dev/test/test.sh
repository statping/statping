#!/usr/bin/env bash

DIR=`pwd`
DOCKER=`which docker`

$DOCKER build -t hunterlong/statping:dev -f ../Dockerfile ../../
$DOCKER run -it -d -p 8080:8080 -v $DIR/app:/app --name statping_dev hunterlong/statping:dev

./node_modules/.bin/start-server-and-test start http://localhost:8080/robots.txt cy:run

$DOCKER stop statping_dev || true && $DOCKER rm -f statping_dev || true

sudo rm -rf $DIR/app
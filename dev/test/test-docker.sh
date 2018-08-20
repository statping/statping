#!/usr/bin/env bash

statup > /dev/null &

./node_modules/.bin/start-server-and-test start http://localhost:8080/robots.txt cy:run
#!/usr/bin/env sh

modd -f ./dev/modd.conf

devd -w ./src http://localhost:8585

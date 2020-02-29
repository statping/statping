#!/usr/bin/env sh

echo "Starting!"

echo "Serving Vue frontend first..."

cd frontend && yarn serve &

echo "Now serving Vue, lets build the golang backend now..."

modd -f dev/modd.conf

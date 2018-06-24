#!/usr/bin/env bash

VERSION=v0.18

wget https://github.com/hunterlong/statup/releases/download/$VERSION/statup-linux-x64
mv statup-linux-x64 /usr/local/bin/statup
chmod +x /usr/local/bin/statup
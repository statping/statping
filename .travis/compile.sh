#!/usr/bin/env bash

cd $GOPATH/src/github.com/hunterlong/statup/core

# RENDERING CSS
gem install sass
sass ../source/scss/base.scss ../source/css/base.css

# MIGRATION SQL FILE FOR CURRENT VERSION
#printf "UPDATE core SET version='$VERSION';\n" >> source/sql/upgrade.sql

# COMPILE SRC INTO BIN
rice embed-go

cd $GOPATH/src/github.com/hunterlong/statup/cmd

go install

mv $GOPATH/bin/cmd $GOPATH/bin/statup

cd $GOPATH/src/github.com/hunterlong/statup
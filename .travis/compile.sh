#!/usr/bin/env bash

# RENDERING CSS
gem install sass
sass html/scss/base.scss html/css/base.css

# MIGRATION SQL FILE FOR CURRENT VERSION
printf "UPDATE core SET version='$VERSION';\n" >> sql/upgrade.sql

# COMPILE SRC INTO BIN
rice embed-go

go install
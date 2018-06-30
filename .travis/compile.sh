#!/usr/bin/env bash

# RENDERING CSS
gem install sass
sass source/scss/base.scss source/css/base.css

# MIGRATION SQL FILE FOR CURRENT VERSION
printf "UPDATE core SET version='$VERSION';\n" >> source/sql/upgrade.sql

# COMPILE SRC INTO BIN
rice embed-go

go install
#!/usr/bin/env bash

# COMPILE BOOTSTRAP
rm -rf bootstrap
git clone https://github.com/twbs/bootstrap.git
cd bootstrap
npm install
rm -f scss/_variables.scss
cp ../html/scss/_variables.scss scss/_variables.scss
npm run dist
mv dist/css/bootstrap.min.css ../html/css/bootstrap.min.css
cd ../
rm -rf bootstrap

# RENDERING CSS
gem install sass
sass html/scss/base.scss html/css/base.css

# MIGRATION SQL FILE FOR CURRENT VERSION
printf "UPDATE core SET version='$VERSION';\n" >> sql/upgrade.sql

# COMPILE SRC INTO BIN
rice embed-go

go install
#!/usr/bin/env bash
VERSION=v0.22
OS=osx
ARCH=x64
if [ `getconf LONG_BIT` = "64" ]
then
    ARCH=x64
else
    ARCH=x32
fi
unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     OS=linux;;
    Darwin*)    OS=osx;;
    CYGWIN*)    OS=windows;;
    MINGW*)     OS=windows;;
    *)          OS="UNKNOWN:${unameOut}"
esac
FILE="https://github.com/hunterlong/statup/releases/download/$VERSION/statup-$OS-$ARCH"
curl -sS $FILE -o statup
chmod +x statup
mv statup /usr/local/bin/
statup version
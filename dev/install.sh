#!/usr/bin/env bash
OS=osx
ARCH=x64
REPO=github.com/hunterlong/statping
VERSION=$(curl -s "https://$REPO/releases/latest" | grep -o 'tag/[v.0-9]*' | awk -F/ '{print $2}')
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
printf "Installing $VERSION for $OS $ARCH...\n"
FILE="https://$REPO/releases/download/$VERSION/statping-$OS-$ARCH.tar.gz"
printf "Downloading latest version URL: $FILE\n"
curl -L -sS $FILE -o statping.tar.gz && tar xzf statping.tar.gz && rm statping.tar.gz
chmod +x statping
echo "Installing Statping to directory: /usr/local/bin/"
mv statping /usr/local/bin/
echo "Statping $VERSION has been successfully installed! Try 'statping version' to check it!"
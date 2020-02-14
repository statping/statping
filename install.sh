#!/bin/bash
# Statping installation script for Linux, Mac, and maybe Windows.
#
# This installation script is a modification of Yarn's installation
set -e

reset="\033[0m"
red="\033[31m"
green="\033[32m"
yellow="\033[33m"
cyan="\033[36m"
white="\033[37m"
gpg_key=64B9C6AAE2D55278
gpgurl=https://statping.com/statping.gpg
repo=https://github.com/hunterlong/statping

statping_get_tarball() {
  fext='tar.gz'
  if [ ${OS} == 'windows' ]; then
    fext='zip'
    ARCH='x64'
  fi
  url="$repo/releases/latest/download/statping-$1-$2.$fext"
  printf "$cyan> Downloading latest version for $OS $ARCH...\n$url $reset\n"
  # Get both the tarball and its GPG signature
  tarball_tmp=`mktemp -t statping.tar.gz.XXXXXXXXXX`
  if curl --fail -L -o "$tarball_tmp" "$url"; then
    # All this dance is because `tar --strip=1` does not work everywhere
    temp=$(mktemp -d statping.XXXXXXXXXX)
    if [ ${OS} == 'windows' ]; then
      unzip $tarball_tmp -d "$temp"
    else
      tar xzf $tarball_tmp -C "$temp"
    fi
    statping_verify_integrity "$temp"/statping
    printf "$green> Installing to $DEST/statping\n"
    mv "$temp"/statping "$DEST"
    newversion=`$DEST/statping version`
    rm -rf "$temp"
    rm $tarball_tmp*
    printf "$cyan> Statping is now installed! $reset\n"
    printf "$white>   Version:  $newversion $reset\n"
    printf "$white>   Repo:     $repo $reset\n"
    printf "$white>   Wiki:     $repo/wiki $reset\n"
    printf "$white>   Issues:   $repo/issues $reset\n"
    printf "$cyan> Try to run \"statping help\" $reset\n"
  else
    printf "$red> Failed to download $url.$reset\n"
    exit 1;
  fi
}

# Verifies the GPG signature of the tarball
statping_verify_integrity() {
  # Check if GPG is installed
  if [[ -z "$(command -v gpg)" ]]; then
    printf "$yellow> WARNING: GPG is not installed, integrity can not be verified!$reset\n"
    return
  fi

  if [ "$statping_GPG" == "no" ]; then
    printf "$cyan> WARNING: Skipping GPG integrity check!$reset\n"
    return
  fi

  printf "$cyan> Verifying integrity with gpg key from $gpgurl...$reset\n"
  # Grab the public key if it doesn't already exist
  gpg --list-keys $gpg_key >/dev/null 2>&1 || (curl -sS -L $gpgurl | gpg --import)

  if [ ! -f "$1.asc" ]; then
    printf "$red> Could not download GPG signature for this Statping release. This means the release can not be verified!$reset\n"
    statping_verify_or_quit "> Do you really want to continue?"
    return
  fi

  # Actually perform the verification
  if gpg --verify "$1.asc" $1 &> /dev/null; then
    printf "$green> GPG signature looks good$reset\n"
  else
    printf "$red> GPG signature for this Statping release is invalid! This is BAD and may mean the release has been tampered with. It is strongly recommended that you report this to the Statping developers.$reset\n"
    statping_verify_or_quit "> Do you really want to continue?"
  fi
}

statping_reset() {
  unset -f statping_install statping_reset statping_get_tarball statping_verify_integrity statping_verify_or_quit statping_brew_install getOS getArch
}

statping_brew_install() {
  if [[ -z "$(command -v brew --version)" ]]; then
    printf "${white}Using Brew to install!$reset\n"
    printf "${yellow}---> brew tap hunterlong/statping$reset\n"
    brew tap hunterlong/statping
    printf "${yellow}---> brew install statping$reset\n"
    brew install statping
    printf "${green}Brew installation is complete!$reset\n"
    printf "${yellow}You can use 'brew upgrade' to upgrade Statping next time.$reset\n"
  else
    statping_get_tarball $OS $ARCH
  fi
}

statping_install() {
  printf "${white}Installing Statping!$reset\n"
  getOS
  getArch
  if [ "$OS" == "osx" ]; then
      statping_brew_install
    else
      statping_get_tarball $OS $ARCH
  fi
  statping_reset
}

statping_verify_or_quit() {
  read -p "$1 [y/N] " -n 1 -r
  echo
  if [[ ! $REPLY =~ ^[Yy]$ ]]
  then
    printf "$red> Aborting$reset\n"
    exit 1
  fi
}

# get the users operating system
getOS() {
    OS="`uname`"
    case $OS in
      'Linux')
        OS='linux'
        DEST=/usr/local/bin
        alias ls='ls --color=auto'
        ;;
      'FreeBSD')
        OS='freebsd'
        DEST=/usr/local/bin
        alias ls='ls -G'
        ;;
      'WindowsNT')
        OS='windows'
        DEST=/usr/local/bin
        ;;
      'MINGW*')
        OS='windows'
        DEST=/usr/local/bin
        ;;
      'CYGWIN*')
        OS='windows'
        DEST=/usr/local/bin
        ;;
      'Darwin')
        OS='osx'
        DEST=/usr/local/bin
        ;;
      'SunOS')
        OS='solaris'
        DEST=/usr/local/bin
        ;;
      'AIX') ;;
      *) ;;
    esac
}

# get 64x or 32 machine arch
getArch() {
    MACHINE_TYPE=`uname -m`
    if [ ${MACHINE_TYPE} == 'x86_64' ]; then
      ARCH="x64"
    else
      ARCH="x32"
    fi
}

cd ~
statping_install $1 $2

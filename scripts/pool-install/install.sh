#! /bin/bash

poolInstallDir="~/pool-install-files"
GOPATH="/usr/lib/go-1.10"
GOBIN="$GOPATH/bin"
PATH="$PATH:$GOBIN"
export "$PATH"
sudo add-apt-repository ppa:gophers/archive
sudo apt-get update
sudo apt-get install golang-1.10-go


sudo apt-get install software-properties-common
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install ethereum


echo "$PATH" >> "~/.bashrc"
mkdir "$poolInstallDir"
cd "$poolInstallDir"
git config --global http.https://gopkg.in.followRedirects true
git clone https://github.com/sammy007/open-ethereum-pool.git
cd open-ethereum-pool
make
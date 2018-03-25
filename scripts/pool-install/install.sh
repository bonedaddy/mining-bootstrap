#! /bin/bash

poolInstallDir="~/pool-install-files"
GOPATH="/usr/lib/go-1.10"
GOBIN="$GOPATH/bin"
PATH="$PATH:$GOBIN"
export "$PATH"
sudo add-apt-repository ppa:gophers/archive
sudo apt-get update -y
sudo apt-get install golang-1.10-go -y


sudo apt-get install software-properties-common
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get update -y
sudo apt-get install ethereum -y


echo "$PATH" >> "~/.bashrc"
mkdir "$poolInstallDir"
cd "$poolInstallDir"
git config --global http.https://gopkg.in.followRedirects true
git clone https://github.com/sammy007/open-ethereum-pool.git
cd open-ethereum-pool
make

sudo apt-get update -y
sudo apt-get install redis-server -y


echo "[INFO] To  run the pool backend use the following command"
echo "$poolInstallDir/build/bin/open-ethereum-pool $poolInstallDir/config.json"

echo "[INFO] Starting to build pool frontend"
sudo apt-get install nodejs -y
sudo apt-get install npm -y

#cd "$poolInstallDir/www"
#su
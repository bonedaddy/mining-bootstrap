#! /bin/bash

# install that cuda goodness baby
cd ~
# download cuda 9
wget https://developer.nvidia.com/compute/cuda/9.2/Prod/local_installers/cuda-repo-ubuntu1604-9-2-local_9.2.88-1_amd64
sudo dpkg -i cuda-repo-ubuntu1604-9-2-local_9.2.88-1_amd64
sudo apt-key add /var/cuda-repo-9-2-local/7fa2af80.pub
sudo apt-get update -y
sudo apt-get install cuda-9.2 -y
# install updates
sudo add-apt-repository ppa:ubuntu-toolchain-r/test
sudo apt-get update -y
sudo apt-get upgrade -y
# download ethminer
wget https://github.com/ethereum-mining/ethminer/releases/download/v0.15.0.dev11/ethminer-0.15.0.dev11-Linux.tar.gz
tar zxvf ethminer-0.15.0.dev11-Linux.tar.gz
# make out boot script directory
sudo mkdir /boot_scripts
sudo cp -r bin /boot_scripts





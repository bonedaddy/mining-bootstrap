#! /bin/bash

# install that cuda goodness baby
cd ~
# download cuda 8
wget https://developer.nvidia.com/compute/cuda/8.0/Prod2/local_installers/cuda-repo-ubuntu1604-8-0-local-ga2_8.0.61-1_amd64-deb
sudo dpkg -i cuda-repo-ubuntu1604-8-0-local-ga2_8.0.61-1_amd64-deb
sudo apt-get update -y
# install cuda
sudo apt-get install cuda -y
# download patch 2
wget https://developer.nvidia.com/compute/cuda/8.0/Prod2/patches/2/cuda-repo-ubuntu1604-8-0-local-cublas-performance-update_8.0.61-1_amd64-deb
sudo dpkg -i cuda-repo-ubuntu1604-8-0-local-cublas-performance-update_8.0.61-1_amd64-deb
# refresh repo
sudo apt-get update -y 
# run the upgrade
sudo apt-get upgrade -y
# download ethminer
wget https://github.com/ethereum-mining/ethminer/releases/download/v0.15.0.dev11/ethminer-0.15.0.dev11-Linux.tar.gz
tar zxvf ethminer-0.15.0.dev11-Linux.tar.gz
# make out boot script directory
sudo mkdir /boot_scripts
sudo cp -r bin /boot_scripts

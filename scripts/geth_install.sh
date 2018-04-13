#! /bin/bash

# script to install go-ethereum on ubuntu 

sudo apt-get install software-properties-common
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install ethereum
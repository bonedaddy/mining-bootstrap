#! /bin/bash

sudo add-apt-repository ppa:graphics-drivers/ppa -y
sudo apt-get update -y
sudo apt upgrade -y
sudo apt install nvidia-384 -y

#! /bin/bash

DISTRO=$(lsb_release -sc)

# Different ways to install nvidia
# * https://www.linuxbabe.com/ubuntu/install-nvidia-driver-ubuntu-18-04

# validate the proper version
if [[ "$DISTRO" != "bionic" ]]; then
    echo "[ERROR] Installation only supported for 18.04"
    exit 1
fi

echo "[INFO] Updating system"
sudo apt update -y
echo "[INFO] Upgrading system"
sudo apt upgrade -y
echo "[INFO] Installing openssh-server"
sudo apt install openssh-server -y
echo "[INFO] Installing git"
sudo apt install git -y
echo "[INFO] Downloading mining-bootstrap repo"
cd ~ || exit
git clone https://github.com/RTradeLtd/mining-bootstrap.git
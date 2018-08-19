#! /bin/bash

# Fresh Install Setup
# 1) Install Ubuntu 18.04 Desktop first, without any GPU's installed
# 2) After instalation of the system, boot into the system, and confirm initial setup instructions
# 3) Shutdown the machine, plug in your GPU, and start the machine
# 4) After starting up, run this script and after reboot verify installation
# 5) Shutdown, plug the remaining GPUs in and you'll be good to go


DISTRO=$(lsb_release -sc)
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
echo "[INFO] Installing vim"
sudo apt install vim -y
echo "[INFO] Downloading mining-bootstrap repo"

cd ~ || exit

git clone https://github.com/RTradeLtd/mining-bootstrap.git

echo "[INFO] Beginning nvidia driver installation"
RECOMMENDED_VERSION=$(sudo ubuntu-drivers devices | grep -i driver | grep -i nvidia | awk '{print $3}')
echo "[INFO] Installing driver version $RECOMMENDED_VERSION"
sudo apt install "$RECOMMENDED_VERSION" -y

if [[ "$?" != "0" ]]; then
    echo "[ERROR] Installation of nvidia drivers failed"
    exit 1
fi

echo "[INFO] Nvidia drivers installed. Setting default boot target to multi-user.target"
sudo systemctl set-default multi-user.target
echo "[INFO] Ubuntu 18.04 mining bootstrap finished, sleeping for 15 seconds before rebooting"
echo "[INFO] Hit CTRL+C to cancel reboot"
sleep 15
sudo shutdown -r now
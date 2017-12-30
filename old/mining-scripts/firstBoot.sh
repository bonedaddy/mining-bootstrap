#! /bin/bash


cudaVersion=8

echo "[INFO] Ensure that you have installed the nvidia drivers first"

# perform initial updates
sudo apt-get update -y
sudo apt-get upgrade -y
# install dependences
sudo apt-get install libssl-dev libcurl4-openssl-dev vim automake git -y
# create scripts directory
mkdir ~/scripts
cd ~/scripts
# download some pre-written scripts
git clone https://github.com/createthis/linux_gpu_mining.git
cd linux_gpu_mining
# enable fan control
./enable_nvidia_fan_control.sh
# create boot script directory
sudo mkdir /boot_scripts
sudo chmod a+rw /boot_scripts
# copy underclock
cp underclock.sh /boot_scripts
# make directory to store boot scripts
mkdir -p nvidia-temp/cuda
cd nvidia-temp/cuda
# download cuda repos
wget https://developer.download.nvidia.com/compute/cuda/repos/ubuntu1604/x86_64/cuda-repo-ubuntu1604_9.0.176-1_amd64.deb
# install repo keys
sudo apt-key adv --fetch-keys http://developer.download.nvidia.com/compute/cuda/repos/ubuntu1604/x86_64/7fa2af80.pub
# install repo
sudo dpkg -i cuda-repo* -y
# update to refresh repo list
sudo apt-get update -y
mkdir /logs
if [[ "$cudaVersion" -eq 8 ]]; then
    # install cuda 
    sudo apt-get install cuda-8-0 -y
elif [[ "$cudaVersion" -eq 9 ]]; then
    # install cuda
    sudo apt-get install cuda-9-0 -y
fi

mkdir /tmp/ccminer_download
cd /tmp/ccminer_download
wget https://github.com/djm34/ccminer-msvc2015/archive/v0.2.1.tar.gz
tar zxvf v0.2.1.tar.gz
cd *
./autogen.sh
./configure.sh
make
sudo make install
which ccminer
if [[ "$?" -eq 0 ]]; then
    echo "ccminer installed rebooting" > /dev/null
    sudo reboot now
else
    echo "ccminer not installed"
fi

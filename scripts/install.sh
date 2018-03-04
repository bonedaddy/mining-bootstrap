#! /bin/bash

# make boot script dir
sudo mkdir /boot_scripts
# set permissinos
sudo chmod a+rw /boot_scripts
# set the services
cd /home/rtrade
cp /home/rtrade/mining-bootstrap/services.tar .
tar xvf services.tar
sudo cp services/*.service /etc/systemd/
sudo cp /home/rtrade/ethminer.sh /boot_scripts

# uncomment depending on whether or not you want
# to setup any services
# sudo systemctl enable ccminer.service
# sudo systemctl enable miner.service
# sudo systemctl enable ethminer.service
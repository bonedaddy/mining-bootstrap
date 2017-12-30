git clone https://github.com/postables/mining-bootstrap.git
cd mining-bootstrap
tar xvf temp.tar
cd temp
cp scripts/* /boot_scripts
chmod a_x /boot_scripts/*.sh
sudo cp services/*.service /etc/systemd/system
sudo systemctl enable /etc/systemd/system/ccminer.service
sudo systemctl enable /etc/systemd/system/miner.service
sudo reboot now

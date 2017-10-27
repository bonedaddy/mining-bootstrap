#! /bin/bash

# Used to install the service

sudo cp bootstrapper.service /etc/systemd/system
sudo systemctl daemon-reload
sudo systemctl enable bootstrapper.service
sudo systemctl start  bootstrapper.service
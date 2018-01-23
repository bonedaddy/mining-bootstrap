#! /bin/bash

# used to fix nvidia driver issues

sudo apt remove bbswitch-dkms nvidia* cuda* --purge -y

sudo bash /home/rtrade/cuda_8_install.run

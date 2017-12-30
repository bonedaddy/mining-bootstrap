#! /bin/bash

# author: postables
# version: 1.0.1
# description .sh: shell script to mine any equihash algorithm
# description .j2: template file used by ansible

/boot_scripts/miner --server kmd.suprnova.cc --port 6250 --user postables."$(hostname)" --pass password --cuda_devices 6 7 8 9 10 11

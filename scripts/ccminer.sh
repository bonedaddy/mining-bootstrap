#! /bin/bash

# author: postables
# version: 1.0.1
# description .sh: ccminer script used to mine any algorithm supported by the current ccminer version
# description .j2: ccminer template used by ansible

/usr/local/bin/ccminer --algo=lyra2z --url=stratum+tcp://us-east.lyra2z-hub.miningpoolhub.com:20581 --user=postables."$(hostname)" --pass=password --devices=0,1,2,3,4,5

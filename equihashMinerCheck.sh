#! /bin/bash

# Author: Postables
# Version: 0.0.1alpha
# Description: Used by Zabbix to check for equihash miner

equiCheck=$(ps aux | grep -i miner | grep -v grep)

if [[ -z "$equiCheck" ]]; then
    echo "equihash miner not detected"
    echo 0 > /tmp/equicheck.log
else
    echo "equihash miner detected"
    echo 1 > /tmp/equicheck.log
fi
#! /bin/bash

# Author: Postables
# Version: 0.0.1alpha
# Description: Used by Zabbix to check for ccminer, designed to be called by the zabbix agent

cCheck=$(ps aux | grep -i ccminer | grep -v grep)
if [[ -z "$cCheck" ]]; then
    echo "ccminer not detected" > /dev/null
    echo 0 > /tmp/ccminerCheck.log
else
    echo "ccminer detected" > /dev/null
    echo 1 > /tmp/ccminerCheck.log
fi
#! /bin/bash

# quick network fix for device unit hang issues

device_id=$(ifconfig | head -n 1 | awk -F ':' '{print $1}')
ethtool -K "$device_id" gso off gro off tso off

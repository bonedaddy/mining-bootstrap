#! /bin/sh

# Used to launch the mining process 
hostname="$(hostname)"
if [ -z "$STY" ]; then exec screen -dm -S "$hostname" /boot_scripts/miner.sh "$0"; fi

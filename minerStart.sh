#! /bin/sh

# Used to launch the mining process 

if [ -z "$STY" ]; then exec screen -dm -S miningscreen /boot_scripts/miner.sh "$0"; fi

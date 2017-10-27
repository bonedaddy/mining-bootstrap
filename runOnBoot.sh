#! /bin/bash

# Author: Postables
# Version: 0.0.1alpha
# Description: This is used to execute our startup/bootstrap sequence during a boot to auto-start the miners, and configure relevant system variables

bootScriptDir="/boot_scripts"
logDir="/logs"
rigName="$(hostname)"
poolUrl="stratum+tcp://hub.miningpoolhub.com:20507"
poolUser="postables.$rigName"
poolPassword="password"
minerAlgo="lyra2v2"


# Execute underclock
mv "$logDir/boot.log" "$logDir/previousBoot.log"
touch "$logDir/boot.log"

printf "[INFO] Starting Underclock Script\n\n"   >> "$logDir/boot.log"
bash -x "$bootScriptDir/underclock.sh"           >> "$logDir/boot.log"
if [[ "$?" -eq 0 ]]; then
	printf "[INFO] Underlock Script Finished \n\n"   >> "$logDir/boot.log"
else
	printf "[ERROR] Underclock script failed to start, going to exit" >> "$logDir/boot.log"
	exit
fi

bash /boot_scripts/miner.sh &
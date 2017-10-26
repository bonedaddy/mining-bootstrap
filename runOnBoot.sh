#! /bin/bash

# Author: Postables
# Version: 0.0.1alpha
# Description: This is used to execute our startup/bootstrap sequence during a boot to auto-start the miners, and configure relevant system variables

bootScriptDir="/boot_scripts"
logDir="/logs"

# Execute underclock
mv "$logDir/boot.log" "$logDir/previousBoot.log"
touch "$logDir/boot.log"

printf "[INFO] Starting Underclock Script\n\n"   >> "$logDir/boot.log"
bash -x "$bootScriptDir/underclock.sh"           >> "$logDir/boot.log"
if [[ "$?" -eq 0 ]]; then
	printf "[INFO] Underlock Script Finished \n\n"   >> "$logDir/boot.log"
else
	printf "[ERROR] Underclock script failed to start, going to exit"
fi

printf "[INFO] Starting Miner Launch Script\n\n" >> "$logDir/boot.log"
# Note: We must use /bin/sh to do this, because if we use /bin/bash, it will pass in the whole argument string to screen as a single option, whereas with /bin/sh it will pass it in as a list of different words. Also note that the "$0" refers to the script name, so we're going to call that script independently since it will be using a different shell
"$bootScriptDir/minerStart.sh" 

if [[ "$?" -eq 0 ]]; then
	printf "[INFO] Miner launched successfully, going to exit\n\n" >> "$logDir/boot.log"
else
	printf "[ERROR] Miner not launched successfully\n\n" >> "$logDir/boot.log"
fi

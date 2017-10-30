#! /bin/bash

# Author: Postables
# Version: 0.0.3alpha
# Description: This is used to execute our startup/bootstrap sequence during a boot to auto-start the miners, and configure relevant system variables

bootScriptDir="/boot_scripts"
logDir="/logs"
rigName="$(hostname)"
poolUrl="stratum+tcp://hub.miningpoolhub.com:20507"
poolUser="postables.$rigName"
poolPassword="password"
minerAlgo="lyra2v2"
minerToRun="ccminer"
equiHashUrl="zdash.suprnova.cc"
equiHashPort=4048
equiHashUser="postbales.$rigName"
equiHashPassword="password"

# Execute underclock
mv "$logDir/mine.log" "$logDir/previous.log"
touch "$logDir/mine.log"

printf "[INFO] Starting Underclock Script\n\n"   &>> "$logDir/mine.log"
bash "$bootScriptDir/underclock.sh"           &>> "$logDir/mine.log"

if [[ "$?" -eq 0 ]]; then
	printf "[INFO] Underlock Script Finished \n\n"   &>> "$logDir/mine.log"
else
	printf "[ERROR] Underclock script failed to start, going to exit" &>> "$logDir/mine.log"
	exit
fi


if [[ "$minerToRun" -eq ccminer ]]; then
	# since we don't care about supporting pre bash 4, we use the new version of appending stderr and stdout
	ccminer --algo="$minerAlgo" --url="$poolUrl" --user "$poolUser" --pass "$poolPassword" &>> "$logDir/mine.log"
elif [[ "$minerToRun" -eq miner ]]; then	
	miner  --server "$equiHashUrl" --port "$equiHashPort" --user "$equiHashUser" --pass "$equiHashPassword" &>> "$logDir/mine.log"
fi
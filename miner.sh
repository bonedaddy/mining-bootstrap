#! /bin/bash

# Author: Postables
# Version: 0.0.1alpha
# Description: Script is used to launch ccminer, and provide configurable variabels for easy management

bootDir="/boot_scripts"
runmode="$1"
rigName="$(hostname)"
suprnovaurl="suprnova.cc"
suprnovazdashport=4048
suprnovauser=.....
suprnovapassword=....

if [[ "$runmode" -eq "hush" || "$runmode" -eq "zdash" ]]; then
	"$bootdir"/miner --server "zdash.$suprnovaurl:suprnovazdashport" --user "suprnova.rigName" --pass "$suprnovapassword"
fi
#! /bin/bash

# used to check if auto updates are enabled

grep 'APT::Periodic::Update-Package-Lists "1";' /etc/apt/apt.conf.d/10periodic
if [[ "$?" -eq 0 ]]; then
	# auto updates enabled so we need to disable
	sed -i 's/'APT::Periodic::Update-Package-Lists "1";'/'APT::Periodic::Update-Package-Lists "0";'/g' /etc/apt/apt.conf.d/10periodic
fi


grep 'APT::Periodic::Update-Package-Lists "0";' /etc/apt/apt.conf.d/10periodic
if [[ "$?" -eq 0 ]]; then
	# everything okay, exit silently
	exit 0
else
	# not okay, exit and make noise
	echo "[ERROR] AUTO UPDATES ENABLED STILL OR OTHER ERROR"
	exit 1
fi

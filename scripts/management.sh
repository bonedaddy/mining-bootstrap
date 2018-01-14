#! /bin/bash

# management script

case "$1" in
	restart-miner)
		sudo systemctl restart miner.service
		;;
	restart-ccminer)
		sudo systemctl restart ccminer.service
		;;
	start-miner)
		sudo systemctl start miner.service
		;;
	start-ccminer)
		sudo systemctl start ccminer.service
		;;
	stop-miner)
		sudo systemctl stop miner.service
		;;
	stop-ccminer)
		sudo systemctl stop ccminer.service
		;;
	update)
		sudo apt update
		;;
	upgrade)
		sudo apt upgrade
		;;
	remove-rc)
		# removes packages marked 'rc' by dpkg
		for pkg in $(dpkg --list | grep "^rc" | awk '{print $2}'); do
			echo 'remove "$pkg"'
			sudo dpkg --purge "$pkg"
		done
		;;
	*)
		echo $"Usage: $0 {restart-miner|restart-ccminer|start-miner|start-ccminer|stop-miner|stop-ccminer|update|upgrade|remove-rc}"
		exit 1
esac


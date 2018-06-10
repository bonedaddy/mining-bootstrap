# V4 (WIP)

Collection of scripts, and various utilities used to setup, manage, and monitor mining rigs.
Each mining rig is presumed to be running Ubuntu 16.04LTS Desktop

# Setup Process

Install Ubuntu 16.04LTS Desktop x86_64 edition, without your GPUs plugged in. After installation shutdown, connect your GPUs, and power on again

Run `scripts/install.sh` which will install nvidia drivers, as well as upgrade a package needed to run ethminer. It will also copy the scripts needed to monitor, and manage and ethereum miner, as well as pull stats from said ethereum miner. A script to pull stats from the nvidia gpus is also copied

Run `scripts/zabbix_agent_install.sh`  which will install zabbix agent. 
Copy `zabbix/userparameter_gpu.conf` to `/etc/zabbix/zabbix_agentd.d` and run `systemctl restart zabbix-agent`
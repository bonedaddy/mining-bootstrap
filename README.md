# Mining Bootstrap

In this repository you'll find a collection of docs, scripts, and various programs designed to assist mining farm operators.
Scripts vary from ones to monitor your mining rigs, all the way up to report generation for the tax man.


# Report Generation

Currently reports are supported only for miningpoolhub. The idea behind the reports package is that every 24 hours, you can receive an email detailing how much eth was mined in a 24 hour period, and the USD/CAD evaluation of the mined ethereum. 

# Rig Setup Process

Install Ubuntu 16.04LTS Desktop x86_64 edition, without your GPUs plugged in. After installation shutdown, connect your GPUs, and power on again

Run `scripts/install.sh` which will install nvidia drivers, as well as upgrade a package needed to run ethminer. It will also copy the scripts needed to monitor, and manage and ethereum miner, as well as pull stats from said ethereum miner. A script to pull stats from the nvidia gpus is also copied

Run `scripts/zabbix_agent_install.sh`  which will install zabbix agent. 
Copy `zabbix/userparameter_gpu.conf` to `/etc/zabbix/zabbix_agentd.d` and run `systemctl restart zabbix-agent`

# Ethminer troubleshooting

https://github.com/ethereum-mining/ethminer/issues/314
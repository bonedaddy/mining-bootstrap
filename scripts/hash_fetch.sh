#! /bin/bash

echo '{"id":17,"jsonrpc":"2.0","method":"miner_getstat1"}' | netcat localhost 3000 | awk '{print $1}' | awk -F ':' '{print $4}' | awk -F ',' '{print $3}' | awk -F ';' '{print $1}' | sed 's/\"//g' | cut -b1-3

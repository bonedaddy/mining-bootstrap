#! /bin/bash

USERNAME="......"
WORKER_NAME=$(hostname)
PORT="20535"
URL="stratum2+tcp://$USERNAME.$WORKER_NAME:x@us-east.ethash-hub.miningpoolhub.com:$PORT"

/boot_scripts/bin/ethminer -U -P "$URL" --api-port -6767 2>&1 | tee --append /var/log/ethminer.log
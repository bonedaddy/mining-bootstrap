#! /bin/bash

USERNAME="......"
WORKER_NAME=$(hostname)
PORT="20535"
URL="stratum2+tcp://$USERNAME.$WORKER_NAME:x@us-east.ethash-hub.miningpoolhub.com:$PORT"
API="no"

if [[ "$API" == "yes" ]]; then
    ethminer -U -P "$URL" --api-port -6767 2>&1 | tee --append "$HOME/ethminer.log"
else
    ethminer -U -P "$URL" 2>&1 | tee --append "$HOME/ethminer.log"
fi

#! /bin/bash

QUERY_MODE="$1"
SELECTED_GPU_NUMBER="$2"
GPU_NUMBERS=$(nvidia-smi --list-gpus | awk '{print $2}' | tr -d ':')
COUNT=0

if [[ "$SELECTED_GPU_NUMBER" == "" ]]; then
    echo "[ERROR] No valid gpu number provided"
    exit 1
fi

# if this never exits, then a non valid gpu number was passed
while true; do
    if [[ "$SELECTED_GPU_NUMBER" -eq "$COUNT" ]]; then
        break
    fi
    let COUNT++
done


get_temp() {
    TEMP=$(nvidia-smi -q -i "$1" | grep -i "GPU Current Temp" | awk -F ':' '{print $2}' | awk '{print $1}')
    echo "$TEMP"
}


case "$QUERY_MODE" in 

    # query is for GPU temp
    temp) 
        get_temp "$SELECTED_GPU_NUMBER"
        ;;

esac            

#! /bin/bash


# Used primarily  by zabbix agent to parse nvidia-smi output into a consumable format

# particular query to run
QUERY_MODE="$1"
# the gpu to pull metrics from
SELECTED_GPU_NUMBER="$2"
# cleans up output from list-gpu into an iterable format
GPU_NUMBERS=$(nvidia-smi --list-gpus | awk '{print $2}' | tr -d ':')
COUNT=0

# make sure that a non-empty parameter was given
if [[ "$SELECTED_GPU_NUMBER" == "" ]]; then
    echo "[ERROR] No valid gpu number provided"
    exit 1
fi

# will loop for 100 counts unless a valid GPU number is given
# TODO: this needs to be cleaned up to ensure that the number given is in a valid range
# from that provided by GPU_NUMBERS
while true; do
    if [[ "$SELECTED_GPU_NUMBER" -eq "$COUNT" ]]; then
        break
    fi
    let COUNT++
done

# function used to retrieve the temperature of a gpu
get_temp() {
    TEMP=$(nvidia-smi -q -i "$1" | grep -i "GPU Current Temp" | awk -F ':' '{print $2}' | awk '{print $1}')
    echo "$TEMP"
}

# function used to retrieve the power draw of a gpu
get_power_draw() {
    DRAW=$(nvidia-smi -q -i "$1" | grep -i "Power Draw" | awk -F ':' '{print $2}' | awk '{print $1}')
    echo "$DRAW"
}

case "$QUERY_MODE" in 

    # query is for GPU temp
    temp) 
        get_temp "$SELECTED_GPU_NUMBER"
        ;;

    draw)
        get_power_draw "$SELECTED_GPU_NUMBER"
        ;;
    *)
        echo "Invalid invocation"
        echo "./nvidia_smi_parser.sh [temp | draw] <gpu-number>"
        exit 1
        ;;
esac            

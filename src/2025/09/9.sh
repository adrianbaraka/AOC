#!/usr/bin/env bash
input="input.txt"

abs(){
    #echo "called with $1 $2" >&2
    local n1=$1
    local n2=$2
    local diff
    diff=$((n1-n2))
    if ((diff < 0)); then 
        diff=$((-diff))
    fi
    ((diff++))
    echo "$diff"
}


area(){
    local l1=$1
    local w1=$2
    local l2=$3
    local w2=$4

    local l w
    l=$(abs "$l1" "$l2")
    w=$(abs "$w1" "$w2")

    local ar=$((l*w))
    echo "$ar"
}

# area 2 5 11 1

part1(){
    local data len max_area num l1 w1 i j num2 l2 w2 narea iter
    mapfile -t data < "$input"
    len="${#data[@]}"
    max_area=0
    iter=0
    for ((i=0;i<len-1;i++)); do
        printf "\r %d/%d " "$iter" "$len"
        ((iter++)) 
        num="${data[$i]}"
        IFS=, read -r l1 w1 <<< "$num"
        #echo "Starting: $l1 $w1"
        for ((j=i+1;j<len;j++)); do 
            num2="${data[$j]}"
            IFS=, read -r l2 w2 <<< "$num2"
            #echo "$l2 $w2"
            narea=$(area "$l1" "$w1" "$l2" "$w2")
            if ((narea > max_area)); then 
                max_area=$narea
            fi 
        done
    done
    echo
    echo "$max_area"
}
part1
#!/usr/bin/env bash

load(){
    # both data and len are global
    read -r data
    len="${#data}"
}

part1(){
    local start i char 
    start=0
    for ((i=0;i<len;i++)); do
        char="${data:i:1}"
        if [[ $char == '(' ]]; then
            ((start++))
        else
            ((start--))
        fi
    done

    echo "Part 1: $start"
}
# part1

part2(){
    local start i char
    start=0
    for ((i=0;i<len;i++)); do
        char="${data:i:1}"
        if [[ $char == '(' ]]; then
            ((start++))
        else
            ((start--))
        fi

        if ((start == -1)); then
            # basement reached break
            echo "Part 2: $((i+1))"
            break
        fi
    done
}
#part2

both(){
    local start i char p2
    load
    start=0
    for ((i=0;i<len;i++)); do
        char="${data:i:1}"
        if [[ $char == '(' ]]; then
            ((start++))
        else
            ((start--))
        fi

        if ((start == -1)) && [[ -z "$p2" ]]; then
            p2=$((i+1))
        fi
    done
    echo "Part 1: $start"
    echo "Part 2: $p2"
}

# load
both
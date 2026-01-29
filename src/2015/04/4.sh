#!/usr/bin/env bash

key='iwrupvqb'
num=1
while true; do
    # pad num to key
    #echo -ne "\rTrying $num " >&2
    key="${key}$num"
    md5="$(printf "%s" "$key" | md5sum)"
    #echo "$md5"

    first5="${md5:0:5}"

    if [[ $first5 == '00000' ]]; then
        echo "$num"
        break
    fi
    ((num++))

done

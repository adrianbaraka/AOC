#!/usr/bin/env bash

source vardump

# https://youtu.be/uCsD3ZGzMgE?si=Sh1UPNvC3UWj2NU1

num=3005290
binary=$(echo "obase=2;$num" | bc)

new="${binary:1}1"


echo Part 1: "$((2#$new))"

part2() {
    arr=()
    num=5
    i=1
    while (( num > 1 )); do
        ((val = num ))
        while true; do
            if [[ -n ${arr[$i]} ]]; then
                (( i ++ ))
            else
                break
            fi
        done
        (( half = val / 2 ))
        echo half + i  $((half + i)) i = $i
        arr[ half + i]=0
        (( num -- ))
        (( i++ ))
    done

    vardump -v arr

}

opposite() {
    local index total
    index=$1
    total=$2
    (( res = (index + (total/2)) % total )) # res is global
    ((res == 0)) && (( res = total ))
    ((res--)) # index in array
    echo "index = $index, total= $total, res = $res" 
}


testt() {
    arr=()
    tot=3005290
    for (( i=0; i<tot; i++ )); do
        arr+=("$i")
    done
    total=tot
    curr=1
    #vardump arr
    while (( curr <= tot )); do
        opposite $curr $total
        arr[res]=nil
        new_arr=()
        for (( i=0; i<total; i++ )); do
            [[ "${arr[i]}" != nil ]] && new_arr+=("${arr[i]}")
        done
        arr=()
        arr+=("${new_arr[@]}")
        ((total--))
        ((curr++))
        #vardump arr
    done
    vardump arr
}

testt


part22() {
    arr=()
    total=5
    num=5
    curr=1
    while (( curr <= total )); do
        opposite $curr $num

        ((num--))
        ((curr++))
        while true; do
            if [[ -z "${arr[curr]}" ]]; then
                break
            else
                ((curr++))

            fi
        done
        echo $curr
    done
     vardump arr
}

# part22

#!/usr/bin/env bash

source vardump

get_factors(){
    local num=$1
    factors=()

    for (( i=1; i<=num/2; i++ )); do
        if (( num % i == 0 )); then
            factors+=("$i")
            ((  divisor = num / i))
            (( divisor != i )) && factors+=("$divisor")
        fi
    done

    #vardump factors
}

part1(){
    local num=$1
    #factors=()
    total=0
    ((upto = num / 2))
    for (( i=1; i<=upto; i++ )); do
        printf "\r%d/%d " "$i" "$upto"
        if (( num % i == 0 )); then
            (( total += i * 10 ))
            ((  divisor = num / i))
            (( divisor != i )) && (( total += divisor * 10 ))
        fi
    done
    echo
    echo $total

    #vardump factors
}

part1 36000000
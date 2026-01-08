#!/usr/bin/env bash

part1(){
    local tot mem line len s0 val len_val
    tot=0
    mem=0
    while read -r line; do
        len="${#line}"
        ((tot+=len))
        s0="${line:1:((len-2))}"
        # parameter expansion expanding all backslashes
        val="${s0@E}"
        #echo "$val"
        len_val="${#val}"
        ((mem+=len_val))
    done

    echo $((tot-mem))
}
# part1

part2(){
    local tot mem line len s0 val len_val
    tot=0
    mem=0
    while read -r line; do
        len="${#line}"
        ((tot+=len))
        s0="${line//[^\"\\]/}"
        #echo "here $s0"
        # increase the extra chars
        ((mem+="${#s0}"))
        val="${line@Q}"
        #echo "$val"
        len_val="${#val}"
        ((mem+=len_val))
    done
    echo "Tot: $tot"
    echo "Mem: $mem"
    echo $((mem-tot))
}

part2

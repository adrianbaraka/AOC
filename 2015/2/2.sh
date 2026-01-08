#!/usr/bin/env bash

part1(){
    local total l w h s1 s2 s3 smallest i req
    total=0
    while IFS='x' read -r l w h; do
        # multiply each by 2 when adding
        s1=$((l*w))
        s2=$((w*h))
        s3=$((h*l))

        smallest=$s1
        for i in $s2 $s3; do
            if ((i < smallest)); then
                smallest=$i
            fi
        done
        req=$(( (2*s1) + (2*s2) + (2*s3) + smallest ))
        ((total+=req))
    done
    echo "Part 1: $total"
}
part1

part2(){
    local total l w h data s1 s2 bow wrapper req
    total=0
    while IFS='x' read -r l w h; do
        # get the 2 smallest sides 
        mapfile -t data <<< "$(printf "%s\n" "$l" "$w" "$h" | sort -n)"
        s1="${data[0]}"
        s2="${data[1]}"

        bow=$((l * w * h))
        wrapper=$((2 * (s1 + s2)))
        req=$((bow + wrapper))
        ((total+=req))
    done < /dev/stdin
    echo "Part 2: $total"
}
# part2


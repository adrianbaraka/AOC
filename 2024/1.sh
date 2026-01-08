#!/usr/bin/env bash

if [[ -t 0 ]]; then
    # stdin is a terminal → no pipe given → use file
    input="1.txt"
else
    # stdin is NOT a terminal → data piped in → read from stdin
    input="/dev/stdin"
fi

list1=()
list2=()

while read -r num1 num2; do
    list1+=("$num1")
    list2+=("$num2")
done<"$input"

# sort the arrays
mapfile -t slist1 < <(printf "%s\n" "${list1[@]}" | sort)
mapfile -t slist2 < <(printf "%s\n" "${list2[@]}" | sort)


# get len
len=${#slist1[@]}

part1(){
    local i s1 s2 dist distance
    distance=0
    for ((i=0; i<len; i++)); do
        s1="${slist1[i]}"
        s2="${slist2[i]}"
        #echo "here s1=$s1 - s2=$s2"
        dist=$(( s1-s2 ))
        if ((dist < 0)); then
            dist=$((-dist))
        fi
        ((distance+=dist))
    done

    echo $distance
}

part2(){
    local num tot total
    total=0
    for num in "${slist1[@]}"; do
        tot=$(printf "%s\n" "${slist2[@]}" | grep -c "$num")
        #echo "Num=$num Total=$tot"
        similarity_score=$((num * tot))
        ((total += similarity_score))
    done
    echo $total
}
part2
#!/usr/bin/env bash

dist(){
    local speed time rest n times rem
    speed=$1
    time=$2
    rest=$3
    n=$4
    # global
    distance=0

    (( times = n / (time+rest) ))

    (( rem = (n%(time+rest)) ))

    if (( rem <= time )); then
        (( distance+= (rem * speed) ))
    else
        (( distance += (speed * time) ))
    fi

    (( distance += times * speed * time ))

    #echo "$dis"
}

extract(){
    local line regex
    line=$1

    regex='^([a-zA-Z]+) can fly ([0-9]+) km\/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.'

    if [[ $line =~ $regex ]]; then
        name="${BASH_REMATCH[1]}"
        speed="${BASH_REMATCH[2]}"
        duration="${BASH_REMATCH[3]}"
        rest="${BASH_REMATCH[4]}"

        #echo "$name $speed $duration $rest"
    else
        echo "Unmatched regex $line" >&2
        exit 1
    fi


}

part1(){
    local reindeer max owner
    max=0
    owner=""

    for reindeer in "${reindeers[@]}"; do
        extract "$reindeer"
        dist "$speed" "$duration" "$rest" "$finish"
        if (( distance > max )); then 
            ((max=distance))
            owner=$name
        fi
    done

    echo -e "\E[92mPart 1: Distance: $max, Reindeer: $owner\n\E[0m"
}

# part1

part2() {
    local final info reindeer i max deer j speed duration rest deer k
    declare -A final
    declare -A info


    for reindeer in "${reindeers[@]}"; do
        extract "$reindeer"
        info[$name]="$speed $duration $rest"
        final[$name]=0
    done

    for ((i=1; i<=finish; i++)); do
        max=0
        deer=""
        # loop over all elements in info, get dist travelled at i for each, max dist increment name
        for j in "${!info[@]}"; do 
            read -r speed duration rest <<< "${info[$j]}"
            dist "$speed" "$duration" "$rest" "$i" 
            if ((distance > max)); then 
                max=$distance
                deer="$j"
            elif ((distance == max)); then
                (( final["$j"]++ ))
            fi
        done
        (( final["$deer"]++ ))
    done

    winner=("" 0)
    for k in "${!final[@]}"; do
        #echo "$k = ${final[$k]}"
        if (( "${final[$k]}" > "${winner[1]}" )); then
            winner=("$k" "${final[$k]}")
        fi
    done 

    printf "\E[92mPart 2: Winner: "
    printf "%s %s" "${winner[@]}"
    printf "\E[0m\n"
    
}


main() {
    mapfile -t reindeers
    finish=2503
    part1
    part2
}

main
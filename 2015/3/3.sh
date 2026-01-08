#!/usr/bin/env bash

load(){
    read -r line < /dev/stdin
    len="${#line}"
}

in_array(){
    local -n arr=$1
    local search=$2
    local item

    for item in "${arr[@]}"; do
        [[ "$search" == "$item" ]] && return 0
    done

    return 1
}

col(){
    local str change r c
    str=$1
    change=$2

    # extract the column r,c
    IFS=, read -r r c <<< "$str"

    ((c+=change))
    echo "$r,$c"
}

row(){
    local str change r c
    str=$1
    change=$2

    # extract the column r,c
    IFS=, read -r r c <<< "$str"

    ((r+=change))
    echo "$r,$c"
}


part1(){
    #load
    visited=('0,0')
    curr='0,0'
    count=1
    for ((i=0;i<len;i++));do
        echo -ne  "\r$count/$len " >&2
        ((count++))
        char="${line:i:1}"

        case $char in
            '^') 
                curr=$(col "$curr" 1)
            ;;
            'v') 
                curr=$(col "$curr" -1)
            ;;
            '>') 
                curr=$(row "$curr" 1) 
            ;;
            '<') 
                curr=$(row "$curr" -1)
            ;;
            *)
                echo "Invalid character $char"
                exit 1
            ;;
        esac
        #echo "curr: $curr"
        if ! in_array visited "$curr"; then
            visited+=("$curr")
            # printf "%s " "${visited[@]}"
            # echo
        fi

    done
    echo
    echo "Part 1: ${#visited[@]}"
}

# part1

part2(){
    #load
    visited=('0,0')
    santa='0,0'
    robo_santa='0,0'
    count=1
    for ((i=0;i<len;i++));do
        echo -ne  "\r$count/$len " >&2
        ((count++))
        char="${line:i:1}"

        # even santa odd robo_santa
        if ((i%2 == 0)); then
            curr="$santa"
        else
            curr="$robo_santa"
        fi

        case $char in
            '^') 
                curr=$(col "$curr" 1)
            ;;
            'v') 
                curr=$(col "$curr" -1)
            ;;
            '>') 
                curr=$(row "$curr" 1) 
            ;;
            '<') 
                curr=$(row "$curr" -1)
            ;;
            *)
                echo "Invalid character $char"
                exit 1
            ;;
        esac
        #echo "curr: $curr"
        if ! in_array visited "$curr"; then
            visited+=("$curr")
            # printf "%s " "${visited[@]}"
            # echo
        fi

                # even santa odd rsanta
        if ((i%2 == 0)); then
            santa="$curr"
        else
            robo_santa="$curr"
        fi

    done
    echo
    echo "Part 2: ${#visited[@]}"
}
#part2

load
part1
part2
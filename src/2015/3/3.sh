#!/usr/bin/env bash

load(){
    read -r line
    len="${#line}"
}

in_array(){
    local -n arr=$1
    local search=$2
    local item
    # global bool holding if in array
    in_arr=

    for item in "${arr[@]}"; do
        [[ "$search" == "$item" ]] && in_arr=true && return 0
    done

    in_arr=false
    #return 1
}

col(){
    local str change
    str=$1
    change=$2
    # clear r and c first are global vars
    r=
    c=

    # extract the column r,c
    IFS=, read -r r c <<< "$str"

    ((c+=change))
    #echo "$r,$c"
}

row(){
    local str change
    str=$1
    change=$2
    # clear r and c first are global vars
    r=
    c=

    # extract the column r,c
    IFS=, read -r r c <<< "$str"

    ((r+=change))

    #echo "$r,$c"
}


part1(){
    #load
    visited=('0,0')
    curr='0,0'
    count=1
    for ((i=0;i<len;i++));do
        #echo -ne  "\r$count/$len " >&2
        ((count++))
        char="${line:i:1}"

        case $char in
            '^')
                col "$curr" 1
                printf -v curr "%d,%d" "$r" "$c"
                #curr=$(col "$curr" 1)
            ;;
            'v') 
                col "$curr" -1
                printf -v curr "%d,%d" "$r" "$c"
                #curr=$(col "$curr" -1)
            ;;
            '>') 
                row "$curr" 1
                printf -v curr "%d,%d" "$r" "$c"
                #curr=$(row "$curr" 1) 
            ;;
            '<') 
                row "$curr" -1
                printf -v curr "%d,%d" "$r" "$c"
                #curr=$(row "$curr" -1)
            ;;
            *)
                echo "Invalid character $char"
                exit 1
            ;;
        esac
        #echo "curr: $curr"
        in_array visited "$curr"
        if ! "$in_arr"; then
            visited+=("$curr")
            # printf "%s " "${visited[@]}"
            # echo
        fi

    done
    echo
    echo "Part 1: ${#visited[@]}"
}
load
part1

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

# load
# part1
# part2
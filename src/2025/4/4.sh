#!/usr/bin/env bash

if [[ -t 0 ]]; then
    # stdin is a terminal → no pipe given → use file
    input="4.txt"
else
    # stdin is NOT a terminal → data piped in → read from stdin
    input="/dev/stdin"
fi


accessible(){
    index=$1
    local filled to_check
    filled=0
    to_check=()
    

    local left
    left=$((index - 1))
    if (( (index/columns) == (left/columns) )); then 
        to_check+=("$left")
    fi

    local right
    right=$((index + 1))
    if (( (index/columns) == (right/columns) )); then 
        to_check+=("$right")
    fi

    local up
    up=$((index - columns))
    if (( up >= 0 )); then 
        to_check+=("$up")
    fi
    

    local upleft
    upleft=$((index - columns -1))
    if (( (up/columns) == (upleft/columns) )); then 
        to_check+=("$upleft")
    fi

    local upright
    upright=$((index - columns + 1))
    if (( (up/columns) == (upright/columns) )); then
        to_check+=("$upright")
    fi

    local down
    down=$((index + columns))
    if ((down < len)); then
        to_check+=("$down")
    fi

    local downright
    downright=$((index + columns + 1))
    if (( (down/columns) == (downright/columns) )); then
        to_check+=("$downright")
    fi

    local downleft
    downleft=$((index + columns - 1))
    if (( (down/columns) == (downleft/columns) )); then
        to_check+=("$downleft")
    fi

    local item
    for item in "${to_check[@]}"; do 
        if is_full "$item"; then 
            ((filled++))
        fi
    done


    if ((filled < 4)); then
        return 0
    else
        return 1
    fi


}
# function to check if the element at passed index is empty
is_full(){
        local i=$1
        #echo "$len"
        #echo "Compared ${data[$i]} index: $i"
        if (( i < 0 )) || ((i >= len)); then
            #echo "Empty Out of range"
            return 1
        fi
        if [[ "${data[$i]}" = '@' ]]; then 
            #echo "full"
            return 0
        else
            #echo "empty"
            return 1
        fi
}


part1(){
    local a char
    data=()

    # used by function data: columns len

    # load the info in the data array
    while read -r line; do
        # echo "$line"
        columns="${#line}"
        for ((a=0; a<columns; a++)); do
            char="${line:a:1}"
            data+=("$char")
        done
    done < $input

    local b total
    len="${#data[@]}"
    #echo "Len = $len"
    total=0

    for((b=0; b<len; b++)); do
        #var="${data[$b]}"
        if [[ "${data[$b]}" = '@' ]] && accessible "$b"; then
            #echo "Index free: $b"
            ((total++))
        fi
    done

    echo "$total"
    #printf "%s " "${data[@]}"
}
#part1

part2(){
    local a char
    data=()

    # used by function data: columns len

    # load the info in the data array
    while read -r line; do
        # echo "$line"
        columns="${#line}"
        for ((a=0; a<columns; a++)); do
            char="${line:a:1}"
            data+=("$char")
        done
    done < $input

    # begin loop here
    local b total to_rm total_rm
    total_rm=0
    while true; do
        to_rm=()
        len="${#data[@]}"
        #echo "Len = $len"
        total=0

        for((b=0; b<len; b++)); do
            #var="${data[$b]}"
            if [[ "${data[$b]}" = '@' ]] && accessible "$b"; then
                #echo "Index free: $b"
                to_rm+=("$b")
                ((total++))
            fi
        done

        #echo "$total"

        ((total_rm+=total))


        local len_to_rm c
        len_to_rm="${#to_rm}"
        if ((len_to_rm > 0)); then 
            # replace all @ with .
            for c in "${to_rm[@]}"; do 
                data["$c"]='.'
            done
        else
            break
        fi
    done

    echo "$total_rm"

    #printf "%s " "${data[@]}"
}

part2


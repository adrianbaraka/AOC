#!/usr/bin/env bash

if [[ -t 0 ]]; then
    # stdin is a terminal → no pipe given → use file
    input="input.txt"
else
    # stdin is NOT a terminal → data piped in → read from stdin
    input="/dev/stdin"
fi

loader(){
    local line a char
    # global values arr, col_len, arr_len, start_index
    arr=()
    while read -ra line; do 
        col_len="${#line}"
        for((a=0; a<col_len; a++)); do 
            char="${line:$a:1}"
            if [[ $char = 'S' ]]; then 
                start_index="$a"
            fi
            arr+=("$char")
        done
    done < "$input"

    arr_len="${#arr[@]}"

    #printf "%s" "${arr[@]}"
}
part1(){
    loader
    to_check=("$start_index")
    split=0
    while true; do 
        next_index="${to_check[-1]}"
        next_index=$((next_index + col_len))
        if ((next_index > arr_len )); then 
            echo "breaking"
            break
        fi
        echo "$next_index"
        val="${arr[$next_index]}"
        if [[ $val = '^' ]] && [[ $val != '#' ]]; then  # already checked
            index1=$((next_index - 1))
            index2=$((next_index + 1 ))
                        # to the left
            if (((next_index / col_len) == (index1 / col_len))); then 
                echo "here -1"
                to_check+=("$index1")
                arr[index1]='#'
                ((split++))
            fi 

            if (((next_index / col_len) == (index2 / col_len))); then 
                echo "here +1"
                to_check+=("$index2")
                arr[index2]='#'
                ((split++))
            fi 
            break
        else
            to_check+=("$next_index")
        fi
    done
    printf "%s " "${to_check[@]}"
    echo -e  "\nSplit: $split"
    #echo "Len: ${#to_check[@]}"
    
    #printf "%s " "${arr[@]}"
    #done
}
part1


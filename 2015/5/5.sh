#!/usr/bin/env bash

cond1(){
    # at least 3 vowels
    local str new_str len
    str=$1
    # parameter expansion removing all consonants
    new_str="${str//[^aeiou]/}"
    len="${#new_str}"

    (( len < 3 )) && return 1
    
    return 0

}

cond2(){
    local str len i char next_char
    str=$1
    len="${#str}"
    for((i=0; i<len-1; i++)); do
        char="${str:i:1}"
        next_char="${str:((i+1)):1}"

        [[ "$char" == "$next_char" ]] && return 0 # if same the condition is met

    done

    return 1
}

cond3(){
    local str regex
    str=$1
    regex='ab|cd|pq|xy'

    [[ "$str" =~ $regex ]] && return 1 # if anything matches return 1

    return 0
}

cond4(){
    local str len i char new_str
    str=$1
    len="${#str}"
    # take 2 remove from string and regex match for removed string
    for ((i=0; i<len-1; i++)); do
        char="${str:i:2}"

        # new string replace first occurrence of char with -
        new_str="${str/$char/-}"
        if [[ $new_str =~ $char ]]; then
            return 0
        fi

    done

    return 1
}

cond5(){
    local str len i char next_char
    str=$1
    len="${#str}"
    for ((i=0; i<len-2; i++)); do
        char="${str:i:1}"
        next_char="${str:((i+2)):1}"

        if [[ "$char" == "$next_char" ]]; then
            return 0
        fi
    done

    return 1
}

main(){
    local line nice nice2
    nice=0
    nice2=0
    while read -r line; do
        #echo "$line"
        if cond1 "$line" && cond2 "$line" && cond3 "$line"; then
            ((nice++))
        fi

        if cond4 "$line" && cond5 "$line"; then
            #echo "$line"
            ((nice2++))
        fi
    done

    echo "Part 1: $nice"
    echo "Part 2: $nice2"
}

main
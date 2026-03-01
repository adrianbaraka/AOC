#!/usr/bin/env bash
# shellcheck disable=SC1091
source vardump

input=/dev/stdin
if (( $# > 0 )) && [[ $1 != "-" ]]; then
    input=$1
fi

sortword() {
    local word len wa i char sq arr
    word=$1
    len="${#word}"
    wa=()
    for (( i=0; i < len; i++ )); do
        char="${word:i:1}"
        wa[i]="$char"
    done

    sq=$(printf "%s\n" "${wa[@]}" | sort)

    mapfile -t arr <<< "$sq"

    sortedWord=""
    printf -v sortedWord "%s" "${arr[@]}"
}

isValid() {
    local line words len i word j w2 
    line=$1
    read -ra words <<< "$line"
    len="${#words[@]}"
    for (( i = 0; i < len; i++ )); do
        word="${words[i]}"
        for (( j = i+1; j < len; j++ )); do
            w2="${words[j]}"
            if [[ "$word" == "$w2" ]]; then
                return 1
            fi

        done
    done
    return 0
}

# this implementation is slow due to the external call to sort
isValid2() {
    local line words len i word j w2 
    line=$1
    read -ra words <<< "$line"
    len="${#words[@]}"
    for (( i = 0; i < len; i++ )); do
        word="${words[i]}"
        sortword "$word"
        sw1="$sortedWord"

        for (( j = i+1; j < len; j++ )); do
            w2="${words[j]}"

            sortword "$w2"
            sw2="$sortedWord"

            if [[ "$sw1" == "$sw2" ]]; then
                return 1
            fi

        done
    done
    return 0
}

isValid21() {
    local line words len i word j w2 k lenBefore lenAfter w1
    line=$1
    read -ra words <<< "$line"
    len="${#words[@]}"
    for (( i = 0; i < len; i++ )); do
        word="${words[i]}"
        lenw="${#word}"
        for (( j = i+1; j < len; j++ )); do
            w2="${words[j]}"
            w1="$word"
            #echo "$w2"
            # loop through word removing every letter in w1 from w2 
            # if resulting string is empty it is not valid
            for (( k = 0; k < lenw; k++)); do
                char="${word:k:1}"

                lenBefore="${#w2}"
                w2="${w2/$char/}"
                lenAfter="${#w2}"

                # if a letter was removed from w2 also remove from w1
                (( lenBefore != lenAfter )) && w1="${w1/$char/}"
                # echo "w2 $w2"
                # echo "w1 $w1"
            done
            if [[ "$w2" == "" && "$w1" == "" ]]; then
                #echo "$word"
                return 1
            fi

        done
    done
    return 0
}

main() {
    valid=0
    #valid2=0
    valid21=0
    idx=1
    while read -r line; do
        printf "\rLine %d. " "$idx"
        (( idx++ ))

        if isValid "$line"; then
            (( valid++ ))
        fi
        # slow
        # if isValid2 "$line"; then
        #     echo "[Old] $line"
        #     (( valid2++ ))
        # fi

        if isValid21 "$line"; then
            #echo "[New] $line"
            (( valid21++ ))
        fi

    done < "$input"
    echo
    echo "Part 1: $valid"
    # echo "Part 2: $valid2"
    echo "Part 2: $valid21"
}

main


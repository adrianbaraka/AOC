#!/usr/bin/env bash

declare -A wires
todo=()

in_array(){
    local point
    local item=$1
    for point in "${!wires[@]}"; do
        [[ $point == "$item" ]] && return 0
    done

    return 1
}

get_val(){
    local -n _value=$1
    local val=$2
    # if val is a number set value to that else get the value
    if [[ $val == *[0-9] ]]; then
        _value=$val
    else
        ! in_array "$val" && todo+=("$val") && return 1
        _value="${wires[$val]}"
    fi
}

set_val(){
    local key value
    key=$1
    value=$2
    wires["$key"]="$value"
}



match(){
    local line case1 case2 case3 val wire value w1 op w2 w3 word1 word2
    line=$1
    case1='^([0-9a-z]+) -> ([a-z]+)$' # 123 -> x lx -> a
    case2='^([a-z0-9]+) (AND|OR|LSHIFT|RSHIFT) ([a-z0-9]+) -> ([a-z]+)$' # bn R|LSHIFT 2 -> bo, cj OR|AND cp -> cq
    case3='^NOT ([a-z0-9]+) -> ([a-z]+)$' # NOT y -> i

    if [[ $line =~ $case1 ]]; then
        val="${BASH_REMATCH[1]}"
        wire="${BASH_REMATCH[2]}"

        get_val value "$val" || return 0

        set_val "$wire" "$value"
    
    elif [[ $line =~ $case2 ]]; then
        w1="${BASH_REMATCH[1]}"
        op="${BASH_REMATCH[2]}"
        w2="${BASH_REMATCH[3]}"
        w3="${BASH_REMATCH[4]}"

        word1=
        word2=
        get_val word1 "$w1" || return 0
        get_val word2 "$w2" || return 0


        case $op in
            AND) ((value = "$word1" & "$word2" )) ;;
            OR) ((value = "$word1" | "$word2" )) ;;
            LSHIFT) ((value = "$word1" << "$word2" )) ;;
            RSHIFT) ((value = "$word1" >> "$word2" )) ;;
            *) 
                echo "Unsupported value $op" >&2
                exit 1

            ;;
        esac

        set_val "$w3" "$value"

    elif [[ $line =~ $case3 ]]; then
        w1="${BASH_REMATCH[1]}"
        w2="${BASH_REMATCH[2]}"

        word1=
        get_val word1 "$w1" || return 0

        ((value = ~ "$word1" & 65535))

        set_val "$w2" "$value"
    
    else
        echo "Unsupported pattern: <$line>" >&2
        exit 1
    fi
}

part1(){
    mapfile -t data
    while true; do 
        for li in "${data[@]}"; do
            #echo "$li"
            match "$li"
        done

        a_val="${wires[a]}"
        if [[ -n $a_val ]]; then
            echo "$a_val"
            break
        fi

        if (( "${#todo[@]}" == 0)); then
            todo=()
            echo "Looping back."
            break
        fi
    done
}

# part1

part2(){
    mapfile -t data
    while true; do 
        match "46065 -> b"
        for li in "${data[@]}"; do
            #echo "$li"
            match "$li"
        done

        a_val="${wires[a]}"
        if [[ -n $a_val ]]; then
            echo "$a_val"
            break
        fi

        if (( "${#todo[@]}" != 0)); then
            todo=()
            #echo "Looping back."
        else
            echo "A not found"
            break
        fi
    done
}

part2
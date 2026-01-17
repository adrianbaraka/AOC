#!/usr/bin/env bash

# TODO rewrite using a list and avoiding subshells [144]

# global vars
declare -A lights

extract(){
    local line regex r1 c1 r2 c2 op
    line=$1
    regex='^(.*) ([0-9]{1,3}),([0-9]{1,3}) through ([0-9]{1,3}),([0-9]{1,3})$'

    if [[ ! $line =~ $regex ]]; then
        # early exit
        echo "Did not match line $line" >&2
        exit 1
    fi
    # remove the string turn from op remaining with off or on
    op="${BASH_REMATCH[1]/turn/}"
    r1="${BASH_REMATCH[2]}"
    c1="${BASH_REMATCH[3]}"
    r2="${BASH_REMATCH[4]}"
    c2="${BASH_REMATCH[5]}"

    echo "$op $r1 $c1 $r2 $c2"

}

handle_point(){
    local op row column current
    op=$1
    row=$2
    column=$3
    case "$op" in
        'off')
            lights[$row,$column]=0
        ;;

        'on')
            lights[$row,$column]=1
        ;;

        'toggle')
            current="${lights[$row,$column]}"
            # if 0 or it is unset
            if [[ $current == '0' || -z $current ]]; then
                lights[$row,$column]=1
            else
                lights[$row,$column]=0
            fi
        ;;

        "*")
            echo "Unsupported operation '$op'" >&2
            exit 1
        ;;
    esac
}

part1(){
    local line operation r1 c1 r2 c2 i j lit point count
    count=1
    while read -r line; do

        # hard coding num lines just sad
        printf "\rLine %d/300 " "$count"
        ((count++))

        read -r operation r1 c1 r2 c2 <<< "$(extract "$line")"
        #echo "<$operation> $r1 $c1 $r2 $c2"
        # generate all spots to be changed
        for ((i=r1; i<=r2; i++)); do
            for ((j=c1; j<=c2; j++)); do
                handle_point "$operation" $i $j
            done
        done
    done

    lit=0
    for point in "${lights[@]}"; do
        ((point == 1)) && ((lit++))
    done

    echo
    echo "$lit"

}

#part1

handle_light(){
    local row column change value
    row=$1
    column=$2
    change=$3

    # get the value at that point
    value="${lights[$row,$column]}"

    # if unset set it to 0
    [[ -z $value ]] && lights[$row,$column]=0

    ((value = value + change))

    # if value < 0 set it to 0
    ((value < 0)) && value=0

    lights[$row,$column]=$value
}
handle_point2(){
    local op row column current
    op=$1
    row=$2
    column=$3
    case "$op" in
        'off')
            handle_light "$row" "$column" -1
        ;;

        'on')
            handle_light "$row" "$column" 1
        ;;

        'toggle')
            handle_light "$row" "$column" 2
        ;;

        "*")
            echo "Unsupported operation '$op'" >&2
            exit 1
        ;;
    esac
}

part2(){
    local line operation r1 c1 r2 c2 i j lit point count
    count=0
    while read -r line; do

        # hard coding num lines just sad
        printf "\rLine %d/300 " "$count"
        ((count++))

        read -r operation r1 c1 r2 c2 <<< "$(extract "$line")"
        #echo "<$operation> $r1 $c1 $r2 $c2"
        # generate all spots to be changed
        for ((i=r1; i<=r2; i++)); do
            for ((j=c1; j<=c2; j++)); do
                handle_point2 "$operation" $i $j
            done
        done
    done

    lit=0
    for point in "${lights[@]}"; do
        (( lit+=point ))
    done

    echo
    echo "$lit"

}

part2



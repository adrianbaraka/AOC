#!/usr/bin/env bash

export LC_ALL=C

part1(){
    num=$1
    max=$2
    #((max++))
    times=0
    while true; do
        printf "\r%d/%d " "$((times+1))" "$max"
        new_num=""
        len="${#num}"

        i=0
        while ((i != len)); do
            # get the first char 
            char="${num:i:1}"
            count=1
            ((i++))
            for((j=i; j<len; j++)); do
                next_char="${num:j:1}"
                if ((char == next_char)); then
                    ((count++))
                    ((i++))
                else
                    break
                fi
            done

            new_num+="${count}${char}"
        done

        num="${new_num}"

        ((times++))
        if ((times == 40)); then
            printf "\n Part 1: %d\n" "${#num}"
        fi
        if (( times == max)); then
            printf "\n Part 2: %d\n" "${#num}"
            break
        fi
    done
}

echo -e "Warning Using bash in this problem is terribly slow.\nRather use the python version (same algorithm) or the go one(TODO)" >&2
part1 1321131112 50


# different approach still slower probably slower than the first
part11(){
    local num curr_digit len
    num=$1    
    times=0
    max=$2

    while true; do
        printf "\r%s/%s " "$((times+1))" "$max"
        # get the first digit
        len="${#num}"
        start=0
        new_s=""
        #echo "<$num>"
        while true; do 
            #echo "num: $num"
            curr_digit="${num:0:1}"
            #echo "Curr digit: $curr_digit"
            if [[ $num =~ ^$curr_digit+ ]]; then
                occurr="${#BASH_REMATCH[0]}"
                #echo "occurrences: $occurr"
            else
                echo "not matched <$num> ~= <^$curr_digit+>" >&2
                exit 1
            fi

            new_s+="${occurr}${curr_digit}"
            ((start+=occurr))
            # echo "start: $start"
            num="${num:$occurr}"
            # echo "new num: $num"
            # echo "new_s: $new_s"
            # echo

            ((start==len)) &&  num=$new_s && break
        done
        ((times++))
        ((times == max )) && echo -e "\n${#num}" && break
    done

}
# part11 1 5
# part11 1321131112 30



#!/usr/bin/env bash

source library.sh

is_safe(){
    local -n arr=$1
    local i
    local len="${#arr[@]}"

    local dir=inc
    if ((arr[0] > arr[1])); then 
        dir=desc
    fi

    # handle inreasing
    if [[ $dir == "inc" ]]; then 
        for((i=0; i<len-1; i++)); do
            local curr next
            curr="${arr[$i]}"
            next="${arr[((i+1))]}"
            if (((next-curr) < 1 || (next-curr) > 3)); then
                #echo "invalid" 
                return 1
            fi
            #echo "curr: $curr next: $next"
        done
        #echo valid
        return 0
    fi

    # handle decreasing
    if [[ $dir == "desc" ]]; then 
        for((i=0; i<len-1; i++)); do
            local curr next
            curr="${arr[$i]}"
            next="${arr[((i+1))]}"
            if (((curr-next) < 1 || (curr-next) > 3)); then
                #echo "invalid" 
                return 1
            fi
            #echo "curr: $curr next: $next"
        done
        #echo valid
        return 0
    fi

}

is_safe2(){
    local -n arr=$1
    local i

    local reset=0
    local count=0
    #local bad_index
    while true; do 
        echo "Count: $count" >&2
        echo "Reset: $reset" >&2
        ((count++))
        ((reset > 1)) && return 1

        local len="${#arr[@]}"
        printf "Starting.. " >&2
        printf "%s " "${arr[@]}"
        echo

        local dir=inc
        if ((arr[0] > arr[1])); then 
            dir=desc
        fi
        echo "Direction. $dir" >&2

        # handle inreasing
        if [[ $dir == "inc" ]]; then 
            for((i=0; i<len-1; i++)); do
                local curr next
                curr="${arr[$i]}"
                next="${arr[((i+1))]}"
                echo "curr: $curr next: $next" >&2
                if (((next-curr) < 1 || (next-curr) > 3)); then
                    # if i is the second last one reset the last one
                    if((i == len-2)); then
                        printf "\t\tUnsetting index incrementing by 1  Before: %d " "$i"
                        ((i++))
                        printf "After %d\n" "$i"
                    fi
                    echo -e "\t\tunsetting index $i" >&2
                    unset "arr[$i]"
                    ((reset++))
                    arr=("${arr[@]}") # refill hole left behind
                    continue 2 # finally used it continues the 2 loop which is the while rather than the for immediate
                fi
            done
            #echo valid
        fi
        # unset that pos

        # handle decreasing
        if [[ $dir == "desc" ]]; then 
            for((i=0; i<len-1; i++)); do
                local curr next
                curr="${arr[$i]}"
                next="${arr[((i+1))]}"
                echo "curr: $curr next: $next" >&2
                if (((curr-next) < 1 || (curr-next) > 3)); then
                    # if i is the second last one reset the last one
                    if((i == len-2)); then
                        printf "\t\tUnsetting index incrementing by 1  Before: %d " "$i"
                        ((i++))
                        printf "After %d\n" "$i"
                    fi
                    echo -e "\t\tunsetting index $i" >&2
                    unset "arr[$i]"
                    ((reset++))
                    arr=("${arr[@]}") # refill hole left behind
                    continue 2
                fi
            done
            #echo valid
        fi
        # if reached here valid
        secho "safe" green >&2
        return 0
    done

}
# myarr=(17 16 17 19 20 23 24) #a > d > a > a...
# myarr3=(17 16 17 15 14 13 12) # a > d > d> d..

myarr=(17 18 16 19 20 23 24) #a > d > a > a...
myarr2=(16 18 17 15 14 13 12) # a > d > d> d..
myarr3=(18 16 17 15 14 13 12) # d > a > d> d..
myarr4=(19 17 18  20 24 25) # d > a > a>a..
is_safe2 myarr4

part1(){
    local line nums safe
    safe=0
    while read -r line; do 
        #echo "$line"
        IFS=' ' read -ra nums <<< "$line"
        is_safe nums && ((safe++))
    done < /dev/stdin

    echo "$safe"
}
# part1

part2(){
    local line nums safe
    safe=0
    while read -r line; do 
        #echo "$line"
        IFS=' ' read -ra nums <<< "$line"
        if is_safe2 nums; then 
            secho "\tsafe line $line" "green" >&2
            ((safe++))
        else
            secho "Unsafe line $line" red
        fi
    done < /dev/stdin

    echo "$safe"
}

#part2
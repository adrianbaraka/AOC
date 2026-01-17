#!/usr/bin/env bash

declare -A target

# populate target
while read -r line; do
    IFS=': ' read -r feature val <<< "$line"
    target[$feature]=$val
done < "target.txt"
regex='Sue [0-9]+: (.*)'


main(){
    current=1
    while read -r sue; do
        if [[ ! $sue =~ $regex ]]; then
            echo "Unsuported regex $sue" >&2
            exit 1
        fi
        # remove all space with parameter expansion
        unwanted1=0
        unwanted2=0
        IFS=',' read -ra features <<<"${BASH_REMATCH[1]//' '/}"
        for feat in "${features[@]}"; do
            IFS=: read -r f v <<< "$feat"
            if (( ${target[$f]} != v )); then
                #echo "Unmatched $f trager: ${target[$f]} != found: $v"
                unwanted1=1
                #break
            fi

            # part 2
            re='trees|cats|pomeranians|goldfish'
            if [[ ! $f =~ $re ]]; then
                if (( ${target[$f]} != v )); then
                    #echo "Unmatched $f trager: ${target[$f]} != found: $v"
                    unwanted2=1
                    break
                fi
            fi

            if [[ $f == 'trees' || $f == 'cats' ]]; then
                if (( v < ${target[$f]} )); then
                    unwanted2=1
                fi
            fi

            if [[ $f == 'pomeranians' || $f == 'goldfish' ]]; then
                if (( v > ${target[$f]} )); then
                    unwanted2=1
                fi
            fi
        done
        if ((unwanted1 == 0)); then
            echo "Part 1: $current"
            #break
        fi

        if ((unwanted2 == 0)); then
            echo "Part 2: $current"
            #break
        fi
        ((current++))
    done < "input.txt"

}

main
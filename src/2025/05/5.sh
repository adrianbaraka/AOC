#!/usr/bin/env bash

if [[ -t 0 ]]; then
    # stdin is a terminal → no pipe given → use file
    input="5.txt"
else
    # stdin is NOT a terminal → data piped in → read from stdin
    input="/dev/stdin"
fi

loader(){
    local ranges line
    ranges=0
    fresh=() # global contains all fresh ids
    to_check=() # global contains all ids to check
    # load data
    while read -r line; do
        if [[ $line = '' ]]; then 
            ranges=1
            continue
        fi

        if ((ranges == 0)); then
            fresh+=("$line")

        else
            to_check+=("$line")
        fi
    done < "$input"

}


part1(){
    loader
    local total_fresh fresh_range  upper lower is_fresh b
    total_fresh=0
    b=1
    for tc in "${to_check[@]}"; do
        is_fresh=1
        echo -ne "Checking $b/${#to_check[@]} \r" >&1
        for fresh_range in "${fresh[@]}"; do 
            #echo "Checking in $fresh_range"
            while IFS=- read -r lower upper; do 
                # if there it is in that range
                if ((tc >= lower)) && ((tc <=upper)); then 
                    is_fresh=0
                    #echo "Breaking at $tc=$a"
                    break
                fi
            if ((is_fresh==0)); then 
                #echo "Breaking at reading $lower and $upper"
                break
            fi
            done <<< "$fresh_range"
        if ((is_fresh==0)); then
            #echo "Breaking at checking fresh range $fresh_range" 
            break
        fi
    
        done
    if ((is_fresh==0)); then 
        ((total_fresh++))
    fi
    ((b++))
    done
    echo 
    echo "Total fresh: $total_fresh"

}

part2(){
    loader
    local fresh_range in_range
    total=0
    for fresh_range in "${fresh[@]}"; do 
        #echo "$fresh_range"
        while IFS=- read -r lower upper; do 
            in_range=$(((upper -lower) +1))
            #echo "$in_range"
            ((total+=in_range))
        done <<< "$fresh_range"    
    done
    #echo "$total"
}

part22(){
    loader 
    local a len b  c d
    c=1
    d=1
    while true; do 
        if ((c==0)); then 
            break
        fi
        c=0
        echo "Optimizing loop $d" >&2

        fresh=("${fresh[@]}")
        mapfile -t fresh <<< "$(printf "%s\n" "${fresh[@]}" | sort -t- -k1,1n)"

        len="${#fresh[@]}"

        for((a=0; a<len; a++)); do 
            local range="${fresh[$a]}"
            echo -ne "\tChecking range $range. Index: $a/$len\r" >&2 
            IFS=- read -r lower upper <<< "$range"
            local b=$((a+1))
            if ((b>len)); then 
                continue
            fi
            local next_range="${fresh[$b]}"
            IFS=- read -r min max <<< "$next_range"

            if [[ -z "$min" ]] || [[ -z "$max" ]]; then 
                    #echo "Empty min=$min max=$max"
                    #continue
                    break
            fi
            if ((min<=upper)); then 
                # merge the lower and max
                unset "fresh[$a]"
                unset "fresh[$b]"
                fresh+=("$lower-$max")
                c=1
                break
            fi 
        done
        echo -e "\tFinished loop $d"
        ((d++))
    done

    # reindex
    echo "Finished Optimizing."
    #fresh=("${fresh[@]}")

    local fresh_range in_range total lo up
    total=0
    for fresh_range in "${fresh[@]}"; do 
        echo "$fresh_range"
        while IFS=- read -r lo up; do 
            if ((lower > upper)); then 
                echo "uh oh $lo-$up"
            fi
            in_range=$(((up -lo) +1))
            #echo "$in_range"
            ((total+=in_range))
        done <<< "$fresh_range"    
    done

    echo "$total"
}

part23() {
    loader

    # ---- Local variables ----
    local fresh_range lo up total

    echo "Optimizing..." >&2

    # ---- 1. Sort ranges once by lower bound ----
    mapfile -t fresh <<< "$(printf "%s\n" "${fresh[@]}" | sort -t- -k1,1n)"

    # ---- 2. Merge intervals in one linear pass ----
    local merged=()
    local cur_lo cur_up next_lo next_up i range

    # initialize merge with first range
    IFS=- read -r cur_lo cur_up <<< "${fresh[0]}"

    for ((i=1; i<${#fresh[@]}; i++)); do
        range="${fresh[$i]}"
        IFS=- read -r next_lo next_up <<< "$range"

        if (( next_lo <= cur_up )); then
            # extend current interval
            (( next_up > cur_up )) && cur_up=$next_up
        else
            # finalize current interval
            merged+=("$cur_lo-$cur_up")
            cur_lo=$next_lo
            cur_up=$next_up
        fi
    done

    # add the last merged interval
    merged+=("$cur_lo-$cur_up")

    echo "Finished Optimizing."

    # ---- 3. Print merged ranges and compute total ----
    total=0
    for fresh_range in "${merged[@]}"; do
        echo "$fresh_range"
        IFS=- read -r lo up <<< "$fresh_range"

        if (( lo > up )); then
            echo "uh oh $lo-$up" >&2
        fi

        (( total += (up - lo + 1) ))
    done

    echo "$total"
}


part23
#344771884978261
#part1
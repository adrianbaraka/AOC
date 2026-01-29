#!/usr/bin/env bash

if [[ -t 0 ]]; then
    # stdin is a terminal → no pipe given → use file
    input="3.txt"
else
    # stdin is NOT a terminal → data piped in → read from stdin
    input="/dev/stdin"
fi

max_combo1(){
    local num fdigit max i j len digit
    num=$1
    max="${num:0:2}"
    len="${#num}"

    for ((i=0; i<len; i++)); do
        fdigit="${num:$i:1}"
        for((j=i+1; j<len; j++)); do
            digit="$fdigit${num:$j:1}" # build the next number to check
            # compare against current max
            if ((digit > max)); then
                #echo "$digit > $max new max = $digit" 
                max=$digit
            fi
        done
    done

    echo "$max"
}



final_combo(){
    local num len_num a num_array answer
    num=$1
    len_num="${#num}"
    num_array=()

    # convert the number to an array
    for ((a=0; a<len_num; a++)); do
        num_array+=("${num:a:1}")
    done

    # get max number from right upto -11
    # chop off every number before that. It is the first digit.
    # Repeat until 12 digits or no more left and append the remaining

    answer=()
    local max_num first_index b c answer_len to_reserve
    c=0
    while true; do
        ((c > 100)) && echo "Recursion limit" >&2 && exit 1
        
        max_num=0
        first_index=0
        len_num="${#num}"
        answer_len="${#answer[@]}"
        to_reserve=$((11 - answer_len))
        #to_loop=$((13 - to_reserve))
        #echo "Reserving $to_reserve looping till $to_loop"

        echo -e "\tlen num $len_num"  >&2
        if ((answer_len != 12)); then   
            for ((b=0; b<len_num-to_reserve; b++)) do
                var="${num:b:1}"
                echo "here $var" >&2
                if ((var > max_num)); then 
                    max_num=$var
                    first_index=$b
                fi
            done

            # append that max to answer
            answer+=("$max_num")
            printf "Before chopping num: %s length = %s \n" "$num" "${#num}" >&2
            #chop off everything before

            ((first_index++))
            num="${num:$first_index}"

            printf "After chopping num: %s Length: ${#num} \n" "$num" >&2

            printf "Answer: " >&2
            printf "%s" "${answer[@]}" >&2
            echo " Length ${#answer[@]}" >&2
            


            echo "Max_num: $max_num First Index: $first_index" >&2
            #((c++))
        else
            local d
            d=$(printf "%s" "${answer[@]}")

            echo "$d"
            return 0
        fi
        ((c++))

    done
}


part1(){
    local line max sum
    sum=0
    while read -r line; do
        max=$(max_combo1 "$line")
        ((sum += max))
    done < $input
    echo "$sum"
}
#part1

part2(){
    local line max sum
    sum=0
    while read -r line; do
        max=$(final_combo "$line")
        echo "[Part2] $max" >&2
        ((sum += max))
    done < $input
    echo "$sum"
}

#final_combo 818181911112111
part2


#!/usr/bin/env bash

if [[ -t 0 ]]; then
    # stdin is a terminal → no pipe given → use file
    input="2.txt"
else
    # stdin is NOT a terminal → data piped in → read from stdin
    input="/dev/stdin"
fi

is_valid(){
    local num=$1
    local num_digits
    num_digits="${#num}"
     # odd number of digits a valid number
    if (( num_digits %2 != 0 )); then 
        return 0
    fi

    # get the first half of digits and compare with the second half if same it is not valid
    local first_half="${num:0:((num_digits/2))}"
    local second_half="${num:(((num_digits/2)))}"

    #echo "First half=$first_half Second Half=$second_half"

    if [[ "$first_half" == "$second_half" ]]; then
        return 1
    else
        return 0
    fi
}
#5,529,687-5,587,329

# any number where both start and end has an odd no. of digits is valid no need to check
early_valid(){
    local strt=$1
    local ed=$2
    local lens="${#strt}"
    local lene="${#ed}"

    #echo "start= $strt #=$lens End=$ed #=$lene"

    if ((lens %2 != 0 )) && ((lene %2 != 0)); then
        return 0
    else
        return 1
    fi
}

is_validi(){
    # 0 is valid 1 is not valid
    local num=$1
    local num_digits
    num_digits="${#num}"
    #echo "From is valid 2 num=$num"

    # if num digits is one it is valid
    if ((num_digits == 1)); then
        return 0
    fi
    
    # for 2 digit numbers use the old is valid
    if ((num_digits == 2)); then
        if ! is_valid "$num"; then
            #echo "repeated string $num" 
            return 1
        else
            #echo "valid $num"
            return 0
        fi
    fi
    #echo "$num"
    for ((i=1; i<=(num_digits/2)+1; i++)); do
        test_string="${num:0:$i}"
        #echo "$i. $test_string"

        
        for (( j=0; j<=(num_digits -1); j+=i )); do
            local eq=0
            val="${num:$j:$i}"
            # echo -e "\t teststring: $test_string val: $val j=$j i=$i"
            if [[ $val != "$test_string" ]]; then
                eq=1
                #echo -e "\tIssuing break with val=$val "
                break
            fi
        done
        if ((eq == 0)); then
            #echo "repeated string $num repeated portion: $test_string"
            return 1
        fi
    done
    return 0
    
}

part1(){
    read -r name <"$input" 
    #echo "$name"
    local sum i 
    IFS=','
    sum=0
    for val in $name; do
        while IFS='-' read -r start end; do
            #echo "Start=$start End=$end"
            # check for early validity 
            if early_valid "$start" "$end"; then
                echo "Early valid found skipping range $start-$end"
                break
            fi
            for((i=start; i <= end; i++)); do
                if ! is_valid $i; then
                    echo $i
                    ((sum += i))
                fi
            done
        done <<< "$val"
    done

    echo $sum
}

part2(){
    read -r name <"$input" 
    #echo "$name"
    IFS=','
    sum=0
    for val in $name; do
        while IFS='-' read -r start end; do
            #echo "Start=$start End=$end"
            for((k=start; k <= end; k++)); do
                #echo $i
                if ! is_validi $k; then
                    echo $k
                    ((sum += k))
                fi
            done
        done <<< "$val"
    done

    echo $sum
}

# leaning
learn(){
    IFS=, read -ra elements <"$input"
    for element in "${elements[@]}"; do
        echo "$element"
    done
}

part1
#learn
#part2
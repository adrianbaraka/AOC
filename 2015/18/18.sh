#!/usr/bin/env bash
source vardump
# global vars
declare -a data
columns=0
len=0

input="input.txt"
load_data(){
    local i char
    data=()
    while read -r line; do
        columns="${#line}"
        for ((i=0; i<columns; i++)); do
            char="${line:$i:1}"
            data+=("$char")
        done
    
    done<"$input"
    len="${#data[@]}"
}


get_neighbours() {
    local index=$1 columns=$2 len=$3
    neighbours=()

    local row=$((index / columns))
    local col=$((index % columns))

    # Define relative movements: [row_diff, col_diff]
    # This covers Up, Down, Left, Right, and all 4 Diagonals
    for rd in -1 0 1; do
        for cd in -1 0 1; do
            ((rd == 0 && cd == 0)) && continue # Skip the center (itself)

            local r=$((row + rd))
            local c=$((col + cd))

            # Check if new row/col are within grid boundaries
            if ((r >= 0 && r < (len / columns) && c >= 0 && c < columns)); then
                local n_index=$((r * columns + c))
                # Final check to ensure the index exists in the flat array
                if ((n_index >= 0 && n_index < len)); then
                    neighbours+=("$n_index")
                fi
            fi
        done
    done
}

draw_arr(){
    local i j len
    local cols=$1
    local -n arr
    arr=$2
    len="${#arr[@]}"
    for (( i=0; i<len; i+=cols )); do
        for(( j=i; j<i+cols; j++ )); do
            #printf "%s" "$j"
            printf "%s" "${arr[$j]}"
        done
        echo
    done
}


part1(){
    local iter steps new_data i num_on n curr_state total p
    load_data
    iter=1
    steps=100
    while (( iter <= steps )); do
        printf "\r%s/%s " "$iter" "$steps"
        new_data=()
        # draw_arr "$columns" data
        # echo
        for i in "${!data[@]}"; do
            get_neighbours "$i" "$columns" "$len"
            num_on=0
            for n in "${neighbours[@]}"; do
                [[ ${data[$n]} == '#' ]] && ((num_on++))
            done

            curr_state=${data[$i]}
            if [[ $curr_state == '#' ]]; then
                if (( num_on == 2 || num_on == 3  )); then
                    new_data+=('#')
                else
                    new_data+=('.')
                fi
            else
                if ((num_on == 3)); then 
                    new_data+=('#')
                else
                    new_data+=('.')
                fi
            fi
        done
        data=()
        data+=("${new_data[@]}")

        #draw_arr "$columns" data
        #echo
        (( iter++ ))
    done

    # get total on
    total=0
    for p in "${data[@]}"; do
        [[ $p == '#' ]] && ((total++))
    done
    echo
    echo "$total"
}



part2(){
    local iter steps new_data i num_on n curr_state total p
    load_data
    iter=1
    steps=100

    for i in 0 $((columns-1)) $((len-columns)) $((len-1)); do
        data["$i"]='#'
    done 
    #draw_arr "$columns" data
    while (( iter <= steps )); do
        printf "\r%s/%s " "$iter" "$steps"
        new_data=()
        # draw_arr "$columns" data
        # echo
        for i in "${!data[@]}"; do
            #echo "$i"
            if (( i == 0 || i == columns-1 || i == len-columns || i == len-1 )); then
                #echo -e "\nSkipping i=$i\n"
                new_data["$i"]='#'
                continue
            fi
            get_neighbours "$i" "$columns" "$len"
            num_on=0
            for n in "${neighbours[@]}"; do
                [[ ${data[$n]} == '#' ]] && ((num_on++))
            done

            curr_state=${data[$i]}
            if [[ $curr_state == '#' ]]; then
                if (( num_on == 2 || num_on == 3  )); then
                    new_data+=('#')
                else
                    new_data+=('.')
                fi
            else
                if ((num_on == 3)); then 
                    new_data+=('#')
                else
                    new_data+=('.')
                fi
            fi
        done
        data=()
        data+=("${new_data[@]}")

        #echo
        #draw_arr "$columns" data
        #echo
        (( iter++ ))
    done

    # get total on
    total=0
    for p in "${data[@]}"; do
        [[ $p == '#' ]] && ((total++))
    done
    echo
    echo "$total"
}
part2

#!/usr/bin/env bash

if [[ -t 0 ]]; then
    # stdin is a terminal → no pipe given → use file
    input="input.txt"
else
    # stdin is NOT a terminal → data piped in → read from stdin
    input="/dev/stdin"
fi
mkdir -p temp

loader(){
    local num_columns i col
    num_columns=$(head -n 1 "$input" | wc -w)
    # clear new.txt first
    printf "" > temp/1.txt
    for ((i=1; i<=num_columns; i++)); do
        col=$(awk -v col="$i" '{print $col}' "$input" | tr '\n' ' ')
        echo "$col" > temp/1.txt
    done
    input="new_test.txt"
    echo "Done Transposing the input." >&2
}

#loader
part1(){
    loader
    local a total line operation len answer val
    total=0
    while read -ra line; do 
        operation="${line[-1]}"
        len="${#line[@]}"
        case $operation in
                '*'| '/')
                    answer=1 
                ;;
                '+' | '-')
                    answer=0 
                ;;
        esac
        for ((a=0; a<len-1; a++)); do
            # shellcheck disable=SC2034
            val="${line[$a]}" 
            # shellcheck disable=SC1102
            answer=$((answer "$operation" val))
        done
        ((total+=answer))

    done < "$input"
    echo "$total"
}
part1

pad(){
    local num required diff len_num a to_return
    num=$1
    required=$2
    len_num="${#num}"
    diff=$((required-len_num))
    to_return=()
    if ((required < len_num)); then 
        echo "Error. Cannot truncate number" >&2
        exit 1
    fi

    while ((diff>0)); do 
        to_return+=("_")
        ((diff--))
    done
    to_return+=("$num")
    printf "%s" "${to_return[@]}"
}

loader2(){
    local len a max_digits b c

    while read -ra line; do 
        len="${#line[@]}"
        #echo "len=$len"
        max_digits=0
        for ((a=0; a<len-1; a++)); do 
            len_digit="${#line[$a]}"
            if ((len_digit > max_digits)); then 
                max_digits=$len_digit
            fi
        done
        operation="${line[-1]}"
        # build the new digits
        newline=()
        
        for ((b=max_digits; b>=0; b--)); do # loop equal to max digits
            for ((c=0; c<len-1; c++)); do #loop over all numbers in line extracting the bth digit
                num="${line[$c]}"
                num=$(pad "$num" "$max_digits")
                char="${num:$b:1}"
                if [[ -n "$char" ]] && [[ ! "$char" = '_' ]]; then 
                    newline+=("$char")
                fi
            done
            # add a space after each new number
            newline+=(" ")
        done
        newline+=("$operation")
        printf "%s" "${newline[@]}"
        echo

    done < "$input"
}


# starts here
width(){
    local ref len  a val 
    separators=()
    # get width of each column
    ref=$(tail -n 1 $input | tr ' ' '#')
    len="${#ref}"
 
    for ((a=0; a<len; a++)); do 
        val="${ref:$a:1}"
        if [[ "$val" != '#' ]] && ((a != 0)); then 
            val=$((a-1))
            separators+=("$val")
        fi
    done
    # printf "<%s> " "${separators[@]}"
    echo "Finshed width 1"
}
#width
replace(){
    local original index_to_replace new_char prefix suffix new_string
    original=$1
    index_to_replace=$2
    new_char=$3

    prefix="${original:0:$index_to_replace}"
    suffix="${original:$((index_to_replace + 1))}"

    new_string="${prefix}${new_char}${suffix}"

    echo "$new_string"
}
translate(){
    width
    local nlines a nstring sep
    printf "" > "1.txt"
    nlines=$(wc -l "$input" | cut -f1 -d " ")
    for((a=1; a<=nlines; a++)); do 
        nstring=$(head -n $a $input | tail -n 1)
        for sep in "${separators[@]}"; do 
            nstring=$(replace "$nstring" "$sep" '@' )
        done
        if ((a != nlines)); then 
            nstring=$(tr ' ' '-' <<< "$nstring")
        fi
        echo "$nstring" >> "1.txt"
    done
    input="1.txt"
    echo "Finished 2 translate"
    #cat "temp.txt"

}
#part22

loader22(){
    translate
    local num_columns i col 
    num_columns=$(tail -n 1 "$input" | wc -w)
    echo "here: $num_columns"
    printf "" > 2.txt
    for ((i=1; i<=num_columns; i++)); do
        #col=$(awk -v col="$i" '{print $col}' "$input" | tr '\n' ' ')
        col=$(cut -d '@' -f "$i" < "$input" | tr '\n' ' ')
        echo "$col" >> 2.txt
    done
    input="2.txt"
    echo "Done Transposing the input." >&2
}

loader3(){
    loader22
    local len a max_digits b c
    printf "" > 3.txt

    while read -ra line; do 
        len="${#line[@]}"
        #echo "len=$len"
        max_digits=0
        for ((a=0; a<len-1; a++)); do 
            len_digit="${#line[$a]}"
            if ((len_digit > max_digits)); then 
                max_digits=$len_digit
            fi
        done
        operation="${line[-1]}"
        # build the new digits
        newline=()
        
        for ((b=max_digits; b>=0; b--)); do # loop equal to max digits
            for ((c=0; c<len-1; c++)); do #loop over all numbers in line extracting the bth digit
                num="${line[$c]}"
                num=$(pad "$num" "$max_digits")
                char="${num:$b:1}"
                if [[ -n "$char" ]] && [[ ! "$char" = '_' ]]; then 
                    newline+=("$char")
                fi
            done
            # add a space after each new number
            newline+=(" ")
        done
        newline+=("$operation")
        local nl
        nl=$(printf "%s" "${newline[@]}" | tr '-' ' ')
        echo "$nl" >> 3.txt

    done < "$input"
    input="3.txt"
    echo "done 3 loader 3"
}
part3(){
    loader3
    local a total line operation len answer val
    total=0
    while read -ra line; do 
        #printf "<%s>" "${line[@]}"
        operation="${line[-1]}"
        len="${#line[@]}"
        case $operation in
                '*'| '/')
                    answer=1 
                ;;
                '+' | '-')
                    answer=0 
                ;;
        esac
        for ((a=0; a<len-1; a++)); do
            # shellcheck disable=SC2034
            val="${line[$a]}" 
            # shellcheck disable=SC1102
            answer=$((answer "$operation" val))
        done
        ((total+=answer))

    done < "$input"
    echo "$total"
}


# part3
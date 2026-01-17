#!/usr/bin/env bash
. vardump

setup() {
    local abc char i current next
    declare -gA next_chars
    abc=()
    for char in {a..z}; do
        abc+=("$char")
    done

    for (( i=0; i<=24; i++ )); do
        current="${abc[$i]}"
        next="${abc[$((i+1))]}"

        next_chars["$current"]="$next"
    done

    next_chars[z]=a

    #vardump next_chars

}

to_decimal () {
    to_dec=""
    printf -v to_dec "%d" \'"$1"
}

to_ascii () {
    local hex
    to_asc=""
    printf -v hex "%x" "$1"
    printf -v to_asc "%b" "\x$hex"

}

to_ascii 0
echo "1=> '$to_asc'"

get_next_seq() {
    local prev prev_a add final_a len i carry j val diff k
    prev=$1
    prev_a=()
    add=()
    final_a=()
    len="${#prev}"

    # convert ascii to decimal
    for (( i=0; i<len; i++ )); do
        to_decimal "${prev:i:1}"
        prev_a+=("$to_dec")
        add+=(0)
    done
    add[-1]=1
    carry=0
    # perform the arithmetic
    for (( j=len-1; j>=0; j-- )); do 
        (( val = "${prev_a[$j]}" + "${add[$j]}" + "$carry" ))

        if (( val > 122 )); then
            ((diff = val - 122))
            (( val = 96 + diff ))
            carry=$diff
        else
            carry=0
        fi

        final_a[j]=$val
    done

    # convert back to ascii
    next_seq=""
    if (( carry > 0 )); then
        to_ascii "$(( carry + 96 ))"
        next_seq+="$to_asc"
    fi

    for k in "${final_a[@]}"; do
        to_ascii "$k"
        next_seq+="$to_asc"
    done
    # global val return value = next_seq
}

conditions() {
    local str len dec i j d1 d2 d3 cond13 cond3 k 
    # return value condition1: bool
    str=$1
    len="${#str}"
    dec=()
    for (( i=0; i<len; i++ )); do
        if [[ $i == 'i' || $i == 'o' || $i == 'l' ]]; then
            # condition 2 failed return failed
            return 1
        fi
        to_decimal "${str:i:1}"
        dec+=("$to_dec")
    done

    # cond13 is a variable set to 0 and incremented if either condition 1 or 3 is met
    # if it ends up at 2 both are met return true else false
    cond13=0

    # condition 1
    for (( j=0; j<=len-3; j++ )); do
        # get three digits 
        d1="${dec[$j]}"
        d2="${dec[$((j+1))]}"
        d3="${dec[$((j+2))]}"

        if (( d2-d1 == 1 && d3-d1 == 2 )); then
            # condition 1 met return true
            ((cond13++))
            break
        fi
    done

    # condition 3
    cond3=0
    for (( k=0; k<=len-2; k++ )); do
        # get two digits 
        d1="${dec[$k]}"
        d2="${dec[$((k+1))]}"

        if (( d1 == d2 )); then
            ((cond3++))

            # increment k as it is not needed to check
            ((k++))
        fi

        if ((cond3 >= 2)); then
            # 2 pairs have been found break
            ((cond13++))
            break
        fi
    done

    if (( cond13 == 2 )); then
        return 0
    fi

    return 1

}

part1() {
    pass1="cqjxjnds"
    pass="cqjxxyzz"
    while true; do
        get_next_seq "$pass"
        if (( "${#next_seq}" != "${#pass}" )); then
            echo "Answer too big" >&2
            exit 1
        fi
        #printf "\rChecking %s " "$next_seq"
        if conditions "$next_seq"; then
            echo -e "\nPass: $next_seq"
            break
        else
            pass=$next_seq
        fi
        
    done
}
part1


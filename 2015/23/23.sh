#!/usr/bin/env bash

source ../../utils.sh

main() {
    get_input "$@"
    mapfile -t inst < "$input"
    declare -A registers


    for j in 0 1; do
        i=0
        registers[a]=$j
        registers[b]=0
        len="${#inst[@]}"

        while true; do
            line="${inst[$i]}"
            f3="${line:0:3}"
            case $f3 in
                "hlf")
                    reg="${line: -1}"
                    (( registers[$reg] /= 2 ))
                    #(( register = register / 2 ))
                    (( i++ ))
                ;;
                "tpl")
                    reg="${line: -1}"
                    (( registers[$reg] *= 3 ))
                    #(( register = register * 3 ))
                    (( i++ ))
                ;;
                "inc")
                    reg="${line: -1}"
                    (( registers[$reg]++ ))
                    #(( register = register + 1 ))
                    (( i++ ))
                ;;
                "jmp")
                    read -r _ offset <<< "$line"
                    (( i+= offset))
                ;;
                "jie")
                    IFS=, read -r in offset <<< "$line"
                    reg="${in: -1}"

                    if (( registers[$reg] % 2 == 0 )); then
                        (( i += offset ))
                    else
                        ((i++))
                    fi
                ;;
                "jio")
                    IFS=, read -r in offset <<< "$line"
                    reg="${in: -1}"
                    if (( registers[$reg] == 1 )); then
                        (( i += offset ))
                    else
                        ((i++))
                    fi
                ;;
                *)
                    echo "Instruction '$f3' undefined. Breaking."
                    break
                ;;
            esac

            if (( i >= len || i < 0 )); then
                #echo "Index i out of bounds breaking."
                break
            fi
        done
        echo "Part $((j+1)): ${registers[b]}"
    done
}
main "$@"

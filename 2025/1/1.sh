#!/usr/bin/env bash

if [[ -t 0 ]]; then
    # stdin is a terminal → no pipe given → use file
    input="1.txt"
else
    # stdin is NOT a terminal → data piped in → read from stdin
    input="/dev/stdin"
fi

main(){
    local current=50
    local pass=0

    while read -r line; do
        local action="${line:0:1}"
        local num="${line:1}"
        

        # handle left
        if [[ "$action" = "L" ]]; then
            local new=$(( current - num ))


            if (( new == 0)); then
                # increment pass
                ((pass++))
                ((current = new))
            elif (( new < 0 )); then
                if (( new <= -100)); then
                    ((new = new % 100))
                fi

                if ((new == 0)); then
                    ((pass++))
                    ((current = 0))
                else
                    ((current = new + 100))
                fi
            else
                ((current = new))
            fi
        else
            # handle right
            new=$(( current + num ))

            # if new == 100 increment pass
            if (( new == 100 )); then
                ((pass++))
                ((current = 0))

            elif (( new > 100 )); then
                ((new = new % 100))
                # full rotation increment pass and reset current to 0
                if ((new == 0)); then
                    ((pass++))
                    ((current = 0))
                else
                    ((current = (new % 100 )))
                fi
            else
                (( current = new ))
            fi
        fi

    done < "$input"
    echo "Part 1: $pass"
}

part2(){
    local current=50
    local pass=0
    while read -r line; do
        local action="${line:0:1}"
        local num="${line:1}"

        # handle left
        if [[ "$action" = "L" ]]; then
            local new=$(( current - num ))
            
            # increment pass
            if (( new == 0)); then
                ((pass++))
                ((current = new))
            elif (( new < 0 )); then
                # increment every time it passes 0. If current is not 0, add 1 
                if ((current != 0)); then
                    times=$(( (-new / 100) + 1 ))
                    ((pass += times))
                    #echo "left times=$times line=$line current=$current pass=$pass num=$num"
                else
                    times=$(( -new / 100 ))
                    ((pass += times))
                fi

                if (( new <= -100)); then
                    ((new = new % 100))
                fi

                if ((new == 0)); then
                    #((pass++))
                    ((current = 0))
                else
                    ((current = new + 100))
                fi
            else
                ((current = new))
            fi
        else
            # handle right
            new=$(( current + num ))

            # if new == 100 increment pass
            if (( new == 100 )); then
                ((pass++))
                ((current = 0))

            elif (( new > 100 )); then
                # everytime it touches 0 increment.
                times=$(( new / 100 ))
                ((pass += times))


                ((new = new % 100))
                # full rotation increment pass and reset current to 0
                if ((new == 0)); then
                    #((pass++))
                    ((current = 0))
                else
                    ((current = (new % 100 )))
                fi
            else
                (( current = new ))
            fi
        fi
        #echo -e "\tPass = $pass next_start=$current instruction=$line"
    done < "$input"
    echo "Part two: $pass"
}


main
part2


#!/usr/bin/env bash

get_input_(){
    local OPTIND opt
    local testing
    testing=false
    while getopts 't' opt; do
        case "$opt" in
            t) testing=true;;
            *) return 1 ;;
        esac
    done

    if $testing; then
        input="test.txt"
    else
        input="input.txt"
    fi

    export input
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
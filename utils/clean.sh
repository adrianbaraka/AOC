#!/usr/bin/env bash

removeBinaries() {
    local toDel file type ans
    shopt -s globstar
    toDel=()
    for file in ../**/*; do
        [[ ! -f "$file" ]] && continue # do not get directories or the literal string ./**/*

        type="$(file -b --mime-type "$file")"
        if [[ $type == *'application/x-executable'* ]]; then 
            toDel+=("$file")
        fi
    done

    #no files found
    (( "${#toDel[@]}" == 0 )) && echo "No binaries found." && return

    echo "Found the following ${#toDel[@]} files..."
    printf "'%s'\n" "${toDel[@]}"
    read -rp "Delete the files? y or n  " ans
    
    [[ "$ans" == 'y' || "$ans" == 'Y' ]] && rm -v "${toDel[@]}"
}

cleanPermissions() {
    local file
    shopt -s globstar
    for file in **/*; do
        if [[ -d "$file" ]]; then
            chmod 775 "$file" # directories
            continue
        fi

        if [[ "$file" == *.sh || "$file" == *.py ]]; then
            chmod 774 "$file" #  executable files
            continue
        fi

        chmod 664 "$file" # any other file
    done
}

removeBinaries
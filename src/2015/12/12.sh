#!/usr/bin/env bash

# part1(){
#     read -r line
#     mapfile -t results < <(grep -oE '\-?[0-9]+' <<< "$line")

#     s=0
#     for n in "${results[@]}"; do
#         ((s+=n))
#     done
#     echo "$s"
# }

# part1
part2(){
  cat input.json | jq '
    walk(
      if type == "object" and any(.[]; . == "red") 
      then empty 
      else . 
      end
    ) 
    | [.. | numbers] 
    | add
  '
}
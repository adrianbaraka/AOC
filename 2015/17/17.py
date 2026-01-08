#!/usr/bin/env python3

from itertools import combinations


data =[]
with open("input.txt", "r") as f:
    for line in f:
        data.append(int(line.strip()))

def main():
    tot=0
    cont = {}
    for i in range(len(data)):
        for n in combinations(data, i):
            if sum(n) == 150:
                tot+=1

                # part 2
                if len(n) in cont.keys():
                    cont[len(n)]+=1
                else:
                    cont[len(n)] = 1


    print(f"Part 1: {tot}")
    print(f"Part 2: {cont[min(cont.keys())]}")

main()

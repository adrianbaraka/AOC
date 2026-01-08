#!/usr/bin/env python3

import sys
from itertools import combinations
import math

def load_data() -> list:
    input = "input.txt"
    if len(sys.argv) > 1 and sys.argv[1] == "-t":
        input = "test.txt"

    data = []
    with open(input, "r") as f:
        for line in f:
            data.append(int(line.strip()))
    return data

def get_ideal(groups: int, data: list):
    target = sum(data) / groups

    possible = []
    minn = None
    # get only the possibilities with mean number 
    for i in range(len(data)):
        if minn is not None and i > minn:
            break
        for j in combinations(data, i):
            if sum(j) == target:
                if minn is None:
                    minn = i
                possible.append(j)

    # get the one with the least quantum entanglement
    qe = math.inf
    for a in possible:
        q=1
        for k in a:
            q *= k
        if q < qe:
            qe = q

    return qe

def main():
    data = load_data()
    print(f"Part 1: {get_ideal(3, data)}")
    print(f"Part 2: {get_ideal(4, data)}")

main()

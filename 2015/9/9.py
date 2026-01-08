#!/usr/bin/env python3

from itertools import permutations
import sys
import re
import math

def load_data():
    input = "input.txt"
    if len(sys.argv) > 1 and sys.argv[1] == "-t":
        input = "test.txt"


    cities = {}
    distances = {}

    regex = '^([A-Za-z]+) to ([A-Za-z]+) = ([0-9]+$)'
    re_object = re.compile(regex)
    with open(input, "r") as file:
        for line in file:
            myl = re_object.findall(line)

            cities[myl[0][0]] = None
            cities[myl[0][1]] = None

            # both to and fro

            distances[f"{myl[0][0]},{myl[0][1]}"] = int(myl[0][2])
            distances[f"{myl[0][1]},{myl[0][0]}"] = int(myl[0][2])
    
    return cities, distances

def main():
    cities, distances = load_data()
    short = math.inf
    path_short = ""
    long = 0
    path_long = ""

    for perm in permutations(cities.keys()):
        # get the distance for the perm
        #print(perm)
        dist = 0
        for d in range(len(perm) - 1):
            dist += distances[f"{perm[d]},{perm[d+1]}"]

        if dist < short:
            short = dist
            path_short = perm
        if dist > long:
            long = dist
            path_long = perm


    print(f"Part 1 shortest distance: {short}")
    print(f"\t{path_short}")

    print(f"Part 2. Longest distance: {long}")
    print(f"\t{path_long}")

main()
#!/usr/bin/env python3

import re
import sys

def solve():
    # Initialize grids: 1000x1000 
    # Part 1 uses booleans (or 0/1), Part 2 uses integers
    grid1 = [0] * 1000000
    grid2 = [0] * 1000000

    # Regex to capture the command and the four coordinates
    pattern = re.compile(r'(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)')

    # Read from stdin or a file
    for line in sys.stdin:
        match = pattern.match(line.strip())
        if not match:
            continue
            
        op, r1, c1, r2, c2 = match.groups()
        r1, c1, r2, c2 = map(int, [r1, c1, r2, c2])

        # Iterate through the defined rectangle
        for r in range(r1, r2 + 1):
            row_offset = r * 1000
            for c in range(c1, c2 + 1):
                idx = row_offset + c
                
                if op == 'turn on':
                    grid1[idx] = 1
                    grid2[idx] += 1
                elif op == 'turn off':
                    grid1[idx] = 0
                    grid2[idx] = max(0, grid2[idx] - 1)
                elif op == 'toggle':
                    grid1[idx] = 1 - grid1[idx]
                    grid2[idx] += 2

    print(f"Part 1 (Lit lights): {sum(grid1)}")
    print(f"Part 2 (Total brightness): {sum(grid2)}")

if __name__ == "__main__":
    solve()
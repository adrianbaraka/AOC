#!/usr/bin/env python3

def part1(num, max):
    num=str(num)
    times = 0
    while True:
        print(f"\r{times+1}/{max} ", flush=True, end="")

        new_num=""
        i=0
        while i != len(num):
            char = num[i]
            count = 1
            i+=1
            for j in range(i, len(num)):
                next_char = num[j]
                if char == next_char:
                    count+=1
                    i+=1
                else:
                    break
            
            new_num+=f"{count}{char}"
        
        num = new_num
        times += 1
        if times == 40:
            print(f"\nLen part 1: {len(str(num))}")
            #return 0

        if times == max:
            print(f"\nLen part 2: {len(str(num))}")
            return 0


part1(1321131112, 50)

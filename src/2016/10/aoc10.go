package main

import (
	"bufio"
	"fmt"
	"local/lib"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)
var re = regexp.MustCompile(`^value ([0-9+]+) goes to ([a-zA-Z]+ [0-9]+)$`)
var re2 = regexp.MustCompile(`^(bot [0-9]+) gives low to (\w+ [0-9]+) and high to (\w+ [0-9]+)$`)

func load() []string{
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	data := [] string {}

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		data = append(data, scanner.Text())
	}

	return data

}

func toInt(s string) int {
	val, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	data := load()
	bots := make(map [string] [] int)

	//condition := [] int {5, 2}
	//condition2 := [] int {2, 5}

	condition := [] int {61, 17}
	condition2 := [] int {17, 61}
	var p1 bool
	//bs := [] string {"output 0", "output 1", "output 2"}

	stopping := func (s string) bool {
		
		if slices.Compare(condition, bots[s]) == 0 || slices.Compare(condition2, bots[s]) == 0 {
			if ! p1 {
				fmt.Println("Part 1:", s)
				p1 = true
			}
			return true
		}
		return false
	}

	addVal := func (key string, val int) bool {
		inf := bots[key]

		if len(inf) >= 2 && ! strings.Contains(key, "output") {
			return false
		}

		bots[key] = append(bots[key], val)
		return true
		
	}
	finished := make(map [string] bool )
	for {
		for _, line := range data {
			if finished[line] {
				continue
			}
			reg1 := re.FindStringSubmatch(line)
			if reg1 != nil {

				value := toInt(reg1[1])
				dest := reg1[2]

				if addVal(dest, value) {
					finished[line] = true
				}

				//bots[dest] = append(bots[dest], value)
				
				stopping(dest)
				continue
			}

			reg2 := re2.FindStringSubmatch(line)
			if reg2 != nil {
				giver := reg2[1]
				lowRec := reg2[2]
				highRec := reg2[3]

				if len(bots[giver]) != 2 {
					continue	
				}

				// indexes
				low := 0
				high := 1
				if bots[giver][0] > bots[giver][1] {
					high = 0
					low = 1
				}

				// low := slices.Min(bots[giver])
				// high := slices.Max(bots[giver])
				newGiver := [] int {}
				if ! addVal(lowRec, bots[giver][low]) {
					newGiver = append(newGiver, bots[giver][low])
				}
				if ! addVal(highRec, bots[giver][high]) {
					newGiver = append(newGiver, bots[giver][high])
				}

				bots[giver] = newGiver

				if len(newGiver) == 0 {
					finished[line] = true
				}


				stopping(lowRec) 
				stopping(highRec)

				continue

			}
			// if reached here unmatched line error
			fmt.Fprintf(os.Stderr, "Unmatched line '%v'\n", line)
			os.Exit(1)
		}
		if len(finished) == len(data) {
			p2 := bots["output 0"][0] * bots["output 1"][0] * bots["output 2"][0]
			fmt.Println("Part 2:", p2)
			break
		}
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func load_data() ([]string, map[string]int, error) {
	// get data from input or testing
	input := "input.txt"
	if len(os.Args) > 1 && os.Args[1] == "-t" {
		input = "test.txt"
	}
	file, err := os.Open(input)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// func to convert string to integer panics if invalid
	convert := func (str string) int  {
		val, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		return val
	}

	namesT := make(map[string]struct{})
	happiness := make(map[string]int)

	regex := `^(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+).$`
	re := regexp.MustCompile(regex)

	scanner := bufio.NewScanner(file)
	line := 1
	for scanner.Scan() {
		myl := re.FindStringSubmatch(scanner.Text())

		if myl == nil {
			return nil, nil, fmt.Errorf("Could not match line %v; '%v'",line, scanner.Text())
		}
		
		// get all names
		namesT[myl[1]] = struct{}{}
		namesT[myl[4]] = struct{}{}

		d1 := fmt.Sprintf("%v,%v", myl[1], myl[4])
		h1 := convert(myl[3])

		if myl[2] == "lose" {
			h1 = - h1
		}

		happiness[d1] = h1
		line ++
	}
	var names[]string
	for name := range namesT {
		names = append(names, name)
	}

	return names, happiness, scanner.Err()
}


func main() {
	names, happiness, err := load_data()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error. %v\n", err)
		os.Exit(1)
	}
	// part 1
	run(names, happiness, "Part 1")

	// part 2
	// add myself
	me := "ME"
	// adds all happiness possibilities
	for _ , name := range names {
		//fmt.Println(name)
		h1 := fmt.Sprintf("%v,%v", me , name)
		h2 := fmt.Sprintf("%v,%v", name, me)

		happiness[h1] = 0
		happiness[h2] = 0
	}

	// add myself to list of names
	names = append(names, me)

	run(names, happiness, "Part 2")
}

func run(names []string, happiness map[string]int, part string) {
	change := 0
	var best string

	gen := combin.NewPermutationGenerator(len(names), len(names))

	for gen.Next() {
		p := gen.Permutation(nil)

		perm := make([]string, len(p))
		for i, v := range p {
			perm[i] = names[v]
		}

		//fmt.Println(perm)
		happ := 0 // total happiness change

		for h := range len(perm) {
			before := h - 1
			after := h + 1
			// link the first person to last person
			if h == 0 {
				before = len(perm) - 1
			}
			// link the last person to the first
			if h == len(perm)-1 {
				after = 0
			}
			h1 := fmt.Sprintf("%v,%v", perm[h], perm[before])
			h2 := fmt.Sprintf("%v,%v", perm[h], perm[after])

			happ += happiness[h1]
			happ += happiness[h2]
		}

		if happ > change {
			change = happ
			best = strings.Join(perm, " ")
		}
	}

	fmt.Println(part,":", change)
	fmt.Println("\t", best)
}
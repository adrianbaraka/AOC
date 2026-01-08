package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func load_data() ([]string, map[string]int, error) {
	input := "input.txt"
	if len(os.Args) > 1 && os.Args[1] == "-t" {
		input = "test.txt"
	}
	file, err := os.Open(input)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	citiesT := make(map[string]struct{})
	distances := make(map[string]int)

	regex := `^([A-Za-z]+) to ([A-Za-z]+) = ([0-9]+$)`
	re := regexp.MustCompile(regex)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		myl := re.FindStringSubmatch(scanner.Text())
		
		// get all cities 
		citiesT[myl[1]] = struct{}{}
		citiesT[myl[2]] = struct{}{}

		// both to and fro
		d1 := fmt.Sprintf("%v,%v", myl[1], myl[2])
		d2 := fmt.Sprintf("%v,%v", myl[2], myl[1])
		distances[d1], _ = strconv.Atoi(myl[3]) 
		distances[d2], _ = strconv.Atoi(myl[3])
	}
	var cities[]string
	for city := range citiesT {
		cities = append(cities, city)
	}

	return cities, distances, scanner.Err()
}

func main() {
	cities, distances, e := load_data()
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}

	short := math.MaxInt
	var shortPath string

	long := 0
	var longPath string


	gen := combin.NewPermutationGenerator(len(cities), len(cities))

	for gen.Next() {
		p := gen.Permutation(nil)

		perm := make([]string, len(p))
		for i, v := range p {
			perm[i] = cities[v]
		}

		dist := 0
		for d := range len(perm) - 1 {
			d1 := fmt.Sprintf("%v,%v", perm[d], perm[d+1])
			dist += distances[d1]
		}

		if dist < short {
			short = dist
			shortPath = strings.Join(perm, " ")
		}

		if dist > long {
			long = dist
			longPath = strings.Join(perm, " ")
		}

	}

	fmt.Println("Part 1", short)
	fmt.Println("\t", shortPath)

	fmt.Println("Part 2:", long)
	fmt.Println("\t", longPath)
}
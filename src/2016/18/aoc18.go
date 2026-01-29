package main

import (
	"bufio"
	"fmt"
	"local/lib"
	"local/lib/coords"
	"os"
	"strings"
)

func load() string {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	var input string
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		input = scanner.Text()
		break
	}
	return input

}

func isSafe(left, centre, right string) bool{
	if left == "^" && centre == "^" && right == "." {
		return false
	}
	if centre == "^" && right == "^" && left == "." {
		return false
	}
	if left == "^" && centre == "." && right == "." {
		return false
	}
	if left == "." && centre == "." && right == "^" {
		return false
	}
	return true

}

func run( input string, rows int) int{

	var safeTiles int

	arr := coords.NewTwo_Darray[string](rows * len(input), len(input))
	// set initial state
	for a, val := range input {
		arr.Set(string(val), coords.Coordinate{Row: 0, Column: a})
		if string(val) == "." {
			safeTiles ++
		}
	}

	getVal := func (str string, neighbours map[string]coords.Coordinate) string {		
		n, ok := neighbours[str]
		if ok {
			val, _ := arr.GetVal(n)
				return val
			}
		return "."
	}
	for i := 1; i < rows; i++ {
		for j := range len(input) {
			c := coords.NewCoordinate(i, j)
			neighbours := arr.GetNeighboursMap(c)


			left := getVal("upLeft", neighbours)
			centre := getVal("up", neighbours)
			right := getVal("upRight", neighbours)

			if isSafe(left, centre, right) {
				arr.Set(".", c)
				safeTiles ++
			} else {
				arr.Set("^", c)
			}
		}
	}

	//arr.Visual()

	return safeTiles
}

func main() {
	input := load()
	p1 := run(input, 40)
	fmt.Println("Part 1:", p1)

	p2 := run(input, 400000)
	fmt.Println("Part 2:", p2)
}
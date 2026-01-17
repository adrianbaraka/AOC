package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"local/lib"
)


func validTriangle(l1, l2, l3 int) bool{
	// only 3 combinations just hard code
	// l1,l2 l1,l3, l2,l3
	if ((l1 + l2) <= l3) || ((l1 + l3) <= l2) || ((l2 + l3) <= l1) {
		return false
	}
	return true
}

func loadData() ([][]int, error) {
	// reader, err := GetReader()
	reader, err := lib.GetReader()
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	// a slice of lists of int
	var data [][]int

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		vals := strings.Fields(scanner.Text())
		ivals := make([]int, 3)
		for i, val := range vals {
			ival, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			ivals[i] = ival

		}
		data = append(data, ivals)
	}
	return data, scanner.Err()
}

func main() {
	data, err := loadData()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var part1 int
	var part2 int
	for _, val := range data {
		if validTriangle(val[0], val[1], val[2]) {
			part1++
		}
	}
	fmt.Println("Part 1:", part1)

	for i:= 0; i < len(data); i+=3 {
		for j := range 3 {
			if validTriangle(data[i][j], data[i+1][j], data[i+2][j]) {
				part2++
			}
		}
	}
	
	fmt.Println("Part 2:", part2)
}
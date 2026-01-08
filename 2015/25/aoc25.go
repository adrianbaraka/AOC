package main

import "fmt"


func position(row int, col int) int {

	firstn := func (num int) int {
		return (num * ( num + 1 )) / 2
	}

	return firstn(row + col - 1) - (row-1)
}

func next_code(prev int) int {
	return (prev * 252533) % 33554393
}

func main() {
	row := 2947
	col := 3029
	code := 20151125
	finish := position(row, col)
	i := 2
	for {
		code = next_code(code)
		if i == finish {
			fmt.Println(code)
			break
		}
		i++

	}
}
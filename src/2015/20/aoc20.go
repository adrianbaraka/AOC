package main

import (
	"fmt"
	"math"
)

func num_presents(house int) (int) {

	if house < 1 {
		return 0
	}
	total := 0

	for i := 1; i <= int(math.Sqrt(float64(house))) + 1; i++ {
		if house % i == 0 {
			total += i * 10
			
			divisor := house/i
			if divisor != i {
				total += divisor * 10
			}
		}
	}
	//fmt.Printf("%v=%v\n", house, total)
	return total
}

func num_presents2(house int) (int) {

	if house < 1 {
		return 0
	}
	total := 0

	add := func (a int) {
		//fmt.Println("a:",a*50, "house", house)
		if house <= a*50 {
			total += a*11
		}
	}

	for i := 1; i <= int(math.Sqrt(float64(house))) + 1; i++ {
		if house % i == 0 {
			add(i)
			
			divisor := house/i
			if divisor != i {
				add(divisor)
			}
		}
	}
	//fmt.Printf("%v=%v\n", house, total)
	return total
}


func main() {

	newFunction()
}

func newFunction() {
	i := 0
	input := 36000000

	for {
		if num_presents2(i) >= input {
			fmt.Println(i)
			break
		}
		i++
	}
}
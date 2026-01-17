package main

import (
	"bufio"
	"fmt"
	"local/lib"
	"local/lib/coords"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func rect (a int, b int, arr *coords.Two_Darray[string] ) error {
	for i:= range b {
		for j := range a {
			coord := coords.Coordinate{Row: i, Column: j}
			err := arr.Set("1", coord)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func rotateRow (a, b int, arr *coords.Two_Darray[string]) error{
	// keep a map of the original row and the values
	before := make(map[coords.Coordinate] string )
	for i := range arr.Columns {
		coord := coords.Coordinate{Row: a, Column: i}
		val, err := arr.GetVal(coord)
		if err != nil {
			return err
		}
		before[coord] = val

	}

	for k := range before {
		ncol := (k.Column + b) % arr.Columns
		ncoord := coords.Coordinate{Row: k.Row, Column: ncol}
		err := arr.Set(before[k], ncoord)
		if err != nil {
			return err
		}
		
	}
	return nil

}

func rotateColumn(a, b int, arr *coords.Two_Darray[string]) error {
	rows := arr.Length / arr.Columns
	before := make(map[coords.Coordinate] string )
	for i := range rows {
		coord := coords.Coordinate{Row: i, Column: a}
		val, err := arr.GetVal(coord)
		if err != nil {
			return err
		}
		before[coord] = val
	}

	for k := range before {
		nrow := (k.Row + b) % rows
		ncoord := coords.Coordinate{Row: nrow, Column: k.Column}
		err := arr.Set(before[k], ncoord)
		if err != nil {
			return err
		}
		
	}


	return nil
}

func numLit(arr *coords.Two_Darray[string]) int {
	var num int
	for _, val := range arr.FlatArray {
		if val == "1" {
			num ++
		}
	}
	return num
}

// custom visualization
func visual(on string, off string, a * coords.Two_Darray[string]) {
	fmt.Println()
	for i := 0; i < a.Length; i += a.Columns {
		for j:= i; j < i + a.Columns; j++ {
			val :=  a.FlatArray[j]
			if val == "1" {
				//fmt.Printf("%v", val)
				fmt.Printf("\x1B[92m") // green
				fmt.Printf("%v", on)
				fmt.Printf("\x1B[0m")
			} else {
				fmt.Printf("\x1B[30m") // black
				fmt.Printf("%v", off)
				fmt.Printf("\x1B[0m")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

//die
func fatal(str string ,e error) {
	fmt.Fprintf(os.Stderr ,"%v, %v", e, str)
	os.Exit(1)
}

func main() {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()

	r1 := regexp.MustCompile(`^rect (\d+)x(\d+)$`)
	r2 := regexp.MustCompile(`^rotate (\w+) [x|y]=(\d+) by (\d+)$`)

	scanner := bufio.NewScanner(reader)

	// initialize the grid with values
	wide := 50
	tall := 6
	// wide := 7
	// tall := 3
	
	arr := coords.NewTwo_Darray[string](wide * tall, wide)
	// convert a string to int
	toInt := func (str string) int {
		a, err := strconv.Atoi(str)
		if err != nil {
			fatal(str, err)
		}
		return a
		
	}

	for scanner.Scan() {
		// empty string
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}

		// handle rect
		case1 := r1.FindStringSubmatch(scanner.Text())
		if case1 != nil {
			a := toInt(case1[1])
			b := toInt(case1[2])
			err := rect(a, b, arr)
			if err != nil {
				fatal(scanner.Text(), err)
			}
			continue
		}

		// handle rotate
		case2 := r2.FindStringSubmatch(scanner.Text())
		if case2 != nil {
			a := toInt(case2[2])
			b := toInt(case2[3])

			if case2[1] == "column" {
				e := rotateColumn(a, b, arr)
				if e != nil {
					fatal(scanner.Text(), e)
				}
				continue
			} else {
				e1 := rotateRow(a, b, arr)
				if e1 != nil {
					fatal(scanner.Text(), e1)
				}
				continue
			}
		}

		// if reached here no match found is an error
		fatal(scanner.Text(), fmt.Errorf("No match found"))
	}


	fmt.Println("Part 1:",numLit(arr))
	fmt.Println("Part 2: ")
	visual("#", " ", arr)

}
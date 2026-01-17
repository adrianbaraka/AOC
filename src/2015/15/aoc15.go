package main

import (
	"bufio"
	"fmt"
	"strconv"
	"io"
	"os"
	"regexp"
)

type properties struct {
	capacity, durability, flavor, texture, calories int
}

var (
    regex = `^([A-Za-z]+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)$`
    r     = regexp.MustCompile(regex)
)

func extract(line string) (properties, error) {
	match := r.FindStringSubmatch(line)

	if len(match) != 7 {
		return properties{}, fmt.Errorf("Unsupported regex, %v", line)
	}

	parseInt := func (s string) int  {
			val, err := strconv.Atoi(s)

			if err != nil{
				panic(err)
			}
			return val
		
	}

    // Convert strings to ints. 
	return properties{
			capacity:   parseInt(match[2]),
			durability: parseInt(match[3]),
			flavor:     parseInt(match[4]),
			texture:    parseInt(match[5]),
			calories:   parseInt(match[6]),
		}, nil
}

func load_data() ([]properties, error) {
	var reader io.Reader

	reader = os.Stdin
	if len(os.Args) > 1 && os.Args[1] != "-" {
		file, err := os.Open(os.Args[1])
		if err != nil {
			return nil, err
		}
		defer file.Close()
		reader = file
	}

	scanner := bufio.NewScanner(reader)

	// logic to increase buffer size
	// // 1. Create a buffer (starting size 64KB)
    // const maxCapacity = 1024 * 1024 // 1MB
    // buf := make([]byte, 64*1024)

    // // 2. Tell the scanner to use this buffer and set the max limit
    // scanner.Buffer(buf, maxCapacity)

	//info := make(map[string]properties)
	info := []properties{}



	for scanner.Scan() {
		// do some work on the line
	
		props, err := extract(scanner.Text())

		if err != nil {
			fmt.Fprintln(os.Stderr, "Skipping line:", err)
			continue
		}

		info = append(info, props)
	}

	return info, scanner.Err()
}


func check(quantities [] int, data [] properties, consider_calories bool) (int, error) {

	var cap, dur, flav, text, cal int


	for i := range quantities {
		//fmt.Println("len:",len(quantities), "i:", i)
		//fmt.Println(i)
		cal  += quantities[i] * data[i].calories
		cap  += quantities[i] * data[i].capacity
		dur  += quantities[i] * data[i].durability
		flav += quantities[i] * data[i].flavor
		text += quantities[i] * data[i].texture
	}

	normalized := func (val int) int  {
		if val < 0 {
			val = 0
		}
		return val
	}


	score := normalized(cap) * normalized(dur) * normalized(flav) * normalized(text)

	if consider_calories && cal != 500{
		score = 0
	}

	//fmt.Printf("\t%v\n", score)

	return score, nil
}

func get_score (quantities[]int, data []properties, part2 bool) (int, error){
	score, err := check(quantities, data, part2)
	if err != nil {
		return 0, err
	}

	return score, nil
}


func main() {
	data, err := load_data()
    if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var max1, max2 int
	quantities := make([]int, 4)

	for i:= 0; i <= 100; i++ {
		for j:= 0; j<= 100; j++ {
			for k:= 0; k<= 100; k++ {
				l := 100 - (i + j + k)
				quantities[0], quantities[1], quantities[2], quantities[3] = i, j, k, l 
				s1, err := get_score(quantities, data, false)
				s2, err2 := get_score(quantities, data, true)

				if err != nil || err2 != nil {
					fmt.Printf("An error occurred %v %v", err, err2)
					os.Exit(1)
				}

				if s1 > max1 {
					max1 =s1
				}

				if s2 > max2 {
					max2 =s2
				}
			}
		}
	} 

	fmt.Println("Part 1:", max1)
	fmt.Println("Part 2:", max2)
}
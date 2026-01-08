package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)


func load_data() (string, map[string][]string, error) {
	var reader io.Reader

	reader = os.Stdin
	if len(os.Args) > 1 && os.Args[1] != "-" {
		file, err := os.Open(os.Args[1])
		if err != nil {
			return "", nil, err
		}
		defer file.Close()
		reader = file
	}

	scanner := bufio.NewScanner(reader)

	info := make(map[string][]string)
	mol := false
	var molecule string

	for scanner.Scan() {
		if mol {
			molecule = scanner.Text()
			continue
		}

		if scanner.Text() == "" {
			mol = true
			continue
		}

		sections := strings.Split(scanner.Text(), " => ")
		info[sections[0]] = append(info[sections[0]], sections[1])
	}

	return molecule, info,  scanner.Err()
}

func load_data2() (string, map[string] string, error) {
	var reader io.Reader

	reader = os.Stdin
	if len(os.Args) > 1 && os.Args[1] != "-" {
		file, err := os.Open(os.Args[1])
		if err != nil {
			return "", nil, err
		}
		defer file.Close()
		reader = file
	}

	scanner := bufio.NewScanner(reader)

	info := make(map[string]string)
	mol := false
	var molecule string

	for scanner.Scan() {
		if mol {
			molecule = scanner.Text()
			continue
		}

		if scanner.Text() == "" {
			mol = true
			continue
		}

		sections := strings.Split(scanner.Text(), " => ")
		info[sections[1]] = sections[0]
		// info[sections[0]] = append(info[sections[0]], sections[1])
	}

	return molecule, info,  scanner.Err()
}


func part2() {
	/*
		Works by generating all possible substrings of the molecule if that molecule can be replaced it is replaced.
		This is down recursively if the state of the molecule is the same as after the loop no change break.
		The not efficient part is the substring generation On sqaured. got lucky ;)

		Not sure but this looks like Bottom up parsing.
			A valid string defined by that grammar should end in e any other symbols indicate invalid.
	*/
	molecule, info, err := load_data2()
	checkErr(err)
	steps := 0
	for {
		start := molecule
		for i := range len(molecule) {
			for j := i+1 ; j < len(molecule) +1; j++ {
				to_check := molecule[i:j]
				val, ok := info[to_check]
				if ! ok {
					continue
				}
				//fmt.Printf("Before: %v\n", molecule)
				molecule = replace(molecule, val, i, len(to_check))
				//fmt.Printf("%v\n", molecule)
				steps ++

			}
		}
		if start == molecule {
			break
		}
	}
	fmt.Println("Part 2:", steps)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func in_map(str string, dict map[string][]string) (bool) {
	_, in_dict := dict[str]
	if ! in_dict {
		return false
	}
	return true
}

func replace(molecule string, str string, index int, width int) (string) {
	return molecule[:index] + str + molecule[index+width:]
}


func part1() {
	molecule, info, err := load_data()
	checkErr(err)

	distinct := make(map[string]int)

	for i := 0; i < len(molecule); i++ {
		// if char in info replace with all possible occurences in info
		// put these in a map for uniq
		// else get the next char and do as above
		char := string(molecule[i])
		built := false
		if ! in_map(char, info) {
			// get next char and get that
			if i == len(molecule)-1 {
				continue
			}
			//fmt.Printf("Not found %v: ", char)
			next_char := string(molecule[i+1])
			var str_build strings.Builder
			str_build.WriteString(char)
			str_build.WriteString(next_char)
			new_str := str_build.String()

			//fmt.Printf("Checking %v: ", new_str)

			if ! in_map(new_str, info) {
				//fmt.Fprintf(os.Stderr, "Not found %v %v\n", new_str, char)
				//fmt.Printf(" Not found %v continuing i=%v\n", new_str, i)
				continue
			}
			char = new_str
			built = true
			//fmt.Printf(" found %v continuing i=%v\n", new_str, i)
		}

		//fmt.Println("Checking char: ", char, "i=", i)
		for _, val := range info[char] {
			new_str := replace(molecule, val, i, len(char))
			distinct[new_str]++
		}
		if built {
			i++
		}
	}

	// for k, v := range distinct {
	// 	fmt.Printf("%v = %v\n\n", k, v)
	// }
	fmt.Println( "Part 1:", len(distinct))
	
}
func main() {
	part1()
	part2()
}
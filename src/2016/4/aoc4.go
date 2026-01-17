package main

import (
	"bufio"
	"fmt"
	"local/lib"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// compile the regex globally
const regex = `((?:[a-z]-?)+)([0-9]+)\[([a-z]+)\]`
var re = regexp.MustCompile(regex)

func extract(line string) ( string, map[string]int, int, string, error) {
	data := re.FindStringSubmatch(line)
	if data == nil {
		return "", nil, 0, "", fmt.Errorf("No match found on line %v", line)
	}
	// coz of part 2 dont replace just remove the last one -
	//name := strings.ReplaceAll(data[1], "-", "")
	name := strings.TrimSuffix(data[1], "-")
	id, err := strconv.Atoi(data[2])
	if err != nil {
		return "",nil, 0, "", err
	} 
	checksum := data[3]

	// loop through name and get frequency of ocurrences
	freq := make(map[string]int)
	for _,char := range name {
		if char == '-' {
			continue // do not include the -
		}
		freq[string(char)] ++
	}



	return name, freq, id, checksum, nil

}

func isReal(freq map[string]int, checksum string) bool {
	//loop through the chars in checksum if freq greater than before not real or not in freq
	current := string(checksum[0])
	for k := range freq {
		if freq[string(k)] > freq[current] {
			current = k
		}
	}
	for _, char := range checksum {
		val, ok := freq[string(char)]
		if ! ok {
			// not in freq
			//fmt.Println("false")
			return false
		}
		if val > freq[current] {
			//fmt.Println("false")
			return false
		} else if val == freq[current] {
			// if same current must be lexicographically less than char basically sorted alphabetically
			if current > string(char) {
				return false
			}
			
		}
		current = string(char)
	}
	//fmt.Println("true")
	return true
}

func rotate(name string, id int) string{
	var final strings.Builder
	for _, char := range name {
		//new := ((char + rune(id)) % 26) + 97
		// convert to alphabet 0-25 add the shift and mod by 26 then convert back to decimal
		new := (((char - 97) + rune(id)) % 26) + 97
		if string(char) == "-" {
			final.WriteString(" ")
			continue
		}
		final .WriteString(string(new))
	}
	

	//fmt.Println(final.String())
	//fmt.Printf("<%v> id: %v\n", final.String(), id)
	return final.String()
}

func main() {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	part1 := 0
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		name, freq, id, checksum, err := extract(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if isReal(freq, checksum) {
			part1 += id
			plain := rotate(name, id)
			// look for any string with north in it
			if strings.Contains(plain, "north") {
				fmt.Printf("Part 2: '%v' id = %v\n", plain, id)
			}
		}

	}

	fmt.Println("Part 1:",part1)
}
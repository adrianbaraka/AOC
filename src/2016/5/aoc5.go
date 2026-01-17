package main

import (
	"slices"
	"bufio"
	"crypto/md5"
	"fmt"
	"local/lib"
	"os"
	"strconv"
	"strings"
)
// checks if all elements of the array have an element
func fullArray (arr []string) bool {
	return !slices.Contains(arr, "")
}

// return the md5hash of a string
func md5Hash(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func main() {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	input := ""
	for scanner.Scan() {
		input = scanner.Text()
	}

	var pass strings.Builder
	pass2 := make([]string, 8)
	var i int
	var part1 bool

	// animation for part 2
	fmt.Printf("  ********\r")

	for {
		str := input + strconv.Itoa(i)
		mdhash := md5Hash(str)
		if mdhash[:5] != "00000" {
			i++
			continue
		}
		// part 1
		char6 := string(mdhash[5])
		if ! part1 {
			pass.WriteString(char6)
		}

		// part 2
		// if char6 is in the range 0-7 add char7 to the list
		char6Int, err := strconv.Atoi(char6)
		if err == nil && char6Int < 8 && char6Int >= 0 {
			char7 := string(mdhash[6])
			//fmt.Printf("Found i = %v, char6 : %v char7: %v\n", i, char6, char7)
			// if already found do not ovewrite
			if pass2[char6Int] == "" {
				pass2[char6Int] = char7 // add the seventh char 
			}

			// cinematic decrypting
			var currentpass strings.Builder
			for _, char := range pass2 {
				if char == "" {
					char = "*"
				}
				currentpass .WriteString(char)
			}
			fmt.Printf("  %v\r", currentpass.String()) // 2 spaces for visibility from cursor
		}
		// done part 1
		if ! part1 && len(pass.String()) == 8 {
			part1 = true
		}
		// stopping condition
		if part1 && fullArray(pass2) {
			fmt.Println("\rPart 1:",pass.String()) // clear the line first
			fmt.Printf("Part 2: %v\n", strings.Join(pass2, ""))
			break
		}

		i++
	}
}
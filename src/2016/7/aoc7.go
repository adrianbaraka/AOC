package main

import (
	"bufio"
	"fmt"
	"local/lib"
	"os"
	"strings"
)


func hasAbba(str string) bool {
	//fmt.Println(str)
	for i := 0; i < len(str) -3; i++ {
		if lib.Is_palindrome(str[i:i+4]) && str[i] != str[i+1] {
			//.Println("True")
			return true
		}
	}
	//fmt.Println("false")
	return false

}

// returns a slice of all aba sequences in the given string
func areaba(str string) ( []string){
	list := [] string{}
	for i := 0; i <= len(str)-3; i++ {
		if str[i] == str[i+2] && str[i] != str[i+1] {
			s := fmt.Sprintf("%v%v%v", string(str[i]), string(str[i+1]), string(str[i+2]))
			list = append(list, s)
		}
	}
	return list
}


func supportSsl(out []string, in []string) bool{

	//fmt.Println("Out:", out)
	//fmt.Println("in:", in)

	sin := strings.Join(in, " ")
	// reverses aba to get bab
	reverseAba := func (s string) string{
		return fmt.Sprintf("%v%v%v", string(s[1]), string(s[0]), string(s[1]))
		
	}
	outaba := [] string {}
	for _, str := range out {
		outaba = append(outaba, areaba(str)...)
	}
	//fmt.Println("outaba", outaba)
	for _, aba := range outaba {
		//fmt.Printf("aba : %v Reversed aba %v\n",aba, reverseAba(aba))
		if strings.Contains(sin, reverseAba(aba)) {
			return true
		}
	}

	//fmt.Println("false")
	return false
}

func main() {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	var supports int
	var supports2 int

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		// this leaves the ones in brackets with spaces around them and a comma in start to show that it was inside
		s := strings.ReplaceAll(scanner.Text(), "[", " ,")
		s = strings.ReplaceAll(s, "]", " ")
		fields := strings.Fields(s)

		var outside bool
		inside := true
		out := [] string {}
		in := [] string {}

		for _,field := range fields {
			//fmt.Println(field)
			if string(field[0]) == "," {
				// slice the string removing the added comma
				field = field[1:]
				in = append(in, field) // part 2

				if hasAbba(field) {
					inside = false
				}
				continue
			}

			out = append(out, field) // part 2
			if hasAbba(field) {
				outside = true
			}

		}
		if outside && inside {
			supports++
		}

		// check for part 2
		if supportSsl(out, in) {
			//fmt.Println(scanner.Text())
			supports2 ++
		}
		

	}

	fmt.Println("Part 1:",supports)
	fmt.Println("Part 2:",supports2)
}
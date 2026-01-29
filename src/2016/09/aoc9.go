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

// TODO part 2 is slow

const regex = `\(\d+x\d+\)`
var	re = regexp.MustCompile(regex)

func load() (string, error) {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()
	data := ""
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data = scanner.Text()
		break
	}
	return data, scanner.Err()

}

func toInt (str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
		
}

func part1() {
	data, err := load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var count int
	var mark bool
	var marker strings.Builder

	for i:= 0; i < len(data); i++ {
		char := string(data[i])
		if char == "(" {
			mark = true
			//fmt.Println("marker set to true")
			continue
		}

		if char == ")" {
			// parse marker
			info := strings.Split(marker.String(), "x")
			chars := toInt(info[0])
			times := toInt(info[1])

			//fmt.Println(info)

			count += chars * times

			i += chars
			mark = false
			marker.Reset()
			continue
		}

		if mark {
			marker .WriteString(char)
			continue
		}
		count ++
	}

	fmt.Println("Part 1:",count)

	part2 := expandstr4(data)
	fmt.Println("Part 2:",part2) // 10,780,403,063
}


// naive approach very slow but works
func expandstr(str string) (int, error) {
	var ns strings.Builder
	var count int

	for {
		fmt.Println(str)
		ns.Reset()
		before, temp, found := strings.Cut(str, "(")
		ns.WriteString(before)
		//count += len(before)
		if count % 10000 == 0 {
			//fmt.Printf("\r %v ", count)
		}

		if ! found {
			//return count , nil
			return len(before), nil
		}

		// parse after for times
		exp, after, f2 := strings.Cut(temp, ")")
		if ! f2 {
			return 0, fmt.Errorf("Unclosed bracket")
		}
		expr := strings.Split(exp, "x")
		a := toInt(expr[0])
		b := toInt(expr[1])

		//fmt.Println("b=",b)


		for range b {
			ns.WriteString(after[:a])
		}
		ns.WriteString(after[a:])

		//fmt.Println(ns.String())
		//fmt.Println("After:  " ,after)
		//count += len(ns.String())
		str = ns.String()
	}

	//return nil


}

func extractMarker(s string) (int, int){
	s = s[1:len(s)-1]
	str := strings.Split(s, "x")

	a := toInt(str[0])
	b := toInt(str[1])

	return a, b

}

func expandstr2(s string) int {
	//original := strings.Clone(s)
	var totalCount int
	var idx int
	info := make( map[string] int)
	for{
		fmt.Println(s)
		fmt.Println("idx", idx)
		//fmt.Println()
		str := re.FindStringIndex(s)

		if str == nil {
			totalCount += len(s)
			break
		}
		//fmt.Println(str)
		y := str[0]
		z := str[1]

		totalCount += y

		a, b := extractMarker(s[y:z])



		uniqva := s[y:z] + strconv.Itoa(idx)
		fmt.Println("uniqva", uniqva)


		va, ok := info[uniqva]

		if ok {
			//fmt.Println("reached")
			b *= va
		}
		//fmt.Println("b:", b)

		nextbchars := s[z:z+a]

		//fmt.Println(nextbchars)

		//fmt.Printf("nextbchars %v len %v b = %v \n",nextbchars, len(nextbchars), b)

		//ex := re.FindAllString(nextbchars, -1)
		ex2 := re.FindAllStringIndex(nextbchars, -1)
		if ex2 == nil {
			totalCount += len(nextbchars) * b
			z += len(nextbchars)
			//idx += len(nextbchars)
		} else {
			// get the index of first occurrence
			//first0 := re.FindStringIndex(nextbchars)
			//fmt.Println("first0",first0)

			// add all not marker
			//m := re.ReplaceAllString(nextbchars, "")

			//idx += len(m) 

			// add all others to map
			// fmt.Println(ex2)
			// for _, val := range ex{
			// 	//fmt.Println(val)
			// 	//info[val] +=  b

			// 	// get the index of val to make it unique
			// 	val = val + strconv.Itoa(idx)
			// 	info[val] = b
			// }
			for _, val := range ex2 {
				ns := nextbchars[val[0]:val[1]]
				ns = ns + strconv.Itoa(val[1] + 1)
				info[ns] = b
			}
		}
		s = s[z:]
		idx += z
		fmt.Println("z=",z)
		fmt.Println(info)
		fmt.Println()


		if s == "" {
			break
		}
	}
	
	//fmt.Println(totalCount)
	return totalCount
	// get the next a chars 
	// if a marker is there add to the map that marker with value += b
	// get count of all before marker multiply by b and to counter

	
}

func explandstr3(s string) int {
	var totalCount int
	for {
		//fmt.Println("s", s, "count", totalCount)
		first := re.FindStringIndex(s)

		if first == nil {
			totalCount += len(s)
			break
		}

		//fmt.Println(first)

		y := first[0]
		z := first[1]

		//beforechars := s[:z]

		//fmt.Println("Before chars", beforechars)

		totalCount += y
		width, multi := extractMarker(s[ y : z ])

		//fmt.Println(width, multi)

		nextwidthchars := s[ z : z + width ]

		//fmt.Println("next width chars", nextwidthchars)

		ex := re.FindAllString(nextwidthchars, -1)


		if ex == nil {
			totalCount += len(nextwidthchars) * multi
			z += len(nextwidthchars)
			s = s[z:]
		} else {
			var newS strings.Builder
			for _, val := range ex {
				aval := updateMarker(val, multi)
				nextwidthchars = strings.Replace(nextwidthchars, val, aval, 1)
			}

			//aval := updateMarker(ex[0], multi)
			//nextwidthchars = strings.Replace(nextwidthchars, ex[0], aval, 1)
			//newS.WriteString(beforechars)
			//fmt.Println(nextwidthchars)
			newS.WriteString(nextwidthchars)
			newS.WriteString(s[z + width:])

			//fmt.Printf("new '%v'\n", s[z + width:])

			s = newS.String()
			newS.Reset()

		}



		if s == "" {
			//fmt.Println("here")
			break
		}

		
		// ex2 := re.FindAllStringIndex(nextwidthchars, -1)
		// if ex2 == nil {
		// 	totalCount += len(nextwidthchars) * multi
		// 	z += len(nextwidthchars)
			
		// } else {
		// 	var newS strings.Builder
		// 	for _, val := range ex2 { 
		// 		//fmt.Println("Ranges with marker", val, nextwidthchars[ val[0] : val[1]  ])
		// 		//fmt.Println("Ranges with marker", val, s[ len(beforechars) + val[0] : len(beforechars) + val[1]  ])

		// 		aval := nextwidthchars[ val[0] : val[1]  ]
		// 		// fmt.Println()
		// 		// fmt.Println("lval", aval, "len", len(aval))
		// 		// fmt.Println("before chars",beforechars, "len", len(beforechars))
		// 		// fmt.Println("next width chars",nextwidthchars, "len", len(nextwidthchars), "z", z)
		// 		// //fmt.Println("after next" , s[z + len(nextwidthchars):])
		// 		// fmt.Println("after next" , s[z + len(aval):])
		// 		// fmt.Println("S:", s)
		// 		newMarker := updateMarker(nextwidthchars[ val[0] : val[1]  ], multi)

		// 		//fmt.Println("Newmarker",newMarker)
		// 		newS.WriteString(beforechars)
		// 		newS.WriteString(newMarker)
		// 		newS.WriteString(s[z + len(aval):])

		// 		s = newS.String()
		// 		// take beforechars > newmarker > s[z + len(aval):

		// 	}
		// 	newS.Reset()
		// }
		//s = s[z:]
		
	}
	//fmt.Println("Total count:", totalCount)
	return totalCount	
}

func expandstr4(s string) int {
    // Base case: if there are no markers, the length is just the length of s
    first := re.FindStringIndex(s)
    if first == nil {
        return len(s)
    }

    // Split the string into: 
    // 1. Prefix (before the marker)
    // 2. The marker itself
    // 3. The data affected by the marker
    // 4. The rest of the string
    
    startOfMarker := first[0]
    endOfMarker := first[1]
    
    // 1. Everything before the marker counts as its literal length
    count := startOfMarker

    // Extract marker info (e.g., 10x5 -> width 10, multi 5)
    width, multi := extractMarker(s[startOfMarker:endOfMarker])

    // 2. The "Section" affected by the marker
    section := s[endOfMarker : endOfMarker+width]
    
    // RECURSIVE STEP: 
    // Instead of expanding, calculate the length of the section 
    // and multiply it by the marker's multiplier.
    count += explandstr3(section) * multi

    // 3. The "Remainder" of the string after the affected section
    remainder := s[endOfMarker+width:]
    count += explandstr3(remainder)

    return count
}

func updateMarker(s string, newNum int) string {
	s = s[ 1 : len(s) -1 ]

	split := strings.Split(s, "x")

	newNum = toInt(split[1]) * newNum

	return fmt.Sprintf("(%vx%v)", split[0], newNum)
}
func main() {
	//s := "(3x3)XYZ"
	//s := "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"
	//s:= "(27x12)(20x12)(13x14)(7x10)(1x12)A"
	// //s := "X(8x2)(3x3)ABCY"
	// ss := [] string {"(3x3)XYZ", 
	// 				"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 
	// 				"(27x12)(20x12)(13x14)(7x10)(1x12)A", 
	// 				"X(8x2)(3x3)ABCY" }
	// //fmt.Println(expandstr2(ss[1]))
	// explandstr3(ss[1])
	// for _, s := range ss {
	// 	fmt.Println(s, explandstr3(s))
	// }
	//s := "(27x12)(20x12)(13x14)(7x10)(1x12)A"
	//fmt.Println(expandstr(s))
	part1()
}
package main

import "fmt"
//TODO finish rewrite in go
func getNextSeq() {
	str := "abc"

	add := fmt.Sprintf("%0*d", len(str), 1)
	var carry byte = 0
	var final []byte

	for i:=len(str)-1; i >= 0; i-- {
		val := str[i] + byte(int(add[i])) + byte(carry)

		fmt.Println(byte(carry))
		if val > 122 {
			diff := val - 122
			val = 96 + diff
			carry = diff
		} else {
			carry = 0
		}
		final = append(final, val)
	}

	//fmt.Println(final)
}

func main() {
	getNextSeq()
}
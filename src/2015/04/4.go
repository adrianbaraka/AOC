package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func ret_hash(str string) string{
	h := md5.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func match(str string) bool {
	target := "000000"
	if str[:6] == target {
		return true
	}
	return false
}

func main() {
	start := "iwrupvqb"

	var i int
	for {
		//fmt.Printf("\rTrying %v ", i)
		newstring := start + strconv.Itoa(i)
		if match(ret_hash(newstring)) {
			fmt.Printf("Num found: %v\n", i)
			break
		}
		i++
	}
	
}
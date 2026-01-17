package secho

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func Secho(a ... any) {
	fmt.Printf("%v", a...)
}



func Sechod() {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Println("terminal")
	} else {
		fmt.Println("Not a terminal")
	}	
}
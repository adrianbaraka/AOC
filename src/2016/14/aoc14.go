package main

import (
	"bufio"
	"fmt"
	"local/lib"
	"local/lib/coords"
	"os"
	"strings"
)

func load() {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		fmt.Println(scanner.Text())
	}

}

type cust struct {
	val int
}

func main() {
	c := coords.NewCoordinate(1, 1)
	fmt.Printf("%T\n", c)

	ar := coords.NewTwo_Darray[int](50, 10)

	ar.Set(5, c)
	

	ar.Visual()



}

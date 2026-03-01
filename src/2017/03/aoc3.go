package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	
	lib "github.com/adrianbaraka/goutils"
)

func main() {
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
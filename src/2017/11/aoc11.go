package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	lib "github.com/adrianbaraka/goutils"
	"github.com/adrianbaraka/goutils/coords"
)

func load() []string {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()
	scanner := bufio.NewScanner(reader)

	var directions []string

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		//fmt.Println(scanner.Text())
		directions = strings.Split(scanner.Text(), ",")
		break
	}
	//fmt.Println(directions)
	return directions

}

func getNextPosition(direction string, current coords.Coordinate) (coords.Coordinate){
	var nextCoord = coords.NewCoordinate(0, 0)
	switch direction {
	case "nw":
		nextCoord.Row = current.Row + 1
		nextCoord.Column = current.Column - 1

	case "n":
		nextCoord.Row = current.Row + 1
		nextCoord.Column = current.Column

	case "ne":
		nextCoord.Row = current.Row + 1
		nextCoord.Column = current.Column + 1

	case "se":
		nextCoord.Row = current.Row - 1
		nextCoord.Column = current.Column + 1
	case "s":
		nextCoord.Row = current.Row - 1
		nextCoord.Column = current.Column
	case "sw":
		nextCoord.Row = current.Row - 1
		nextCoord.Column = current.Column - 1
	default:
		fmt.Fprintln(os.Stderr, "Invalid direction", direction)
		os.Exit(1)
	}
	return nextCoord

}

func main() {
	directions := load()
	//fmt.Println(directions)
	current := coords.NewCoordinate(0, 0)
	for _, direction := range directions {
		current = getNextPosition(direction, current)
	}

	fmt.Println(current)
}

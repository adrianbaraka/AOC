package main

import (
	"bufio"
	"fmt"
	"local/lib"
	"os"
	"strconv"
	"strings"
	"sync"
)

func load() [] string{
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	data := [] string {}
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		data = append(data, scanner.Text())
	}

	return data
}

func run(c int, part int) {
	data := load()

	registers := make(map [string] int)

	registers["a"] = 0
	registers["b"] = 0
	registers["c"] = c
	registers["d"] = 0

	i := 0
	for i < len(data) {
		ins := strings.Fields(data[i])
	
		switch ins[0] {
		case "cpy":
			val := ins[1]
			dest := ins[2]

			ival, err := strconv.Atoi(val)

			if err != nil {
				ival = registers[val]
			}
			registers[dest] = ival
			i ++
		case "inc":
			reg := ins[1]
			val := registers[reg]

			registers[reg] = val + 1
			i++

		case "dec":
			reg := ins[1]
			val := registers[reg]

			registers[reg] = val - 1
			i++
		case "jnz":
			x := ins[1]
			y := ins[2]

			iy, er := strconv.Atoi(y)

			if er != nil {
				panic(er)
			}

			ix, err := strconv.Atoi(x)

			if err != nil {
				ix = registers[x]
			}

			if ix != 0 {
				i += iy
			} else {
				i ++
			}
		default:
			fmt.Fprintln(os.Stderr, "Unsupported instruction ",ins[0])
			os.Exit(1)
		}
	}

	fmt.Printf("Part %v: %v\n",part, registers["a"])

}

func main() {
	var wg sync.WaitGroup

	wg.Go(func() {run(0, 1)})
	wg.Go(func() {run(1, 2)})

	wg.Wait()
}

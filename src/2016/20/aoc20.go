package main

import (
	"bufio"
	"cmp"
	"fmt"
	"local/lib"
	"os"
	"slices"
	"strings"
)

type rng struct {
	min, max int
}

func load() []rng {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	rngs := []rng{}
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		r := strings.Split(scanner.Text(), "-")
		newr := rng{min: lib.ToInt(r[0]), max: lib.ToInt(r[1])}
		rngs = append(rngs, newr)
	}

	return rngs

}

func merge(rngs []rng) []rng {
	newrngs := []rng{}
	var changed bool
	for {
		i := 0
		for i < len(rngs)-1 {
			diff := rngs[i+1].min - rngs[i].max
			if diff <= 1 {
				nr := rng{min: rngs[i].min, max: max(rngs[i].max, rngs[i+1].max)}
				newrngs = append(newrngs, nr)
				newrngs = append(newrngs, rngs[i+2:]...)
				changed = true
				break
			} else {
				newrngs = append(newrngs, rngs[i])
			}
			i++
		}
		// something was merged loop again
		if changed {
			rngs = rngs[:0] // clear the slice but keeps the underlying capacity so no reallocation
			rngs = append(rngs, newrngs...)
			newrngs = newrngs[:0]
			changed = false
		} else {
			break
		}
	}
	//fmt.Println(rngs)
	return rngs

}
func main() {
	rngs := load()

	cmpFunc := func(a rng, b rng) int {
		// better than implementing myself
		return cmp.Compare(a.min, b.min)
	}
	slices.SortFunc(rngs, cmpFunc)

	// normalize to proper ranges
	rngs = merge(rngs)

	fmt.Println("Part 1:", rngs[0].max+1)
	i := 0
	total := 0
	max := 4294967295
	//max := 9
	for i < len(rngs) {
		// first one
		if i == 0 {
			difff := rngs[i].min - 0
			if difff > 0 {
				total += difff
			}
		}
		// last one
		if i == len(rngs)-1 {
			dif := max - rngs[i].max
			if dif > 0 {
				total += dif
			}
			i++
			continue
		}

		diff := rngs[i+1].min - rngs[i].max - 1
		total += diff
		i++
	}

	fmt.Println("Part 2:", total)
}

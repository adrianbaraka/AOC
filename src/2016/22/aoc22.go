package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"local/lib"
	"local/lib/coords"
	"os"
	"regexp"
	"strings"
)

type node struct {
	coord             coords.Coordinate
	size, used, avail int
}

func load() ([]node, int) {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	nodes := []node{}
	re := regexp.MustCompile(`\/dev\/grid\/node-x(\d+)-y(\d+)`)
	i := 0
	maxy := 0 // get the number of columns
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		if i < 2 {
			// skip first 2 lines
			i++
			continue
		}

		fields := strings.Fields(scanner.Text())
		if len(fields) != 5 {
			fmt.Fprintln(os.Stderr, "Unmatched line", scanner.Text())
			os.Exit(1)
		}

		coods := re.FindStringSubmatch(fields[0])
		newCoord := coords.NewCoordinate(lib.ToInt(coods[2]), lib.ToInt(coods[1]))

		// no of columns
		if lib.ToInt(coods[2]) > maxy {
			maxy = lib.ToInt(coods[2])
		}

		size := lib.ToInt(fields[1][0 : len(fields[1])-1])
		used := lib.ToInt(fields[2][0 : len(fields[2])-1])
		avail := lib.ToInt(fields[3][0 : len(fields[3])-1])

		newNode := node{coord: newCoord, size: size, used: used, avail: avail}

		nodes = append(nodes, newNode)

	}
	return nodes, maxy + 1
}

func isViable(a node, b node) bool {
	if a.used == 0 || a == b {
		return false
	}

	if a.used <= b.avail {
		return true
	}
	return false

}

func p1(nodes []node) {
	var viable int
	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {

			if isViable(nodes[i], nodes[j]) {
				viable++
			}
			if isViable(nodes[j], nodes[i]) {
				viable++
			}
		}
	}

	fmt.Println("Part 1:", viable)

}

type path struct {
	position coords.Coordinate
	parent coords.Coordinate
	steps int
}

type priorityQueue [] *path
func (pq priorityQueue) Less (i int, j int) bool {
	return pq[i].steps < pq[j].steps
}
func (pq priorityQueue) Len() int {
	return len(pq)
}
func (pq priorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *priorityQueue) Push(x any) {
	item := x.(*path)
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func p2() {
	pq := &priorityQueue{}
	heap.Init(pq)

}



func main() {
	nodes, columns := load()

	p1(nodes)

	fmt.Println(columns)
}

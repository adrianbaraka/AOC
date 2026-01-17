package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"local/lib"
	"math/bits"
	"os"
	"slices"
	"strconv"
	"strings"
)

type coord struct {
	row, column int
}

type path struct {
	currentPos coord
	parent coord
	steps int
}

type PriorityQueue [] *path

// methods for a priority queue
func (pq  PriorityQueue) Len() int {
	return len(pq)
}

// only uses steps dont know which is more efficient
func (pq PriorityQueue) Less(i, j int) bool {

	return pq[i].steps < pq[j].steps
}

func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*path)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // prevent memory leaking
	*pq = old[0 : n-1]
	return item
}

func load() (int, coord) {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()
	var num string
	var target string
	scanner := bufio.NewScanner(reader)
	i := 0
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		if i == 0 {
			num = scanner.Text()
			i ++
		} else {
			target = scanner.Text()
		}
	}

	toInt := func (s string) (int ){
		i, e := strconv.Atoi(s)
		if e != nil {
			fmt.Fprintf(os.Stderr, "Invalid input %v", err)
			os.Exit(1)
		}
		return i
		
	}
	inum := toInt(num)
	tr := strings.Fields(target)
	tgt := coord{row: toInt(tr[0]), column: toInt(tr[1])}

	return inum, tgt
}

func isOpenSpace(c coord, fav int) bool {
	// x column y row
	num := (c.column * c.column) + (3 * c.column) + 
			(2 * c.column * c.row) + c.row + (c.row * c.row) +
			fav

	n1 := bits.OnesCount16(uint16(num))


	if n1 % 2 == 0 {
		return true
	}
	return false
}

func getNeighbours(c coord) ([] coord) {
	neighbours := [] coord {}
	//n := [...] int { -1, 0, 1}

	n := make(map [string] [2] int)

	n["up"] = [...] int {-1, 0}
	n["down"] = [...] int {1, 0}
	n["left"] = [...] int {0, -1}
	n["right"] = [...] int {0, 1}


	for _, val := range n {
		ro := c.row + val[0]
		co := c.column + val[1]

		if ro >= 0 && co >= 0 {
			ord := coord{ro,co}
			neighbours = append(neighbours, ord)
		}
	}
	//fmt.Println(neighbours)
	return neighbours
}

func run(fav int, target coord) (path, [] coord, error){
	sp := coord {1,1}
	//taken := [] coord {sp}
	steps := 0
	costSoFar := make(map[coord]int)
	pth := make( map [coord] coord) 
	pth[sp] = sp

	init := path{currentPos: sp,  steps: steps, parent: sp }

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &init)
	costSoFar[sp] = 0

	reconstruct := func (c coord) []coord {
		p := [] coord {c}

		for {
			parent := pth[c]

			if c == sp {
				break
			}

			p = append(p, parent)
			c = parent
		}
		slices.Reverse(p)
		return p
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*path)
		newsteps := current.steps + 1
		for _, neigbour := range getNeighbours(current.currentPos) {

			// we have found the target stop here
			if neigbour == target {
				current.steps ++
				pth[neigbour] = current.currentPos
				//fmt.Println(current.parent)
				//fmt.Println(reconstruct(neigbour))
				return * current, reconstruct(neigbour), nil
			}

			// add all open space neigbours
			if ! isOpenSpace(neigbour, fav) {
				continue
			}

			prevsteps, ok := costSoFar[neigbour]
			if ! ok || newsteps < prevsteps {
				costSoFar[neigbour] = newsteps
				next := path {currentPos: neigbour, steps: newsteps, parent: current.currentPos}
				pth[neigbour] = current.currentPos
				heap.Push(pq, &next)
			}
		}

	}
	
	return path{}, []coord{}, fmt.Errorf("Could not find path")
}

func run2(fav int) {
	//num := 0
	checked := make(map [coord] struct{})
	for i := range 50 {
		for j := range 50 {
			c := coord{i, j}
			//fmt.Printf("%v %v \r ", i, j)
			if i + j > 50 || ! isOpenSpace(c, fav) {
				continue
			}

			sp, _, err := run(fav, c)

			if err == nil && sp.steps <= 50 {
				checked[c] = struct{}{}
				//fmt.Printf("%v\n", sp)
			}


		}
	}

	fmt.Println("Part 2:", len(checked))
}


func main() {
	fav, target := load()
	sp, _,  e := run(fav, target)
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}

	fmt.Println("Part 1:", sp.steps)
	//fmt.Println("\t", ptaken)

	run2(fav)
}
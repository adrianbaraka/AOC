package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	lib "github.com/adrianbaraka/goutils"
)

type tower struct {
	name        string
	weight      int
	parent      string
	totalWeight int
	children    []string
}

func load() map[string]*tower {
	reader, err := lib.GetReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	towers := make(map[string]*tower)
	//towers2 := newTowerlist()

	r1 := regexp.MustCompile(`^(\w+) \((\d+)\)$`)
	r2 := regexp.MustCompile(`^(\w+) \((\d+)\) -> (.+)$`)

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}

		// case 1
		case1 := r1.FindStringSubmatch(scanner.Text())
		if case1 != nil {

			// if ok, idx := towers2.inList(case1[1]); ok {
			// 	towers2.list[idx].weight = lib.ToInt(case1[2])
			// 	continue
			// }
			w1 := lib.ToInt(case1[2])
			t1 := tower{
				name:        case1[1],
				weight:      w1,
				totalWeight: w1,
			}
			towers[t1.name] = &t1
			// add to list
			//towers2.list = append(towers2.list, t1)
			continue
		}

		// case 2
		case2 := r2.FindStringSubmatch(scanner.Text())
		if case2 != nil {
			// get the children
			children := strings.Split(case2[3], ", ")
			// ct := make([]tower, 0, len(children))

			// for _, child := range children {
			// 	if ok, idx := towers2.inList(child); ok {
			// 		towers2.list[idx].parent = case2[1]
			// 		ct = append(ct, towers2.list[idx])
			// 	} else {
			// 		t1 := tower {
			// 			name: child,
			// 			parent: case2[1],
			// 		}
			// 		towers2.list = append(towers2.list, t1)
			// 		ct = append(ct, t1)
			// 	}
			// }
			weight := lib.ToInt(case2[2])
			t2 := tower{
				name:        case2[1],
				weight:      weight,
				totalWeight: weight,
				children:    children,
			}
			towers[t2.name] = &t2
			continue
		}

		// if reach here fatal
		fmt.Fprintf(os.Stderr, "No match found %v", scanner.Text())
		os.Exit(1)
	}

	// loop through if it has children update the respective parents
	for name, tow := range towers {
		//fmt.Println(name, tow)
		for _, child := range tow.children {
			t := towers[child]
			t.parent = name
			towers[child] = t

		}
	}

	//fmt.Printf("%+v\n", towers)
	return towers
}

func part1(towers map[string]*tower) (string, error) {
	for _, val := range towers {
		if val.parent == "" {
			fmt.Println("Part 1:", val.name)
			return val.name, nil
		}
	}
	return "", fmt.Errorf("No root node found")
}

// post order traversal of the tree updating the total weights
func postOrder(towers map[string]*tower, p string) {
	for _, child := range towers[p].children {
		postOrder(towers, child)
	}
	// visit action
	parent := towers[p].parent
	//fmt.Printf("%+v\n", p)
	if parent != "" {
		newWeight := towers[parent].totalWeight + towers[p].totalWeight
		//fmt.Println(towers[p].totalWeight)

		towers[parent].totalWeight = newWeight
	}
}

func main() {
	towers := load()
	root, err := part1(towers)
	if err != nil {
		panic(err)
	}
	// assign the total weights to each node
	postOrder(towers, root)

	// function that gets the different int from a slice
	diff := func(elems []tower) (bool, string, int) {
		occurence := make(map[int]int)
		tw := make(map[int]string)

		for _, t := range elems {
			occurence[t.totalWeight]++
			tw[t.totalWeight] = t.name
		}
		var required int
		var different string
		var diffExist bool
		for k, v := range occurence {
			if v == 1 {
				different = tw[k]
				diffExist = true
				//fmt.Println(tw[k])
			} else {
				required = towers[tw[k]].totalWeight
			}
		}
		return diffExist, different, required
		//fmt.Println("required", required)

	}

	// excuse to use my queue implementation
	q := lib.NewQueue[string]()
	q.Enqueue(root)

	for !q.IsEmpty() {
		elem, err := q.Dequeue()
		if err != nil {
			break
		}

		tw := []tower{}
		for _, c := range towers[elem].children {
			tw = append(tw, *towers[c])
		}
		diffExist, dif, _ := diff(tw)
		//fmt.Println(diffExist,dif, req)

		// the total weights of its children should be equal
		// if not enqueue the one not equal
		if diffExist {
			q.Enqueue(dif)
			continue
		}

		// all children are equal
		//fmt.Println("Dif here:", elem, towers[elem])
		tw2 := []tower{}
		culprit := towers[elem].parent
		//fmt.Println("Culprit, ",culprit)
		for _, c := range towers[culprit].children {
			tw2 = append(tw2, *towers[c])
		}
		_, _, req := diff(tw2)

		//fmt.Printf("req: %v, elem.totalweight: %v , elem.weight: %v\n", req, towers[elem].totalWeight, towers[elem].weight)

		fmt.Printf("Part 2: %v\n", towers[elem].weight+(req-towers[elem].totalWeight))
	}
}

package main

import (
	"container/heap"
	"fmt"
	"os"
)

type priorityQueue []*gameState

func (pq priorityQueue) Len() int {
	return len(pq)
}

// needed for the heap less and swap
func (pq priorityQueue) Less(i int, j int) bool {
	return pq[i].mannaSpent < pq[j].mannaSpent
}

func (pq priorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x any) {
	item := x.(*gameState)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type character struct {
	hitPoints, damage int
}

type gameState struct {
	player                          character
	boss                            character
	spellsUsed                      string
	manna                           int
	poison, shield, recharge, armor int
	mannaSpent                      int
}

func playRound(game gameState, nextMove string, hard bool) (g gameState, won bool, winner string, err error) {

	lessManna := func(val int) error {
		if val > game.manna {
			return fmt.Errorf("Not enough manna")
		}
		game.manna -= val
		return nil

	}

	less := func(val *int, diff int) {
		*val += diff
		if *val < 0 {
			*val = 0
		}

	}

	applyEffects := func() {
		if game.shield > 0 {
			game.armor = 7
			//shield -- // func if less
			less(&game.shield, -1)
			//fmt.Println("Shield dealt", shield)
			if game.shield == 0 {
				game.armor = 0
			}
		}
		if game.poison > 0 {
			game.boss.hitPoints -= 3
			less(&game.poison, -1)
			//poison -- // func
			//fmt.Println("poison dealt", poison)
		}
		if game.recharge > 0 {
			game.manna += 101
			//recharge -- //func
			less(&game.recharge, -1)
			//fmt.Println("recharge dealt", recharge)
		}

	}
	if hard {
		game.player.hitPoints -= 1

		if game.player.hitPoints <= 0 {
			// 1 boss won
			return game, true, "boss", nil
		}
		if game.boss.hitPoints <= 0 {
			//player won
			return game, true, "player", nil
		}
	}
	//fmt.Println("Boss hitpoints (player turn)", boss.hitPoints)
	//fmt.Println("Player hitpoints (player turn)", player.hitPoints, "armor", armor, "Manna", manna)
	// player turn apply effects
	applyEffects()

	if game.player.hitPoints <= 0 {
		// 1 boss won
		return game, true, "boss", nil
	}
	if game.boss.hitPoints <= 0 {
		//player won
		return game, true, "player", nil
	}

	switch nextMove {
	case "Magic Missile":
		game.boss.hitPoints -= 4

		err := lessManna(53)
		if err != nil {
			return game, false, "", err
		}
	case "Drain":
		err := lessManna(73)
		if err != nil {
			return game, false, "", err
		}

		game.boss.hitPoints -= 2
		game.player.hitPoints += 2
	case "Shield":
		err := lessManna(113)
		if err != nil {
			return game, false, "", err
		}

		if game.shield > 0 {
			// should return and error
			return game, false, "", fmt.Errorf("Invalid shield ")
		} else {
			game.shield = 6
			//player.armor += 7
		}
	case "Poison":
		err := lessManna(173)
		if err != nil {
			return game, false, "", err
		}

		if game.poison > 0 {
			// error
			return game, false, "", fmt.Errorf("Invalid shield ")
		} else {
			game.poison = 6
			//boss.hitPoints -= 3
		}
	case "Recharge":
		err := lessManna(229)
		if err != nil {
			return game, false, "", err
		}
		if game.recharge > 0 {
			return game, false, "", fmt.Errorf("Invalid Recharge")
		} else {
			game.recharge = 5
			//manna += 101
		}
	default:
		return game, false, "", fmt.Errorf("Invalid unknown %v", nextMove)

	}
	//fmt.Println("----------------------------------------------------------")
	//fmt.Println("Boss hitpoints (boss turn)",boss.hitPoints)
	//fmt.Println("Player hitpoints (boss turn)", player.hitPoints, "armor", armor, "Manna", manna)
	// boss turn apply effects
	applyEffects()

	if game.player.hitPoints <= 0 {
		// 1 boss won
		return game, true, "boss", nil
	}
	if game.boss.hitPoints <= 0 {
		//player won
		return game, true, "player", nil
	}

	// boss attacking
	damage := game.boss.damage
	if game.armor > 0 {
		//fmt.Println("Entered here playor armor", player.armor)
		damage = game.boss.damage - game.armor
		if damage <= 0 {
			damage = 1
		}
	}
	//fmt.Println("Boss deals", damage, "damage")
	game.player.hitPoints -= damage

	//fmt.Println()

	return game, false, "", nil

}

func run(init gameState, hard bool) (gameState, error) {

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &init)

	//spells := []string{"Magic Missile", "Drain", "Shield", "Poison", "Recharge"}
	spells := make(map[string]int)
	spells["Magic Missile"] = 53
	spells["Drain"] = 73
	spells["Shield"] = 113
	spells["Poison"] = 173
	spells["Recharge"] = 229

	for pq.Len() > 0 {
		// get state with lowest manna spent
		current := heap.Pop(pq).(*gameState)

		// try every spell add to priority queue the one that game does not end
		// if it ends and player wins thats it
		for spell, cost := range spells {
			next, won, winner, err := playRound(*current, spell, hard)

			if err != nil {
				continue
			}
			next.mannaSpent = current.mannaSpent + cost

			if won {
				if winner == "player" {
					return next, nil
				}
				continue // boss won do not add
			}
			//next.spellsUsed = append(next.spellsUsed, spell)
			next.spellsUsed = next.spellsUsed + " " + spell
			heap.Push(pq, &next) // only add games that did not finish

		}
	}

	return gameState{}, fmt.Errorf("Cannot find answer.") // won't be reached 

}

func main() {
	//moves := [] string {"Magic Missile", "Drain","Shield" ,"Poison", "Recharge"}
	p1 := character{50, 0}
	b1 := character{51, 9}
	manna := 500

	g := gameState{p1, b1, "" , manna, 0, 0, 0, 0, 0}

	a, e := run(g, false)
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
	fmt.Println("Part 1:", a.mannaSpent)
	fmt.Printf("\t%v\n", a.spellsUsed)

	b, er := run(g, true)
	if er != nil {
		fmt.Fprintln(os.Stderr, er)
		os.Exit(1)
	}
	fmt.Println("Part 2:", b.mannaSpent)
	fmt.Printf("\t%v\n", b.spellsUsed)


}

func t() {
	p1 := character{10, 0}
	b1 := character{13, 8}
	b2 := character{14, 8}
	manna := 250
	m1 := [...]string{"Poison", "Magic Missile"}
	m2 := [...]string{"Recharge", "Shield", "Drain", "Poison", "Magic Missile"}

	test(p1, b1, manna, m1[:])
	fmt.Printf("\n\n\n")
	test(p1, b2, manna, m2[:])

}

func test(player character, boss character, manna int, moves []string) {
	//g := gameState{player, boss, manna, 0, 0, 0, 0}

	m3 := make(map[string]int)
	m3["Magic Missile"] = 53
	m3["Drain"] = 73
	m3["Shield"] = 113
	m3["Poison"] = 173
	m3["Recharge"] = 229

	g := gameState{player, boss,"", manna, 0, 0, 0, 0, 0}
	//mannaSpent := math.MaxInt

	for _, move := range moves {

		fmt.Printf("%+v\n", g)
		fmt.Printf("Player casts '%v'\n", move)

		ga, won, winner, err := playRound(g, move, false)
		ga.mannaSpent += m3[move]

		if err != nil {
			fmt.Println("Error.", err)
			//return
		}
		if won {
			if winner == "player" {

			}
			fmt.Println("Winner: ", winner)
			return
		}
		g = ga
	}
	fmt.Printf("%+v\n", g)
}

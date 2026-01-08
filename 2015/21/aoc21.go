package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type character struct{
	hitPoints, damage, armor int
}

type item struct {
	category string
	description string
	cost, damage, armor int
}

func playGame(player character, boss character) string {

	round := func (attacker *character, defender *character) {
		damage := attacker.damage - defender.armor
		if damage <= 0 {
			damage = 1
		}
		defender.hitPoints -= damage
	}

	turn := "p"

	for {
		if turn == "p" {
			round(&player, &boss)
			turn = "b"
		} else {
			round(&boss, &player)
			turn = "p"
		}
		// TODO return true or false
		if player.hitPoints <= 0 {
			return "boss"
		}
		if boss.hitPoints <= 0 {
			return "player"
		}
	}
}

// read and populate the items in the shop
func populateShop() map[string][]item {
	// TODO process line by line 
	shop, err := os.ReadFile("shop.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	toint := func (str string) int {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		return num
	}
	items := make(map[string][]item)
	current := ""
	for line := range strings.SplitSeq(string(shop), "\n") {
		//fmt.Printf("here %v\n", line)
		if strings.Contains(line, ":") {
			// header. extract the name
			current = strings.Split(line, ":")[0]
			continue
		}
		if line == "" {
			// TODO how to check for an empty line this breaks if it has even a space
			// trimming the line
			continue
		}
		// kinda cheated here and removed spaces around rings ;)
		// create a struct from the line
		islice := strings.Fields(line)
		itemstr := item{}
		itemstr.category = current
		itemstr.description = islice[0]
		itemstr.cost = toint(islice[1])
		itemstr.damage = toint(islice[2])
		itemstr.armor = toint(islice[3])

		items[current] = append(items[current], itemstr)

	}
	return items

}

func buildCharacter(weapon item, armor item, ring item, ring2 item, hitPoints int) (character, int){
	all :=[...]item{weapon, armor, ring, ring2}
	var cost, damage, armorVal int
	for _, a := range all {
		cost += a.cost
		damage += a.damage
		armorVal += a.armor
	}

	// new character
	return character{hitPoints: hitPoints, damage: damage, armor: armorVal}, cost

}

func main() {
	hitpoints := 100
	//TODO populate boss from file
	boss := character{hitPoints: 104, damage: 8, armor: 1}

	items := populateShop()
	// for optional items make an empty option and add to the list
	emptyArmor := item{"Armor", "nil", 0, 0, 0}
	emptyRing := item{"Rings", "nil", 0, 0, 0}

	items["Armor"] = append(items["Armor"], emptyArmor)
	items["Rings"] = append(items["Rings"], emptyRing)


	min_cost := math.MaxInt
	max_cost := 0
	var best_play_min string
	var best_play_max string

	// loop through all
	for _, weapon := range items["Weapons"] {
		for _, armor := range items["Armor"] {
			for _, ring := range items["Rings"] {
				for _, ring2 := range items["Rings"] {
					if ring == ring2 {
						continue // cannot by two same rings but if both are empty possible
					}
					player, cost := buildCharacter(weapon, armor, ring, ring2, hitpoints)
					if playGame(player, boss) == "player" {
						if cost < min_cost {
							min_cost = cost
							best_play_min = weapon.description + 
											" " + armor.description + 
											" " + ring.description + 
											" " + ring2.description
						//best_play_min = 
						}
					} else {
						// part 2 player loses
						if cost > max_cost {
							max_cost = cost
							best_play_max = weapon.description + 
											" " + armor.description + 
											" " + ring.description + 
											" " + ring2.description
						}
					}

				}
			}
			//fmt.Println(weapon.description, armor.description)
		}
	}

	//fmt.Println(best_play)
	fmt.Println("Part 1:",min_cost)
	fmt.Println("\t", best_play_min)
	fmt.Println("Part 2:",max_cost)
	fmt.Println("\t", best_play_max)
}

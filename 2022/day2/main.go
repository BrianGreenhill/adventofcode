package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	getPairs := func() []pair {
		scanner := bufio.NewScanner(os.Stdin)
		pairs := pairs{}
		for scanner.Scan() {
			line := scanner.Text()
			ps := strings.Split(line, " ")
			p := pair{a: ps[0], b: ps[1]}
			pairs = append(pairs, p)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		return pairs
	}

	pairs := getPairs()

	if err := part1(pairs); err != nil {
		log.Fatal(err)
	}
	if err := part2(pairs); err != nil {

	}
}

var lookup = map[string]int{
	"A": 1,
	"X": 1,
	"B": 2,
	"Y": 2,
	"C": 3,
	"Z": 3,
}

const lost = 0
const draw = 3
const win = 6

type pairs []pair

type pair struct {
	a, b string
}

type player struct {
	id            string
	score         int
	currentChoice string
	turn          bool
}

func part2(pairs []pair) error {
	p1 := &player{id: "player"}
	p2 := &player{id: "opponent"}

	scoreRound := func(p1, p2 *player) {
		// p1 has a desired outcome
		// p2 has a choice
		// p1 has to respond to p2s choice such that they get the desired outcome
		lossLookup := map[string]int{"A": 3, "B": 1, "C": 2}
		winLookup := map[string]int{"A": 2, "B": 3, "C": 1}

		switch p1.currentChoice {
		case "X":
			p1.score += lost + lossLookup[p2.currentChoice]
			p2.score += win + lookup[p2.currentChoice]
		case "Y":
			p2.score += draw + lookup[p2.currentChoice]
			p1.score += draw + lookup[p2.currentChoice]
		case "Z":
			p2.score += lost + lookup[p2.currentChoice]
			p1.score += win + winLookup[p2.currentChoice]
		default:
			log.Fatal("invalid input")
		}
	}

	for _, p := range pairs {
		p1.currentChoice = p.b
		p2.currentChoice = p.a
		scoreRound(p1, p2)
	}

	fmt.Println()
	fmt.Println("part 2")
	fmt.Println(p1.id, p1.score)
	fmt.Println(p2.id, p2.score)
	return nil
}

func part1(pairs []pair) error {
	p1 := &player{id: "player", turn: false}
	p2 := &player{id: "opponent", turn: true}
	for _, p := range pairs {
		if p1.turn {
			p1.currentChoice = p.a
			p2.currentChoice = p.b
			p1.turn = false
			p2.turn = true
		}
		if p2.turn {
			p2.currentChoice = p.a
			p1.currentChoice = p.b
			p1.turn = true
			p2.turn = false
		}
		scoreRound(p1, p2)
		// fmt.Println("round", i+1)
		// fmt.Println(p1.id, p1.currentChoice, p1.score)
		// fmt.Println(p2.id, p2.currentChoice, p2.score)
		// fmt.Println("")
	}

	fmt.Println("part 1")
	fmt.Println(p1.id, p1.score)
	fmt.Println(p2.id, p2.score)
	return nil
}

func scoreRound(p, o *player) {
	// handle draw
	if lookup[p.currentChoice] == lookup[o.currentChoice] {
		p.score += draw + lookup[p.currentChoice]
		o.score += draw + lookup[o.currentChoice]
		//fmt.Println(p.id, "draw")
		return
	}

	// handle wins/losses
	switch p.currentChoice {
	case "A", "X":
		if o.currentChoice == "Y" || o.currentChoice == "B" {
			p.score += lost + lookup["A"]
			o.score += win + lookup["B"]
			//fmt.Println(p.id, "lost")
		} else {
			p.score += win + lookup["A"]
			o.score += lost + lookup["B"]
			//fmt.Println(p.id, "won")
		}
	case "B", "Y":
		if o.currentChoice == "C" || o.currentChoice == "Z" {
			p.score += lost + lookup["B"]
			o.score += win + lookup["C"]
			//fmt.Println(p.id, "won")
		} else {
			p.score += win + lookup["B"]
			o.score += lost + lookup["C"]
			//fmt.Println(p.id, "won")
		}
	case "C", "Z":
		if o.currentChoice == "A" || o.currentChoice == "X" {
			p.score += lost + lookup["C"]
			o.score += win + lookup["A"]
			//fmt.Println(p.id, "lost")
		} else {
			p.score += win + lookup["C"]
			o.score += lost + lookup["A"]
			//fmt.Println(p.id, "won")
		}
	default:
		log.Fatalf("not a valid choice %s", p.currentChoice)
	}
}

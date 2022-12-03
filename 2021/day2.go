package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func day2() int {
	c := parseInput()
	insts := getInstructions(c)
	pos := position{x: 0, y: 0}
	for _, i := range insts {
		pos.move(i)
	}
	fmt.Println("day2a:", pos.x*pos.y)
	return pos.x * pos.y
}

func day2b() int {
	c := parseInput()
	insts := getInstructions(c)
	pos := position{x: 0, y: 0}
	for _, i := range insts {
		pos.moveWithAim(i)
	}

	fmt.Println("day2b:", pos.x*pos.y)
	return pos.x * pos.y
}

type instruction struct {
	direction string
	dist      int
}

type position struct {
	x, y, aim int
}

func getInstructions(lines []string) []instruction {
	var insts []instruction
	for _, l := range lines {
		arr := strings.Split(l, " ")
		dist, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal(err)
		}

		insts = append(insts, instruction{
			direction: arr[0],
			dist:      dist,
		})
	}
	return insts
}

func (pos *position) move(i instruction) {
	switch i.direction {
	case "forward":
		pos.x += i.dist
		break
	case "up":
		pos.y -= i.dist
		break
	case "down":
		pos.y += i.dist
		break
	default:
		log.Fatal("invalid instruction")
	}
}

func (pos *position) moveWithAim(i instruction) {
	switch i.direction {
	case "forward":
		pos.x += i.dist
		pos.y += pos.aim * i.dist
		break
	case "up":
		pos.aim -= i.dist
		break
	case "down":
		pos.aim += i.dist
		break
	default:
		log.Fatal("invalid instruction")
	}
}

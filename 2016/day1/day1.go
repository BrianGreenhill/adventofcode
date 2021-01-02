package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	x, y int
	d    string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	pos := Coords{
		x: 0,
		y: 0,
		d: "N",
	}
	set := make(map[Coords]bool)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		split := strings.Split(line, ", ")
		for _, inst := range split {
			direction := string(inst[0])
			steps, _ := strconv.Atoi(inst[1:])
			if pos.d == "N" {
				if direction == "R" {
					pos.d = "E"
					pos.x += steps
				} else if direction == "L" {
					pos.d = "W"
					pos.x -= steps
				}
			} else if pos.d == "E" {
				if direction == "R" {
					pos.d = "S"
					pos.y -= steps
				}
				if direction == "L" {
					pos.d = "N"
					pos.y += steps
				}
			} else if pos.d == "W" {
				if direction == "R" {
					pos.d = "N"
					pos.y += steps
				}
				if direction == "L" {
					pos.d = "S"
					pos.y -= steps
				}
			} else if pos.d == "S" {
				if direction == "R" {
					pos.d = "W"
					pos.x -= steps
				}
				if direction == "L" {
					pos.d = "E"
					pos.x += steps
				}
			}
			if _, ok := set[pos]; ok {
				fmt.Println("Part 2", pos)
				break
			}
			set[pos] = true
			fmt.Println("x:", pos.x, "y:", pos.y)
		}
	}
	// ans := float64(pos.x) + float64(pos.y)
	fmt.Printf("%+v\n", pos)
	// fmt.Println(ans)
}

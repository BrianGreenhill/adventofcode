package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Coord struct {
	x, y int
}

func main() {
	r := regexp.MustCompile("(turn off|toggle|turn on) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)")
	reader := bufio.NewReader(os.Stdin)
	var on map[Coord]bool = make(map[Coord]bool)
	var bright map[Coord]int = make(map[Coord]int)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		parsed := r.FindStringSubmatch(line)
		command := parsed[1]
		x1, _ := strconv.Atoi(parsed[2])
		y1, _ := strconv.Atoi(parsed[3])
		x2, _ := strconv.Atoi(parsed[4])
		y2, _ := strconv.Atoi(parsed[5])
		switch command {
		case "turn off":
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					delete(on, Coord{x, y})
					if bright[Coord{x, y}] >= 1 {
						bright[Coord{x, y}]--
					} else {
						delete(bright, Coord{x, y})
					}
				}
			}
		case "turn on":
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					on[Coord{x, y}] = true
					bright[Coord{x, y}]++
				}
			}
		case "toggle":
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					bright[Coord{x, y}] += 2
					if on[Coord{x, y}] {
						delete(on, Coord{x, y})
					} else {
						on[Coord{x, y}] = true
					}
				}
			}
		}
	}
	fmt.Println("Part 1", len(on))
	sum := 0
	for _, v := range bright {
		sum += v
	}
	fmt.Println("Part 2", sum)
}

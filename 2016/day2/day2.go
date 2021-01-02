package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coord struct {
	X, Y int
}

var pin int

func main() {
	reader := bufio.NewReader(os.Stdin)
	numPad := map[coord]rune{
		coord{-1, -1}: '7',
		coord{0, -1}:  '8',
		coord{1, -1}:  '9',
		coord{-1, 0}:  '4',
		coord{0, 0}:   '5',
		coord{1, 0}:   '6',
		coord{-1, 1}:  '1',
		coord{0, 1}:   '2',
		coord{1, 1}:   '3',
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		input := strings.Split(line, "\n")
		for _, key := range input {
			loc := coord{0, 0}
			for _, inst := range key {
				proposedLoc := loc
				switch inst {
				case 'U':
					proposedLoc.Y += 1
				case 'D':
					proposedLoc.Y -= 1
				case 'L':
					proposedLoc.X -= 1
				case 'R':
					proposedLoc.X += 1
				default:
					fmt.Printf("Invalid direction %q\n", inst)
					return
				}
				if numPad[proposedLoc] != rune(0) {
					loc = proposedLoc
				}
			}
			fmt.Printf("%q", numPad[loc])
		}
		fmt.Println()
	}
}

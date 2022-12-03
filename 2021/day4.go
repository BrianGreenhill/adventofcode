package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type markable struct {
	val    string
	marked bool
}

type board struct {
	grid      [5][5]markable
	score     int
	lastDigit int
	won       bool
}

func day4b() int {
	instructions, boards := prepareData()
	var winningBoards []board
	for _, i := range instructions {
		for _, board := range boards {
			board.selectDigit(i)
			if board.hasWon() {
				board.won = true
				winningBoards = append(winningBoards, *board)
			}
		}
	}
	return winningBoards[len(winningBoards)-1].sumUnmarked() * winningBoards[len(winningBoards)-1].lastDigit
}

func day4() int {
	instructions, boards := prepareData()
	var winningBoards []*board

	for _, i := range instructions {
		for _, board := range boards {
			board.selectDigit(i)
			if board.hasWon() {
				winningBoards = append(winningBoards, board)
				return winningBoards[0].sumUnmarked() * winningBoards[0].lastDigit
			}
		}
	}
	return 0
}

func (b *board) sumUnmarked() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.grid[i][j].marked {
				continue
			}
			val, err := strconv.Atoi(b.grid[i][j].val)
			if err != nil {
				log.Fatal(err)
			}
			sum += val
		}
	}
	return sum
}

func (b *board) hasWon() bool {
	if b.won {
		return false
	}
	for i := 0; i < 5; i++ {
		count := 0
		for j := 0; j < 5; j++ {
			if b.grid[j][i].marked {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		count := 0
		// row
		for j := 0; j < 5; j++ {
			if b.grid[i][j].marked {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}
	return false
}

func (b *board) selectDigit(digit string) {
	// have to add logic to check for a winner each time
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.grid[i][j].val == digit {
				val, err := strconv.Atoi(digit)
				if err != nil {
					log.Fatal(err)
				}
				if !b.grid[i][j].marked {
					b.grid[i][j].marked = true
					b.score += val
					b.lastDigit = val
				}
			}
		}
	}
}

func prepareData() ([]string, map[int]*board) {
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	contents := string(bytes)
	c := strings.Split(contents, "\n")
	instructions := strings.Split(c[0], ",")

	return instructions, prepareBoards(c[2:])
}

func prepareBoards(c []string) map[int]*board {
	var prepared []string
	for _, line := range c {
		if line == "" {
			continue
		}
		prepared = append(prepared, line)
	}
	boards := make(map[int]*board)
	batch := 5
	for i := 0; i < len(prepared); i += batch {
		j := i + batch
		if j > len(c) {
			j = len(c)
		}
		boards[(j/batch)-1] = createBoard(prepared[i:j])
	}
	return boards
}

func createBoard(c []string) *board {
	var markableVals [5][5]markable
	var ret board
	for i := 0; i < 5; i++ {
		vals := strings.Split(c[i], " ")
		var trimmed []string
		for _, v := range vals {
			if v == "" {
				continue
			}
			trimmed = append(trimmed, strings.TrimPrefix(v, " "))
		}
		for m := 0; m < 5; m++ {
			markableVals[i][m] = markable{val: trimmed[m], marked: false}
		}
		ret = board{grid: markableVals, score: 0}
	}
	return &ret
}

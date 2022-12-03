package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type bounds struct {
	x, y int
}

type fissure struct {
	start bounds
	end   bounds
}

func makeRange(min, max int) []int {
	newMin := int(math.Min(float64(min), float64(max)))
	newMax := int(math.Max(float64(min), float64(max)))
	a := make([]int, newMax-newMin+1)
	for i := range a {
		a[i] = newMin + i
	}
	return a
}

func day5b() int {
	data := parseData()
	var upperX, upperY float64
	// consider only horizontal and vertical fissures
	var fissures []*fissure
	for _, f := range data {
		maxX := math.Max(float64(f.start.x), float64(f.end.x))
		if maxX > upperX {
			upperX = maxX
		}

		maxY := math.Max(float64(f.start.y), float64(f.end.y))
		if maxY > upperY {
			upperY = maxY
		}
		fissures = append(fissures, f)
	}
	danger := 0
	yBound := int(upperY) + 1
	xBound := int(upperX) + 1
	grid := make([][]int, yBound)
	for i := range grid {
		grid[i] = make([]int, xBound)
	}
	for _, f := range fissures {
		// vertical
		if f.start.x == f.end.x {
			for _, l := range makeRange(f.start.y, f.end.y) {
				grid[l][f.start.x]++
			}
		} else if f.start.y == f.end.y {
			for _, l := range makeRange(f.start.x, f.end.x) {
				grid[f.start.y][l]++
			}
		} else {
			dx := int(float64(f.end.x-f.start.x) / math.Abs(float64(f.end.x-f.start.x)))
			dy := int(float64(f.end.y-f.start.y) / math.Abs(float64(f.end.y-f.start.y)))
			x := f.start.x
			y := f.start.y

			for y != f.end.y+dy {
				grid[y][x]++
				x += dx
				y += dy
			}
		}
	}
	for i := 0; i < yBound; i++ {
		for j := 0; j < xBound; j++ {
			if grid[i][j] >= 2 {
				danger++
			}
		}
	}
	return danger
}

func day5() int {
	data := parseData()
	var upperX, upperY float64

	// consider only horizontal and vertical fissures
	var fissures []*fissure
	for _, f := range data {
		if f.start.x != f.end.x && f.start.y != f.end.y {
			continue
		}
		maxX := math.Max(float64(f.start.x), float64(f.end.x))
		if maxX > upperX {
			upperX = maxX
		}

		maxY := math.Max(float64(f.start.y), float64(f.end.y))
		if maxY > upperY {
			upperY = maxY
		}
		fissures = append(fissures, f)
	}
	danger := 0
	yBound := int(upperY) + 1
	xBound := int(upperX) + 1
	grid := make([][]int, yBound)
	for i := range grid {
		grid[i] = make([]int, xBound)
	}
	for _, f := range fissures {
		if f.start.x == f.end.x {
			for _, l := range makeRange(f.start.y, f.end.y) {
				grid[l][f.start.x]++
			}
		}
		if f.start.y == f.end.y {
			for _, l := range makeRange(f.start.x, f.end.x) {
				grid[f.start.y][l]++
			}
		}
	}
	for i := 0; i < yBound; i++ {
		for j := 0; j < xBound; j++ {
			if grid[i][j] >= 2 {
				danger++
			}
		}
	}
	return danger
}

func parseData() []*fissure {
	data := parseInput()
	// var upperX, upperY float64
	var fissures []*fissure
	for _, item := range data {
		items := strings.Split(item, " -> ")
		length := len(items)
		for i := 0; i < length; i += 2 {
			j := i + 2
			if j > length {
				j = length
			}
			f1 := items[i:j][0]
			f2 := items[i:j][1]
			x, err := strconv.Atoi(strings.Split(f1, ",")[0])
			y, err := strconv.Atoi(strings.Split(f1, ",")[1])
			x2, err := strconv.Atoi(strings.Split(f2, ",")[0])
			y2, err := strconv.Atoi(strings.Split(f2, ",")[1])

			if err != nil {
				log.Fatal(err)
			}
			fissures = append(fissures, &fissure{
				start: bounds{
					x: x,
					y: y,
				},
				end: bounds{
					x: x2,
					y: y2,
				},
			})
		}
	}
	return fissures
}

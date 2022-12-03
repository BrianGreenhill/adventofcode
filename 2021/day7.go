package main

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func day7() int {
	data := strings.Split(parseInput()[0], ",")
	var positions []int
	for _, pos := range data {
		a, err := strconv.Atoi(pos)
		if err != nil {
			log.Fatal(err)
		}
		positions = append(positions, a)
	}

	// the smallest sum of the differences to each number is the shortest path
	diffs := make([]int, len(positions))
	for i, pos := range positions {
		for j := 0; j < len(positions); j++ {
			diffs[i] += int(math.Abs(float64(pos - positions[j])))
		}
	}
	sort.Ints(diffs)
	return diffs[0]
}

func day7b() int {
	data := strings.Split(parseInput()[0], ",")
	var positions []int
	for _, pos := range data {
		a, err := strconv.Atoi(pos)
		if err != nil {
			log.Fatal(err)
		}
		positions = append(positions, a)
	}

	diffs := make([]int, len(positions)+1)
	for i := 0; i < len(diffs); i++ {
		for j := 0; j < len(positions); j++ {
			diffs[i] += cost(int(math.Abs(float64(i - positions[j]))))
		}
	}
	sort.Ints(diffs)
	return diffs[0]
}

func cost(i int) int {
	return (i * (i + 1)) / 2
}

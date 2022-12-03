package main

import (
	"log"
	"strconv"
	"strings"
)

const (
	A_DAYS = 80
	B_DAYS = 256

	REPRODUCTIONRATE = 6
)

func day6() int {
	data := strings.Split(strings.Join(parseInput(), ","), ",")

	fish := make([]int, 9)
	for _, row := range data {
		a, err := strconv.Atoi(row)
		if err != nil {
			log.Fatal(err)
		}
		fish[a]++
	}

	for i := 0; i < A_DAYS; i++ {
		fish = step(fish)
	}

	sum := 0
	for _, v := range fish {
		sum += v
	}
	return sum
}

func day6b() int {
	data := strings.Split(strings.Join(parseInput(), ","), ",")

	fish := make([]int, 9)
	for _, row := range data {
		a, err := strconv.Atoi(row)
		if err != nil {
			log.Fatal(err)
		}
		fish[a]++
	}

	for i := 0; i < B_DAYS; i++ {
		fish = step(fish)
	}

	sum := 0
	for _, v := range fish {
		sum += v
	}
	return sum
}

func step(fish []int) []int {
	next := make([]int, 9)
	for i := 1; i < 9; i++ {
		next[i-1] = fish[i]
	}
	next[REPRODUCTIONRATE] += fish[0]
	next[REPRODUCTIONRATE+2] += fish[0]
	return next
}

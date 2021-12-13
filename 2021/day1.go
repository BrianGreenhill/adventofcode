package main

import (
	"log"
	"strconv"
)

func day1() int {
	c := parseInput()
	count := 0
	for i := 1; i < len(c)-1; i++ {
		curr := c[i]
		prev := c[i-1]
		if larger(curr, prev) {
			count++
		}
	}

	return count
}

func day1b() int {
	c := parseInput()
	count := 0
	var vals []int
	for _, item := range c {
		intVal, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal("invalid int string")
		}
		vals = append(vals, intVal)
	}

	for i := 0; i < len(vals)-1; i++ {
		a := vals[i] + vals[i+1] + vals[i+2]
		b := vals[i+1] + vals[i+2] + vals[i+3]
		if b > a {
			count++
		}
		vals = RemoveIndex(vals, i)
		if len(vals) < 6 {
			return count
		}
		i--
	}

	return count
}

func larger(a, b string) bool {
	aVal, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	bVal, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}
	return aVal > bVal
}

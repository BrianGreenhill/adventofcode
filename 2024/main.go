package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if err := day1(os.Args); err != nil {
		panic(err)
	}
}

func day1(in []string) error {
	c, err := readInput(in)
	if err != nil {
		return err
	}

	left := []float64{}
	right := []float64{}
	lines := strings.Split(string(c), "\n")
	for _, l := range lines {
		if l == "\n" || l == "" {
			continue
		}
		lr := strings.Split(strings.ReplaceAll(l, "   ", ","), ",")
		l, err := strconv.Atoi(lr[0])
		if err != nil {
			return err
		}
		r, err := strconv.Atoi(lr[1])
		if err != nil {
			return err
		}

		left = append(left, float64(l))
		right = append(right, float64(r))
	}

	sort.Float64s(left)
	sort.Float64s(right)

	total := 0.0
	simScore := 0.0

	for i := 0; i < len(left); i++ {
		appearsCount := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				appearsCount++
			}
		}
		simScore += left[i] * float64(appearsCount)
		total += math.Abs(left[i] - right[i])
	}

	// similarity score
	for i := 0; i < len(left); i++ {
	}

	fmt.Println("a", int(total))
	fmt.Println("b", int(simScore))
	return nil
}

func readInput(in []string) ([]byte, error) {
	if len(in) < 2 {
		return nil, fmt.Errorf("missing input path")
	}

	c, err := os.ReadFile(in[1])
	if err != nil {
		return nil, err
	}

	return c, nil
}

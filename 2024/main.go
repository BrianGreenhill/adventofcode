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

	// day 1 - a
	total := 0.0
	for i := 0; i < len(left); i++ {
		total += math.Abs(left[i] - right[i])
	}

	// day 1 - b
	simScore := 0.0
	frequencyMap := map[float64]int{}
	for _, value := range right {
		frequencyMap[value]++
	}

	for _, value := range left {
		if count, exists := frequencyMap[value]; exists {
			simScore += value * float64(count)
		}
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

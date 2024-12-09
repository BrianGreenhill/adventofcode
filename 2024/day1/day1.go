package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	left := []float64{}
	right := []float64{}
	lines := strings.Split(string(input), "\n")
	for _, l := range lines {
		if l == "\n" || l == "" {
			continue
		}
		lr := strings.Split(strings.ReplaceAll(l, "   ", ","), ",")
		l, err := strconv.Atoi(lr[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(lr[1])
		if err != nil {
			panic(err)
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
}

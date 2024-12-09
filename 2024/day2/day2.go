package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	if err := day2(); err != nil {
		panic(err)
	}
}

func day2() error {
	reports := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		levels := []int{}
		for _, num := range strings.Fields(line) {
			n, err := strconv.Atoi(num)
			if err != nil {
				return err
			}
			levels = append(levels, n)
		}
		reports = append(reports, levels)
	}

	safeReportsA := 0
	safeReportsB := 0

	for _, report := range reports {
		if len(report) == 0 {
			continue
		}
		if checkIsSafe(report) {
			// Safe for day 2 - a
			safeReportsA++
			safeReportsB++ // Safe for day 2 - b as well
		} else if checkSafeWithRemoval(report) {
			// Safe only for day 2 - b
			safeReportsB++
		}
	}

	fmt.Println("day 2 - a", safeReportsA)
	fmt.Println("day 2 - b", safeReportsB)
	return nil
}

// checkIsSafe verifies if a report is safe without removing any levels.
func checkIsSafe(r []int) bool {
	allIncreasing := true
	allDecreasing := true
	validDifferences := true

	for i := 0; i < len(r)-1; i++ {
		diff := r[i+1] - r[i]
		if diff > 0 {
			allDecreasing = false
		} else if diff < 0 {
			allIncreasing = false
		}

		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			validDifferences = false
		}
	}

	return (allIncreasing || allDecreasing) && validDifferences
}

// checkSafeWithRemoval verifies if removing one level makes the report safe.
func checkSafeWithRemoval(r []int) bool {
	for i := 0; i < len(r); i++ {
		// Create a new slice excluding the current level
		modified := append([]int{}, r[:i]...)
		modified = append(modified, r[i+1:]...)

		// Check if the modified report is safe
		if checkIsSafe(modified) {
			return true
		}
	}
	return false
}

package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var debug = false

func main() {
	_ = day1(os.Args)
	_ = day2(os.Args)
	if err := day3(os.Args); err != nil {
		panic(err)
	}
}

func day3(in []string) error {
	testTarget := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	c, err := readInput(in)
	if err != nil {
		return err
	}

	if debug {
		// convert to byte slice
		b := []byte{}
		b = append(b, []byte(testTarget)...)
		c = b
	}
	res := 0
	exp := `mul\((\d{1,3}),(\d{1,3})\)`
	r := regexp.MustCompile(exp)
	match := r.FindAllStringSubmatch(string(c), -1)
	for _, m := range match {
		int1, err := strconv.Atoi(m[1])
		if err != nil {
			return err
		}
		int2, err := strconv.Atoi(m[2])
		if err != nil {
			return err
		}

		res += int1 * int2
	}
	fmt.Println("day 3 - a", res)

	res = 0
	exp = `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	r = regexp.MustCompile(exp)
	match = r.FindAllStringSubmatch(string(c), -1)
	do := true
	for _, m := range match {
		if m[0] == "do()" || m[0] == "don't()" {
			do = m[0] == "do()"
			continue
		}

		if !do {
			continue
		}

		int1, err := strconv.Atoi(m[1])
		if err != nil {
			return err
		}
		int2, err := strconv.Atoi(m[2])
		if err != nil {
			return err
		}

		res += int1 * int2
	}
	fmt.Println("day 3 - b", res)

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

func day2(in []string) error {
	c, err := readInput(in)
	if err != nil {
		return err
	}

	reports := [][]int{}
	for _, line := range strings.Split(string(c), "\n") {
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

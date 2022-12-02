package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fr, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	fb, err := io.ReadAll(fr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("part 1", runA(string(fb)))
	fmt.Println("part 2", runB(string(fb)))
}

func runB(data string) string {
	elves, totalCalories := parseInput(strings.Split(data, "\n"))
	total := 0
	for i := 0; i < 3; i++ {
		total += totalCalories[elves[i]]
	}

	return strconv.Itoa(total)
}

func runA(data string) string {
	elves, totalCalories := parseInput(strings.Split(data, "\n"))

	winningElf := elves[0]
	answer := 0
	if val, ok := totalCalories[elves[0]]; ok {
		answer = val
	}

	return fmt.Sprintf("elf %d is carrying %d calories", winningElf, answer)
}

func parseInput(data []string) ([]int, map[int]int) {
	// map elf to total calories
	totalCalories := map[int]int{}
	elf := 0
	for _, e := range data {
		if e == "\n" || e == "" {
			elf++
			continue
		}
		cals := strings.Split(e, "\n")
		for _, c := range cals {
			num, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			if _, ok := totalCalories[elf]; !ok {
				totalCalories[elf] = num
				continue
			}
			totalCalories[elf] += num
		}
	}

	keys := make([]int, 0)
	for k := range totalCalories {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	sort.SliceStable(keys, func(i, j int) bool {
		return totalCalories[keys[i]] > totalCalories[keys[j]]
	})
	return keys, totalCalories
}

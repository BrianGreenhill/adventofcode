package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var part = flag.String("part", "a", "Part a | b")

func main() {
	var bytes, _ = ioutil.ReadFile("inputb.simple")
	flag.Parse()
	fmt.Println("Day 1 - Part", *part, ":", Solve(string(bytes), *part))
}

func Solve(line, part string) int {
	nums := strings.Split(line, "\n")
	ans := 0
	for _, num := range nums[:len(nums)-1] {
		s := strings.Fields(num)
		seen := make([]int, len(s))
		max := 0
		min := 0
		for i, d := range s {
			n, err := strconv.Atoi(d)
			seen[i] = n
			if err != nil {
				fmt.Printf("couldnt parse %s, because %v.\n", num, err)
				return 0
			}
			if n > max {
				max = n
			}
			if n < min || min == 0 {
				min = n
			}
			if part == "b" {
				for j := 0; j < i; j++ {
					dividend := -1
					divisor := -1
					if seen[j] < n {
						dividend = n
						divisor = seen[j]
					} else {
						dividend = seen[j]
						divisor = n
					}
					if dividend%divisor == 0 {
						ans += dividend / divisor
					}
				}

			}
		}
		if part == "a" {
			ans += (max - min)
		}
	}
	return ans
}

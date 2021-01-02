package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var part = flag.String("part", "a", "Part a | b")

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	flag.Parse()
	fmt.Println("Day 1", Solve(line, *part))
}

func Solve(input string, part string) int {
	vals := strings.Split(input, "")
	vals = vals[:len(vals)-1]
	sum := 0
	for i := 0; i < len(vals); i++ {
		currVal, err := strconv.Atoi(strings.TrimSuffix(vals[i], "\n"))
		if err != nil {
			fmt.Printf("Could not parse digit %v\n", err)
			break
		}
		if part == "b" {
			halfway, err := strconv.Atoi(strings.TrimSuffix(vals[(len(vals)/2+i)%len(vals)], "\n"))
			if err != nil {
				fmt.Printf("Could not parse digit %v\n", err)
				break
			}
			if halfway == currVal {
				sum += currVal
			}
		} else {
			if i == 0 {
				prev, err := strconv.Atoi(strings.TrimSuffix(vals[len(vals)-1], "\n"))
				if err != nil {
					fmt.Printf("Could not parse digit %v\n", err)
					break
				}
				if prev == currVal {
					sum += currVal
				}
			} else {
				prev, err := strconv.Atoi(strings.TrimSuffix(vals[i-1], "\n"))
				if err != nil {
					fmt.Printf("Could not parse digit %v\n", err)
					break
				}
				if currVal == prev {
					sum += currVal
				}
			}
		}
	}
	return sum
}

package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input
var input string

func main() {
	if err := day3(); err != nil {
		panic(err)
	}
}
func day3() error {
	res := 0
	exp := `mul\((\d{1,3}),(\d{1,3})\)`
	r := regexp.MustCompile(exp)
	match := r.FindAllStringSubmatch(string(input), -1)
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
	match = r.FindAllStringSubmatch(string(input), -1)
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

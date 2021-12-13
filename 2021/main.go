package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "input", "relative input file")

func main() {
	flag.Parse()
	day8()
}

func parseInput() []string {
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return nil
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	return split[:len(split)-1]
}

type signalpattern struct {
	patterns, output []string
}

func day8() int {
	var signalpatterns []*signalpattern
	for _, row := range parseInput() {
		patterns := strings.Split(row, " | ")
		signalpatterns = append(signalpatterns, &signalpattern{
			patterns: strings.Split(patterns[0], " "),
			output:   strings.Split(patterns[1], " "),
		})
	}
	letters := make(map[int]map[int][]string, 7)
	for i, pattern := range signalpatterns {
		letters[i] = map[int][]string{}
		for _, out := range pattern.output {
			letters[i][len(out)] = append(letters[i][len(out)], out)
		}
	}
	// signals := make(map[int]map[int][]string, 7)
	// for i, pattern := range signalpatterns {
	// 	signals[i] = map[int][]string{}
	// 	for _, p := range pattern.patterns {
	// 		signals[i][len(p)] = append(signals[i][len(p)], p)
	// 	}
	// }
	ones := 0
	fours := 0
	sevens := 0
	eights := 0
	for i := 0; i < len(signalpatterns); i++ {
		ones += len(letters[i][2])
		sevens += len(letters[i][3])
		fours += len(letters[i][4])
		eights += len(letters[i][7])
	}

	fmt.Println(ones + fours + sevens + eights)
	return 0
}

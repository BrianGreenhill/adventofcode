package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	ans      int
	name     string
	input    string
	testFunc func() int
}

var testCases = map[string]testcase{
	"1a":     {name: "day1a", testFunc: day1, ans: 1400, input: "day1.input"},
	"1b":     {name: "day1b", testFunc: day1b, ans: 1429, input: "day1.input"},
	"2a":     {name: "day2a", testFunc: day2, ans: 2322630, input: "day2.input"},
	"2b":     {name: "day2b", testFunc: day2b, ans: 2105273490, input: "day2.input"},
	"3atest": {name: "day3a test", testFunc: day3, ans: 198, input: "day3.test"},
	"3btest": {name: "day3b test", testFunc: day3b, ans: 230, input: "day3.test"},
	"3a":     {name: "day3a", testFunc: day3, ans: 749376, input: "day3.input"},
	"3b":     {name: "day3b", testFunc: day3b, ans: 2372923, input: "day3.input"},
	"4atest": {name: "day4a test", testFunc: day4, ans: 4512, input: "day4.test"},
	"4btest": {name: "day4b test", testFunc: day4b, ans: 1924, input: "day4.test"},
	"4a":     {name: "day4a", testFunc: day4, ans: 11774, input: "day4.input"},
	"4b":     {name: "day4b", testFunc: day4b, ans: 4495, input: "day4.input"},
	"5atest": {name: "day5a test", testFunc: day5, ans: 5, input: "day5.test"},
	"5a":     {name: "day5a", testFunc: day5, ans: 4728, input: "day5.input"},
	"5btest": {name: "day5b test", testFunc: day5b, ans: 12, input: "day5.test"},
	"5b":     {name: "day5b", testFunc: day5b, ans: 17717, input: "day5.input"},
	"6atest": {name: "day6a test", testFunc: day6, ans: 5934, input: "day6.test"},
	"6a":     {name: "day6a", testFunc: day6, ans: 379114, input: "day6.input"},
	"6btest": {name: "day6b test", testFunc: day6b, ans: 26984457539, input: "day6.test"},
	"6b":     {name: "day6b", testFunc: day6b, ans: 1702631502303, input: "day6.input"},
	"7atest": {name: "day7a test", testFunc: day7, ans: 37, input: "day7.test"},
	"7a":     {name: "day7a", testFunc: day7, ans: 340056, input: "day7.input"},
	"7btest": {name: "day7b test", testFunc: day7b, ans: 168, input: "day7.test"},
	"7b":     {name: "day7b", testFunc: day7b, ans: 96592275, input: "day7.input"},
}

func Test2021(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			inputFile = &tt.input
			assert.Equal(t, tt.ans, tt.testFunc())
		})
	}
}

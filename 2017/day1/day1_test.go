package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bytes, _ = ioutil.ReadFile("input")
var day1Tests = []struct {
	name     string
	in       string
	part     string
	expected int
}{
	{"Part A 1122", "1122\n", "a", 3},
	{"Part A 1111", "1111\n", "a", 4},
	{"Part A 1234", "1234\n", "a", 0},
	{"Part A 91212129", "91212129\n", "a", 9},
	{"Part A real", string(bytes), "a", 1102},
	{"Part B 1212", "1212\n", "b", 6},
	{"Part B 1221", "1221\n", "b", 0},
	{"Part B 123425", "123425\n", "b", 4},
	{"Part B 123123", "123123\n", "b", 12},
	{"Part B 12131415", "12131415\n", "b", 4},
	{"Part B real", string(bytes), "b", 1076},
}

func TestDay1(t *testing.T) {
	for _, tt := range day1Tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Solve(tt.in, tt.part)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

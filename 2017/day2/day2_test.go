package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bytes, _ = ioutil.ReadFile("input")
var simplebytes, _ = ioutil.ReadFile("input.simple")
var simplebytesB, _ = ioutil.ReadFile("inputb.simple")

var dayTests = []struct {
	name     string
	in       string
	part     string
	expected int
}{
	{"Simple A", string(simplebytes), "a", 18},
	{"Real A", string(bytes), "a", 32020},
	{"Simple B", string(simplebytesB), "b", 9},
	{"Real B", string(bytes), "b", 236},
}

func TestDay(t *testing.T) {
	for _, tt := range dayTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Solve(tt.in, tt.part)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

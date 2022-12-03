package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func day3() int {
	c := parseInput()
	// assume all binary strings are of the same length
	binaryStringLength := len(c[0])

	// populate map with row to frequencies of 1 and 0
	// eg:  [{0: {ones: 3, zeroes: 2}}, {1: {ones: 4, zeroes: 7}}]
	m := make(map[int]binaryFrequency, binaryStringLength)
	for _, val := range c {
		vals := strings.Split(val, "")
		for i, v := range vals {
			f := m[i]
			switch v {
			case "1":
				f.ones += 1
			case "0":
				f.zeroes += 1
			}
			m[i] = f
		}
	}

	// create binary strings based on frequencies
	// gamma rate = *most* common bit in the corresponding position of all numbers
	// epsilon rate = *least* ""
	gamma := make(map[int]string, binaryStringLength)
	epsilon := make(map[int]string, binaryStringLength)
	for row, freq := range m {
		if freq.ones > freq.zeroes {
			gamma[row] = "1"
			epsilon[row] = "0"
		}
		if freq.zeroes > freq.ones {
			epsilon[row] = "1"
			gamma[row] = "0"
		}
	}
	var g, e string
	for i := 0; i < binaryStringLength; i++ {
		g += gamma[i]
		e += epsilon[i]
	}

	// convert binary strings to decimal
	gammaDec, err := strconv.ParseInt(g, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsDec, err := strconv.ParseInt(e, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	// return power consumption
	return int(gammaDec * epsDec)
}

func day3b() int {
	c := parseInput()
	// assume all binary strings are of the same length
	binaryStringLength := len(c[0])

	oxGenBin := ""
	for i := 0; i < binaryStringLength; i++ {
		needle := getNeedle(c, i)
		for l := 0; l < len(c); l++ {
			curr := strings.Split(c[l], "")[i]
			if curr != needle {
				c = RemoveIndexStr(c, l)
				if len(c) == 1 {
					oxGenBin = c[0]
				}
				l--
			}
		}
	}

	c = parseInput()
	co2ScrubBin := ""
	for i := 0; i < binaryStringLength; i++ {
		needle := getNeedle(c, i)
		if needle == "1" {
			needle = "0"
		} else {
			needle = "1"
		}
		for l := 0; l < len(c); l++ {
			curr := strings.Split(c[l], "")[i]
			if curr != needle {
				c = RemoveIndexStr(c, l)
				if len(c) == 1 {
					co2ScrubBin = c[0]
				}
				l--
			}
		}
	}

	oRating, err := strconv.ParseInt(oxGenBin, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	cRating, err := strconv.ParseInt(co2ScrubBin, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("day3b", int(oRating*cRating))
	return int(oRating * cRating)
}

type binaryFrequency struct {
	ones, zeroes int
}

func getNeedle(c []string, digitPos int) string {
	ones := 0
	zeroes := 0
	for j := 0; j < len(c); j++ {
		curr := strings.Split(c[j], "")[digitPos]
		if curr == "1" {
			ones++
		} else {
			zeroes++
		}
	}
	needle := "1"
	if zeroes > ones {
		needle = "0"
	}
	return needle
}

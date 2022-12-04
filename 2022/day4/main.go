package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// sections have ID number
	// every elf is assigned a range of section IDs
	// there is overlap in ID ranges so the elves pair up
	// data set is a list of section assignments for each pair
	data := []string{}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		data = append(data, sc.Text())
	}

	fmt.Println("part 1")
	if err := part1(data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("part 2")
	if err := part2(data); err != nil {
		log.Fatal(err)
	}
}

type section struct {
	start, end int
	id         string
}

func newSection(id string) *section {
	s := &section{
		id: id,
	}
	// parse start and end from id
	sections := strings.Split(id, "-")
	s.start, _ = strconv.Atoi(sections[0])
	s.end, _ = strconv.Atoi(sections[1])
	return s
}

type pair struct {
	section []*section
}

func (p *pair) isOverlap() bool {
	return p.section[0].end >= p.section[1].start && p.section[0].start <= p.section[1].end
}

func (p *pair) isFullOverlap() bool {
	res := false
	if p.section[0].start <= p.section[1].start {
		if p.section[0].end >= p.section[1].end {
			res = true
		}
	}

	if p.section[1].start <= p.section[0].start {
		if p.section[1].end >= p.section[0].end {
			res = true
		}
	}
	return res
}

func part2(data []string) error {
	pairs := []*pair{}
	for _, d := range data {
		sections := strings.Split(d, ",")
		p := &pair{}
		for _, s := range sections {
			p.section = append(p.section, newSection(s))
		}
		pairs = append(pairs, p)
	}

	ans := 0
	for _, p := range pairs {
		if p.isOverlap() {
			ans++
		}
	}

	fmt.Println(ans)
	return nil
}

func part1(data []string) error {
	pairs := []*pair{}
	for _, d := range data {
		sections := strings.Split(d, ",")
		p := &pair{}
		for _, s := range sections {
			p.section = append(p.section, newSection(s))
		}
		pairs = append(pairs, p)
	}

	ans := 0
	for _, p := range pairs {
		if p.isFullOverlap() {
			ans++
		}
	}

	fmt.Println(ans)
	return nil
}

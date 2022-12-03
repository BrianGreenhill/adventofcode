package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	backpackItems := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		backpackItems = append(backpackItems, scanner.Text())
	}

	fmt.Println("part1")
	if err := part1(backpackItems); err != nil {
		log.Fatal(err)
	}
	fmt.Println("part2")
	if err := part2(backpackItems); err != nil {
		log.Fatal(err)
	}
}

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type group struct {
	packs []*backpack
	badge string
}

type backpack struct {
	badge  string
	c1, c2 string
	dupes  []string
}

func (b *backpack) findDupes() {
	// find common elements between c1 and c2
	// set them on the dupes property
	for _, item := range strings.Split(b.c1, "") {
		if strings.Contains(b.c2, item) && !inSlice(item, b.dupes) {
			b.dupes = append(b.dupes, item)
		}
	}
}

func (b *backpack) dupePriority() int {
	total := 0
	for _, d := range b.dupes {
		total += toNum(d)
	}
	return total
}

func (g *group) setBadge() {
	p0 := g.packs[0].c1 + g.packs[0].c2
	p1 := g.packs[1].c1 + g.packs[1].c2
	p2 := g.packs[2].c1 + g.packs[2].c2
	lookup := strings.Split(p0, "")
	for _, item := range lookup {
		if inSlice(item, strings.Split(p1, "")) && inSlice(item, strings.Split(p2, "")) {
			g.badge = item
			return
		}
	}
}

func toNum(letter string) int {
	if strings.Contains(lower, letter) {
		return strings.Index(lower, letter) + 1
	}
	if strings.Contains(upper, letter) {
		return strings.Index(upper, letter) + 27
	}
	return -1
}

func inSlice(s string, sl []string) bool {
	for _, item := range sl {
		if item == s {
			return true
		}
	}
	return false
}

func part1(backpackItems []string) error {
	backpacks := []*backpack{}
	for _, item := range backpackItems {
		splitIndex := (len(item) / 2)
		backpacks = append(backpacks, &backpack{
			c1: item[:splitIndex],
			c2: item[splitIndex:],
		})
	}

	for _, b := range backpacks {
		b.findDupes()
	}

	total := 0
	for _, b := range backpacks {
		total += b.dupePriority()
	}
	fmt.Println(total)

	return nil
}

func part2(backpackItems []string) error {
	numGroups := len(backpackItems) / 3
	groups := []*group{}
	for i := 0; i < numGroups; i++ {
		items := backpackItems[i*3 : i*3+3]
		bs := []*backpack{}
		for _, item := range items {
			splitIndex := (len(item) / 2)
			b := &backpack{
				c1: item[:splitIndex],
				c2: item[splitIndex:],
			}
			b.findDupes()
			bs = append(bs, b)
		}
		g := &group{packs: bs}
		g.setBadge()
		groups = append(groups, g)
		bs = nil
	}
	total := 0
	for _, g := range groups {
		total += toNum(g.badge)
	}

	fmt.Println(total)
	return nil
}

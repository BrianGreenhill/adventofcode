package main

import (
	_ "embed"
	"fmt"
	"strings"
)

// Embed your input data
//
//go:embed input
var input string

type Coord = [2]int

// Offsets for "XMAS" patterns
var xmasOffsets = [][]Coord{
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}},       // Horizontal right
	{{0, 0}, {0, -1}, {0, -2}, {0, -3}},    // Horizontal left
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},       // Vertical down
	{{0, 0}, {-1, 0}, {-2, 0}, {-3, 0}},    // Vertical up
	{{0, 0}, {1, 1}, {2, 2}, {3, 3}},       // Diagonal down-right
	{{0, 0}, {-1, -1}, {-2, -2}, {-3, -3}}, // Diagonal up-left
	{{0, 0}, {1, -1}, {2, -2}, {3, -3}},    // Diagonal down-left
	{{0, 0}, {-1, 1}, {-2, 2}, {-3, 3}},    // Diagonal up-right
}

// Offsets for "X-MAS" patterns
var masOffsets = [][]Coord{
	{{-1, -1}, {0, 0}, {1, 1}}, // Diagonal top-left to bottom-right
	{{-1, 1}, {0, 0}, {1, -1}}, // Diagonal top-right to bottom-left
}

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	// Find "XMAS" patterns
	xmasCount := countPatterns(grid, xmasOffsets, "XMAS")
	fmt.Println("a", xmasCount)

	// Find "X-MAS" patterns
	masCount := countXMAS(grid, masOffsets)
	fmt.Println("b", masCount)
}

// Count patterns like "XMAS"
func countPatterns(grid [][]string, offsets [][]Coord, target string) int {
	count := 0
	for y, row := range grid {
		for x := range row {
			for _, offset := range offsets {
				if extractWord(grid, y, x, offset) == target {
					count++
				}
			}
		}
	}
	return count
}

// Count "X-MAS" patterns
func countXMAS(grid [][]string, offsets [][]Coord) int {
	count := 0
	for y, row := range grid {
		for x, char := range row {
			if char != "A" {
				continue
			}

			words := make([]string, len(offsets))
			valid := true
			for i, offset := range offsets {
				word := extractWord(grid, y, x, offset)
				if word != "MAS" && word != "SAM" {
					valid = false
					break
				}
				words[i] = word
			}
			if valid {
				count++
			}
		}
	}
	return count
}

// Extract a word from the grid based on offsets
func extractWord(grid [][]string, startY, startX int, offset []Coord) string {
	height := len(grid)
	width := len(grid[0])
	var word strings.Builder

	for _, coord := range offset {
		y, x := startY+coord[0], startX+coord[1]
		if y < 0 || y >= height || x < 0 || x >= width {
			return ""
		}
		word.WriteString(grid[y][x])
	}

	return word.String()
}

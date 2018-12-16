package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Claim struct {
	i      int
	x      int
	y      int
	width  int
	height int
}

type Grid [][]int

func (g Grid) String() string {
	fmt.Println("In grid String()")
	var sb strings.Builder
	for x := 0; x < len(g); x++ {
		for y := 0; y < len(g[x]); y++ {
			switch g[x][y] {
			case 0:
				sb.WriteRune('.')
			default:
				sb.WriteRune(rune(g[x][y]) + '0')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (c Claim) String() string {
	return fmt.Sprintf("Claim{%d, %d, %d, %d, %d}", c.i, c.x, c.y, c.width, c.height)
}

func byteToInt(b []byte) int {
	out, _ := strconv.Atoi(string(b))
	return out
}

func parseClaim(value string) Claim {
	pattern := regexp.MustCompile("#(\\d+)\\s@\\s(\\d+),(\\d+):\\s(\\d+)x(\\d+)")

	submatch := pattern.FindSubmatch([]byte(value))
	return Claim{
		byteToInt(submatch[1]),
		byteToInt(submatch[2]),
		byteToInt(submatch[3]),
		byteToInt(submatch[4]),
		byteToInt(submatch[5]),
	}
}

func updateGrid(grid [][]int, claim Claim) {
	for x := claim.x; x < claim.x+claim.width; x++ {
		for y := claim.y; y < claim.y+claim.height; y++ {
			grid[x][y]++
		}
	}
}

func gridCount(grid [][]int) int {
	count := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] > 1 {
				count++
			}
		}
	}
	return count
}

func initGrid() Grid {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}
	return grid
}

func First(input string) int {
	// Init grid
	grid := initGrid()
	// Parse claims from input and update grid
	values := strings.Split(input, "\n")

	for _, value := range values {
		if value == "" {
			continue
		}
		claim := parseClaim(value)
		updateGrid(grid, claim)
	}

	return gridCount(grid)
}

func doesOverlap(grid Grid, claim Claim) bool {
	for x := claim.x; x < claim.x+claim.width; x++ {
		for y := claim.y; y < claim.y+claim.height; y++ {
			if grid[x][y] > 1 {
				return true
			}
		}
	}
	return false
}

func Second(input string) int {
	// Init grid
	grid := initGrid()
	// Parse claims from input and update grid
	values := strings.Split(input, "\n")

	var claims []Claim

	for _, value := range values {
		if value == "" {
			continue
		}
		claim := parseClaim(value)
		claims = append(claims, claim)
		updateGrid(grid, claim)
	}

	// Find any claims that do not overlap
	for _, claim := range claims {
		if !doesOverlap(grid, claim) {
			return claim.i
		}
	}

	return -1
}

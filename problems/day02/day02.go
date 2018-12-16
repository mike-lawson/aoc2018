package day02

import (
	"fmt"
	"sort"
	"strings"

	"github.com/agnivade/levenshtein"
)

func First(input string) int {
	values := strings.Split(input, "\n")
	two := 0
	three := 0
	for _, val := range values {
		charmap := make(map[string]bool)
		hasTwo := false
		hasThree := false

		for i := 0; i < len(val); i++ {
			if hasTwo && hasThree {
				break
			}
			char := string(val[i])
			if _, ok := charmap[char]; ok {
				continue
			}
			charmap[char] = true
			count := strings.Count(val, char)
			switch {
			case count == 2 && !hasTwo:
				hasTwo = true
				two++
			case count == 3 && !hasThree:
				hasThree = true
				three++
			}
		}
	}
	return two * three
}

type ByDistance []Comparison
type Comparison struct {
	first    int
	second   int
	distance int
}

func (c ByDistance) Len() int           { return len(c) }
func (c ByDistance) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByDistance) Less(i, j int) bool { return c[i].distance < c[j].distance }

func Second(input string) string {
	values := strings.Split(input, "\n")
	var comparisons []Comparison
	for i := range values {
		for j := i + 1; j < len(values); j++ {
			if j == 0 {
				fmt.Println("wut")
			}
			comparison := Comparison{
				i, j, levenshtein.ComputeDistance(values[i], values[j]),
			}
			comparisons = append(comparisons, comparison)
		}
	}
	sort.Sort(ByDistance(comparisons))
	result := comparisons[0]
	if result.distance != 1 {
		return fmt.Sprintf("Cannot find match, closest distance is %d", result.distance)
	}
	var out string
	first := values[result.first]
	second := values[result.second]
	for i := range first {
		if first[i] == second[i] {
			out = out + string(first[i])
		}
	}
	return out
}

package day05

import (
	"math"
	"regexp"
	"strings"
)

func First(input string) int {
	// Strip out the last new line
	input = React(strings.Split(input, "\n")[0])
	return len(input)
}

func Second(input string) int32 {
	input = strings.Split(input, "\n")[0]
	results := make([]int32, 26)
	for c := 0; c < 26; c++ {
		charInput := regexp.MustCompile("(?i)"+string(c+'a')).ReplaceAllString(input, "")
		results[c] = int32(len(React(charInput)))
	}
	var min int32 = math.MaxInt32
	for _, result := range results {
		if result < min {
			min = result
		}
	}
	return min
}

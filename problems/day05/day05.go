package day05

import (
	"math"
	"strings"
)

func First(input string) int {
	// Strip out the last new line
	input = Reduce(strings.Split(input, "\n")[0])
	return len(input)
}

func Second(input string) int32 {
	input = strings.Split(input, "\n")[0]
	var charmap = make(map[rune]int32)
	for c := 'a'; c <= 'z'; c++ {
		charInput := strings.Replace(input, string(c), "", -1)
		charInput = strings.Replace(charInput, string(c-32), "", -1)
		charmap[c] = int32(len(Reduce(charInput)))
	}
	var min int32 = math.MaxInt32
	for _, v := range charmap {
		if int32(v) < min {
			min = int32(v)
		}
	}
	return min
}

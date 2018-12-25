package day05

import (
	"math"
	"regexp"
	"strings"
)

func First(input string) int {
	// Strip out the last new line
	input = Reduce(strings.Split(input, "\n")[0])
	return len(input)
}

func Second(input string) int32 {
	input = strings.Split(input, "\n")[0]
	out := make(chan int32)
	process := func(out chan int32, c rune) {
		charInput := regexp.MustCompile("(?i)"+string(c)).ReplaceAllString(input, "")
		out <- int32(len(Reduce(charInput)))
	}
	for c := 'a'; c <= 'z'; c++ {
		go process(out, c)
	}
	var min int32 = math.MaxInt32
	for i := 0; i < 26; i++ {
		result := <-out
		if result < min {
			min = result
		}
	}
	return min
}

package day01

import (
	"strconv"
	"strings"
)

func First(input string) int {
	values := strings.Split(input, "\n")
	sum := 0
	for _, value := range values {
		if value == "" {
			continue
		}
		intVal, err := strconv.Atoi(value)
		if err != nil {
			return -1
		}
		sum += intVal
	}
	return sum
}

func Second(input string) int {
	values := strings.Split(input, "\n")
	sum := 0
	state := make(map[int]bool)
	for {
		for _, val := range values {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			sum += intVal
			if _, ok := state[sum]; ok {
				return sum
			}
			state[sum] = true
		}
	}
}

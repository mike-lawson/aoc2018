package day05

import (
	"runtime"
)

func PolarMatch(i, j byte) bool {
	intI, intJ := int(i), int(j)
	return Abs(intI-intJ) == 32
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Map(fn func(string) string, inputs []string) []chan string {
	chs := make([]chan string, len(inputs))

	for i, input := range inputs {
		ch := make(chan string)
		go func(input string) {
			ch <- fn(input)
		}(input)
		chs[i] = ch
	}

	return chs
}

func Reduce(channels []chan string, fn func(string, string) string, acc string) chan string {
	ch := make(chan string)
	go func() {
		slices := make([]string, len(channels))

		for i, c := range channels {
			slices[i] = <-c
		}

		for _, slice := range slices {
			acc = fn(acc, slice)
		}

		ch <- acc
	}()
	return ch
}

func Shard(input string, shards int) []string {
	shardSize := len(input) / shards
	restSize := len(input) % shards
	slices := make([]string, shards)
	for i := 0; i < shards; i++ {
		slice := input[shardSize*i : shardSize*(i+1)]
		slices[i] = slice
	}
	if restSize > 0 {
		slices = append(slices, input[shardSize*shards:shardSize*shards+restSize])
	}
	return slices
}

func mapFn(in string) string {
	input := in

	for i := 0; i < len(input); {
		if i+1 < len(input) && PolarMatch(input[i], input[i+1]) {
			input = input[0:i] + input[i+2:]
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
	return input
}

func reduceFn(acc string, input string) string {
	for {
		if acc == "" {
			return input
		}
		if input == "" {
			return acc
		}
		if PolarMatch(acc[len(acc)-1], input[0]) {
			acc, input = acc[:len(acc)-1], input[1:]
		} else {
			return acc + input
		}
	}
}

func React(input string) string {
	shards := runtime.GOMAXPROCS(0)
	slices := Shard(input, shards)
	return <-Reduce(Map(mapFn, slices), reduceFn, "")
}

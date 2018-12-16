package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mike-lawson/aoc2018/exercise"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	day, err := strconv.ParseInt(os.Args[1], 10, 8)
	if err != nil {
		usage()
		return
	}

	if err := exercise.Run(int(day)); err != nil {
		fmt.Printf(err.Error())
		return
	}
}

func usage() {
	fmt.Printf("usage: aoc2018 <day>\n")
}

package exercise

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/mike-lawson/aoc2018/problems/day01"
	"github.com/mike-lawson/aoc2018/problems/day02"
	"github.com/mike-lawson/aoc2018/problems/day03"
	"github.com/mike-lawson/aoc2018/problems/day04"
	"github.com/mike-lawson/aoc2018/problems/day05"
)

var days = make(map[int]Exercise)

type Part func(string) interface{}
type Exercise []Part

func init() {
	days[1] = Exercise([]Part{
		func(input string) interface{} { return day01.First(input) },
		func(input string) interface{} { return day01.Second(input) },
	})
	days[2] = Exercise([]Part{
		func(input string) interface{} { return day02.First(input) },
		func(input string) interface{} { return day02.Second(input) },
	})
	days[3] = Exercise([]Part{
		func(input string) interface{} { return day03.First(input) },
		func(input string) interface{} { return day03.Second(input) },
	})
	days[4] = Exercise([]Part{
		func(input string) interface{} { return day04.First(input) },
		func(input string) interface{} { return day04.Second(input) },
	})
	days[5] = Exercise([]Part{
		func(input string) interface{} { return day05.First(input) },
		func(input string) interface{} { return day05.Second(input) },
	})
}

func Run(day int) error {
	file := fmt.Sprintf("%v/src/github.com/mike-lawson/aoc2018/exercise/inputs/%d.txt", os.Getenv("GOPATH"), day)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	exercise, ok := days[day]

	if !ok {
		return fmt.Errorf("no exercise found for day %d", day)
	}

	fmt.Printf("Day %d\n", day)

	for part, fn := range exercise {
		start := time.Now()
		fmt.Printf("Part %v: %v (took %s)\n", part+1, fn(string(bytes)), time.Since(start))
	}

	return nil
}

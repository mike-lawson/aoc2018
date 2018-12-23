package day04

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const TimeFormat = "2006-01-02 15:04"

type ByDateTime []LogInput
type LogInput struct {
	dateTime time.Time
	text     string
}

func (l LogInput) String() string {
	return fmt.Sprintf("%s - %s", l.dateTime.Format(TimeFormat), l.text)
}

func (b ByDateTime) Len() int           { return len(b) }
func (b ByDateTime) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByDateTime) Less(i, j int) bool { return b[i].dateTime.Before(b[j].dateTime) }

func parseDateTime(dateTime string) (time.Time, error) {
	return time.Parse(TimeFormat, dateTime)
}

func ParseLogInputs(input string) []LogInput {
	values := strings.Split(input, "\n")
	var logInputs []LogInput
	for _, value := range values {
		if value == "" {
			continue
		}
		dateTimeString, text := value[1:17], value[19:]
		dateTime, _ := parseDateTime(dateTimeString)
		logInputs = append(logInputs, LogInput{dateTime, text})
	}
	sort.Sort(ByDateTime(logInputs))
	return logInputs
}

func Mode(array []int) int {
	sort.Ints(array)
	mode := 0
	modeMax := 0
	for i := 0; i < len(array); i++ {
		currentCount := 1
		for j := i + 1; j < len(array); j++ {
			if array[j] == array[i] {
				currentCount++
			} else {
				break
			}
		}
		if currentCount > modeMax {
			mode = array[i]
			modeMax = currentCount
		}
	}
	return mode
}

func CountOccurences(array []int, match int) int {
	count := 0
	for _, v := range array {
		if v == match {
			count++
		}
	}
	return count
}

func TrackGuardSleep(logInputs []LogInput) map[string][]int {
	guardAsleepMinutes := make(map[string][]int)

	for i := 0; i < len(logInputs); {
		// Would probably use a regex here for real-world use cases
		guard := logInputs[i].text[7:11]
		i++
		asleepMinute := 0

	InnerLoop:
		for {
			if i == len(logInputs) {
				break
			}
			switch logInputs[i].text {
			case "falls asleep":
				asleepMinute = logInputs[i].dateTime.Minute()
				i++
			case "wakes up":
				for min := asleepMinute; min < logInputs[i].dateTime.Minute(); min++ {
					guardAsleepMinutes[guard] = append(guardAsleepMinutes[guard], min)
				}
				i++
			default:
				break InnerLoop
			}
		}
	}
	return guardAsleepMinutes
}

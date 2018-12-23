package day04

import (
	"strconv"
)

func First(input string) int {
	logInputs := ParseLogInputs(input)

	guardAsleepMinutes := TrackGuardSleep(logInputs)

	laziestGuard := ""
	laziestGuardMinutes := 0
	for k, v := range guardAsleepMinutes {
		if len(v) > laziestGuardMinutes {
			laziestGuard = k
			laziestGuardMinutes = len(v)
		}
	}

	minutes := guardAsleepMinutes[laziestGuard]
	commonMinute := Mode(minutes)

	guardNumber, _ := strconv.ParseInt(laziestGuard, 10, 32)

	return int(guardNumber) * commonMinute
}

func Second(input string) int {
	logInputs := ParseLogInputs(input)

	guardAsleepMinutes := TrackGuardSleep(logInputs)

	guardMatch := ""
	mode := 0
	modeMax := 0
	for guard, minutes := range guardAsleepMinutes {
		guardMode := Mode(minutes)
		guardModeMax := CountOccurences(minutes, guardMode)
		if guardModeMax > modeMax {
			mode = guardMode
			guardMatch = guard
			modeMax = guardModeMax
		}
	}

	guardNumber, _ := strconv.ParseInt(guardMatch, 10, 32)
	return int(guardNumber) * mode
}

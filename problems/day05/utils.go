package day05

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

func Min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func Reduce(input string) string {
	i := 1
	for {
		if i >= len(input) {
			break
		}
		if PolarMatch(input[i-1], input[i]) {
			input = input[0:i-1] + input[i+1:]
		} else if i+1 < len(input) && PolarMatch(input[i], input[i+1]) {
			min := Min(i+2, len(input))
			input = input[0:i] + input[min:]
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}
	return input
}

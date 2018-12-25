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

func Reduce(input string) string {
	for i := 0; i < len(input); {
		if i+1 < len(input) && PolarMatch(input[i], input[i+1]) {
			input = input[0:i] + input[i+2:]
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}
	return input
}

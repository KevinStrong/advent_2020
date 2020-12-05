package day1

func Solve(values []int) int {
	for _, first := range values {
		for _, second := range values {
			for _, third := range values {
				const magicNumber = 2020
				if first+second+third == magicNumber {
					return first * second * third
				}
			}
		}
	}

	return -1
}

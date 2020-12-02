package day_1

func Solve(values []int) int {
	for _, first := range values {
		for _, second := range values {
			for _, third := range values {
				if first+second+third == 2020 {
					return first * second * third
				}
			}
		}
	}
	return -1
}

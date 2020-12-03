package main

import (
	"advent_2020/day_1"
	"advent_2020/day_2"
	"advent_2020/day_3"
	"fmt"
)

func main() {
	fmt.Print(day_1.Solve(ReadLinesNumbers("day_1/input.txt")))
	fmt.Print(day_2.Solve(ReadLines("day_2/input.txt")))
	fmt.Print(day_3.Solve(ReadLines("day_3/input.txt")))

}

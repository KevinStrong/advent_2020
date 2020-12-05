package main

import (
	"advent_2020/day1"
	"advent_2020/day2"
	"advent_2020/day3"
	"advent_2020/day4"
	"fmt"
)

func main() {
	fmt.Print(day1.Solve(ReadLinesNumbers("day1/input.txt")))
	fmt.Print(day2.Solve(ReadLines("Day2/input.txt")))
	fmt.Print(day3.Solve(ReadLines("Day3/input.txt")))
	fmt.Print(day4.Solve(ReadLines("Day4/input.txt")))
}

package main

import (
	"advent_2020/day1"
	"advent_2020/day2"
	"advent_2020/day3"
	"advent_2020/day4"
	"advent_2020/day5"
	"advent_2020/day6"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Print(day1.Solve(ReadLinesNumbers("day1/input.txt")))
	fmt.Print(day2.Solve(ReadLines("Day2/input.txt")))
	fmt.Print(day3.Solve(ReadLines("Day3/input.txt")))
	fmt.Print(day4.Solve(ReadLines("Day4/input.txt")))
	fmt.Print(day5.Solve(ReadLines("Day5/input.txt")))
	fmt.Print(day6.Solve(ReadLines("Day6/input.txt")))
	fmt.Printf("Execution took %s", time.Since(start))
}

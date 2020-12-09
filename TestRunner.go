package main

import (
	"advent_2020/day1"
	"advent_2020/day2"
	"advent_2020/day3"
	"advent_2020/day4"
	"advent_2020/day5"
	"advent_2020/day6"
	"advent_2020/day7"
	"advent_2020/day8"
	"advent_2020/day9"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(day1.Solve(ReadLinesNumbers("day1/input.txt")))
	fmt.Println(day2.Solve(ReadLines("Day2/input.txt")))
	fmt.Println(day3.Solve(ReadLines("Day3/input.txt")))
	fmt.Println(day4.Solve(ReadLines("Day4/input.txt")))
	fmt.Println(day5.Solve(ReadLines("Day5/input.txt")))
	fmt.Println(day6.Solve(ReadLines("Day6/input.txt")))
	fmt.Println(day7.Solve(ReadLines("Day7/input.txt")))
	fmt.Println(day8.Solve(ReadLines("Day8/input.txt")))
	fmt.Println(day9.Solve(ReadLinesNumbers("Day9/input.txt")))
	fmt.Printf("Execution took %s", time.Since(start))
}

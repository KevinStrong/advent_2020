package main

import (
	"advent_2020/day14"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(day14.Solve(ReadLines("Day14/input.txt")))
	fmt.Printf("Execution took %s", time.Since(start))
}

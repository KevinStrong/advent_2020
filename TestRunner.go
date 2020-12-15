package main

import (
	"advent_2020/day15"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(day15.Solve([]int{18, 8, 0, 5, 4, 1, 20}))
	fmt.Printf("Execution took %s", time.Since(start))
}

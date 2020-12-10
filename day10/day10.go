package day10

import (
	"fmt"
	"sort"
)

var cache map[int]int = make(map[int]int)

func Solve(input []int) int {
	sort.Ints(input)
	fmt.Println(solvePartOne(input))
	return findAllChainsStartingHere(input, 0, cache)
}

func findAllChainsStartingHere(input []int, currentJoltage int, cache map[int]int) int {
	cachedTotalCombinations, cacheHit := cache[currentJoltage]
	if cacheHit {
		fmt.Printf("found cached value for %d:%d\n", currentJoltage, cachedTotalCombinations)
		return cachedTotalCombinations
	}
	totalCombinations := 0
	for _, adaptorJoltage := range input {
		joltageDifference := adaptorJoltage - currentJoltage
		if joltageDifference <= 3 && joltageDifference > 0 {
			totalCombinations += findAllChainsStartingHere(input, adaptorJoltage, cache)
		}
	}
	// If we are at the end then we have reached a valid adaptor chain
	if currentJoltage == input[len(input)-1] {
		totalCombinations++
	}
	cache[currentJoltage] = totalCombinations
	fmt.Printf("Comations for %d:%d\n", currentJoltage, totalCombinations)
	return totalCombinations
}

func solvePartOne(input []int) int {
	threeGap := getGapCount(input, 3)
	oneGap := getGapCount(input, 1)
	return oneGap * threeGap
}

// Assumes input is sorted
func getGapCount(input []int, gapSize int) int {
	gapCount := 0
	previousJoltage := 0
	for _, currentJoltage := range input {
		if currentJoltage-previousJoltage == gapSize {
			fmt.Printf("Gap between %d & %d\n", currentJoltage, previousJoltage)
			gapCount++
		}
		previousJoltage = currentJoltage
	}
	// There is 1 more three jolt gap count between the last adaptor and yoru device
	if gapSize == 3 {
		gapCount++
	}
	return gapCount
}

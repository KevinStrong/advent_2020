package day9

import "fmt"

func Solve(input []int) int {
	errorValue := solvePartOne(input)
	fmt.Println(errorValue)
	return solvePartTwo(input, errorValue)
}

func solvePartTwo(input []int, value int) int {
	startIndex := 0
	endIndex := 1
	sum := input[startIndex]
	for sum != value && endIndex != len(input) {
		if sum < value {
			sum += input[endIndex]
			endIndex++
		} else if sum > value {
			sum -= input[startIndex]
			startIndex++
		}
	}
	return sumOfSmallestAndLargestNumbers(input[startIndex:endIndex])
}

func sumOfSmallestAndLargestNumbers(ints []int) int {
	return getSmallestNumber(ints) + getLargestNumber(ints)
}

func getSmallestNumber(ints []int) int {
	var smallest int = ints[0]
	for _, value := range ints {
		if value < smallest {
			smallest = value
		}
	}
	return smallest
}

func getLargestNumber(ints []int) int {
	var largest int = ints[0]
	for _, value := range ints {
		if value > largest {
			largest = value
		}
	}
	return largest
}

func solvePartOne(input []int) int {
	var preamble int = 25
	for index, value := range input {
		// skip the preamble
		if index >= preamble {
			if !isValidEntry(input, index, preamble) {
				return value
			}
		}
	}
	return -1
}

func isValidEntry(input []int, entryIndex int, lookbackDistance int) bool {
	earlierValues := input[entryIndex-lookbackDistance : entryIndex]
	for _, firstSum := range earlierValues {
		for _, secondSum := range earlierValues {
			if firstSum+secondSum == input[entryIndex] && firstSum != secondSum {
				return true
			}
		}
	}
	return false
}

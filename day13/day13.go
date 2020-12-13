package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type BusArrivalRestriction struct {
	busInterval      int
	busArrivalOffset int
}

func Solve(input []string) int {
	fmt.Println(solvePartOne(input))
	return solvePartTwo(input)
}

func solvePartTwo(input []string) int {
	restrictions := buildBusArrivalRestrictions(input[1])
	return findValidTime(restrictions)
}

// Assume restrictions are sorted in descending busInterval order
func findValidTime(restrictions []BusArrivalRestriction) int {
	currentIntervalJump := restrictions[0].busInterval
	currentStartingTime := restrictions[0].busInterval
	previousNumberOfMatches := 1
	isCorrectInterval, numberOfMatches := isValidInterval(currentStartingTime, restrictions)
	for !isCorrectInterval {
		if numberOfMatches > previousNumberOfMatches {
			previousNumberOfMatches = numberOfMatches
			currentIntervalJump = getNewIntervalJump(numberOfMatches, restrictions)
		}
		currentStartingTime += currentIntervalJump
		isCorrectInterval, numberOfMatches = isValidInterval(currentStartingTime, restrictions)
	}
	return currentStartingTime
}

func getNewIntervalJump(matches int, restrictions []BusArrivalRestriction) int {
	factor := 1
	for _, restriction := range restrictions[:matches] {
		factor *= restriction.busInterval
	}
	return factor
}

func isValidInterval(interval int, restrictions []BusArrivalRestriction) (bool, int) {
	fmt.Printf("Checking Interval: %d\n", interval)
	for i, restriction := range restrictions {
		if (interval+restriction.busArrivalOffset)%restriction.busInterval != 0 {
			fmt.Printf("Failed on :%d\n", restriction.busInterval)
			return false, i
		}
	}
	fmt.Println("It worked!")
	return true, -1
}

func buildBusArrivalRestrictions(restrictionStringRaw string) []BusArrivalRestriction {
	offset := 0
	intervalStrings := strings.Split(restrictionStringRaw, ",")
	arrivalRestrictions := make([]BusArrivalRestriction, 0)
	for _, intervalString := range intervalStrings {
		if intervalString != "x" {
			interval, _ := strconv.Atoi(intervalString)
			arrivalRestrictions = append(arrivalRestrictions, BusArrivalRestriction{
				busInterval:      interval,
				busArrivalOffset: offset,
			})
		}
		offset++
	}
	return arrivalRestrictions
}

func solvePartOne(input []string) int {
	earliestPossibleDeparture, _ := strconv.Atoi(input[0])
	intervalStrings := strings.Split(input[1], ",")
	// Arbitrary large number instead of learning how to create max possible int value
	// Alternatively calculate and insert the first bus's value.
	var earliestDeparture = 10000000
	earliestBusID := 0
	for _, intervalString := range intervalStrings {
		thisBusDepartureTime := 0
		if intervalString != "x" {
			interval, _ := strconv.Atoi(intervalString)
			for thisBusDepartureTime < earliestPossibleDeparture {
				thisBusDepartureTime += interval
			}
			fmt.Printf("Interval %d leaves at %d\n", interval, thisBusDepartureTime)
			if thisBusDepartureTime < earliestDeparture {
				fmt.Println("and is the current closest departure")
				earliestDeparture = thisBusDepartureTime
				earliestBusID = interval
			}
		}
	}
	return (earliestDeparture - earliestPossibleDeparture) * earliestBusID
}

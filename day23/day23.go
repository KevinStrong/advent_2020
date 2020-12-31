package main

import (
	"fmt"
	"strconv"
)

type Cup struct {
	value int
	right *Cup
}

var numberOfCups = 1000000

//var numberOfCups = 9
var numberOfRound = 10000000

//var numberOfRound = 10
var valueToCupMap = make(map[int]*Cup, numberOfCups)

func main() {
	// Example order
	// startingOrder := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	startingOrder := []int{3, 6, 8, 1, 9, 5, 7, 4, 2}
	calculateCups(buildCups(startingOrder))
	fmt.Println("Product:", multiplyTwoCupsAfterCupOne())
}

func multiplyTwoCupsAfterCupOne() int {
	cupOne := findCupOne()
	firstCup := cupOne.right
	fmt.Println("First Cup: ", firstCup.value)
	secondCup := firstCup.right
	fmt.Println("Second Cup: ", secondCup.value)
	return firstCup.value * secondCup.value
}

func findCupOne() *Cup {
	cup := valueToCupMap[1]
	return cup
}

func buildCups(cupValues []int) *Cup {
	firstCup := buildCup(cupValues[0])
	valueToCupMap[firstCup.value] = firstCup
	previousCup := firstCup
	var nextCup *Cup
	for i := 1; i < numberOfCups; i++ {
		nextCup = buildCup(getCupValue(i, cupValues))
		valueToCupMap[nextCup.value] = nextCup
		(*previousCup).right = nextCup
		previousCup = nextCup
	}
	if nextCup != nil {
		nextCup.right = firstCup
	} else {
		firstCup.right = firstCup
	}
	return firstCup
}

func getCupValue(i int, values []int) int {
	if i < len(values) {
		return values[i]
	}
	return i + 1
}

func buildCup(i int) *Cup {
	return &Cup{value: i}
}

func _(cup *Cup) {
	var firstCup = cup
	fmt.Print(cup.value)
	cup = cup.right
	for firstCup != cup {
		fmt.Print(", ", cup.value)
		cup = cup.right
	}
	fmt.Println()
}

func calculateCups(currentCup *Cup) *Cup {
	moveCount := 1
	for moveCount <= numberOfRound {
		currentCup = performOneRound(currentCup)
		moveCount++
	}
	return currentCup
}

func performOneRound(currentCup *Cup) *Cup {
	// pick 3 cups clockwise of the current cup
	var firstCup = currentCup.right
	var secondCup = firstCup.right
	var thirdCup = secondCup.right
	var fourthCup = thirdCup.right

	destinationCup := findDestinationCup(currentCup, []*Cup{firstCup, secondCup, thirdCup})

	// move cups to be clockwise of the destination cup
	thirdCup.right = destinationCup.right
	destinationCup.right = firstCup
	currentCup.right = fourthCup

	// The cup clockwise of the current cup is the new current cup
	return currentCup.right
}

// destination cup is highest valued cup that is lower than current cup.  If there are now lower cups then pick the highest cup.  Exclude the 3 selected cups.
func findDestinationCup(current *Cup, ignoreList []*Cup) *Cup {
	destination := findHighestCupLowerThatCurrent(current, ignoreList)
	if destination == nil {
		destination = findHighestCup(ignoreList)
	}
	return destination
}

func findHighestCup(ignoreList []*Cup) *Cup {
	for currentValue := numberOfCups; currentValue > 0; currentValue-- {
		if !contains(valueToCupMap[currentValue], ignoreList) {
			return valueToCupMap[currentValue]
		}
	}
	panic("There should always be a highest cup: " + strconv.Itoa(len(ignoreList)))
}

func findHighestCupLowerThatCurrent(current *Cup, ignoreList []*Cup) *Cup {
	currentValue := current.value
	for currentValue := currentValue - 1; currentValue > 0; currentValue-- {
		if !contains(valueToCupMap[currentValue], ignoreList) {
			return valueToCupMap[currentValue]
		}
	}
	return nil
}

func contains(cup *Cup, list []*Cup) bool {
	for i := range list {
		if list[i].value == cup.value {
			return true
		}
	}
	return false
}

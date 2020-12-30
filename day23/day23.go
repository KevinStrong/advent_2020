package main

import "fmt"

type Cup struct {
	value int
	right *Cup
}

func main() {
	// My Input
	calculateCups(buildCups([]int{3, 6, 8, 1, 9, 5, 7, 4, 2}))
	// Sample Input
	//calculateCups(buildCups([]int{3,8,9,1,2,5,4,6,7}))
}

func buildCups(cupValues []int) *Cup {
	firstCup := buildCup(cupValues[0])
	previousCup := firstCup
	var nextCup *Cup
	for i := 1; i < len(cupValues); i++ {
		nextCup = buildCup(cupValues[i])
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

func buildCup(i int) *Cup {
	return &Cup{value: i}
}

func printCups(cup *Cup) {
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
	fmt.Println("Starting Order: ")
	printCups(currentCup)
	moveCount := 1
	for moveCount <= 100 {
		fmt.Println("Round: ", moveCount)
		currentCup = performOneRound(currentCup)
		printCups(currentCup)
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
		destination = findHighestCup(current, ignoreList)
	}
	return destination
}

func findHighestCup(current *Cup, ignoreList []*Cup) *Cup {
	trackingCup := current.right
	var destinationCup *Cup
	for trackingCup != current {
		if !contains(trackingCup, ignoreList) {
			if destinationCup == nil || destinationCup.value < trackingCup.value {
				destinationCup = trackingCup
			}
		}
		trackingCup = trackingCup.right
	}
	return destinationCup
}

func findHighestCupLowerThatCurrent(current *Cup, ignoreList []*Cup) *Cup {
	trackingCup := current.right
	var destinationCup *Cup
	for trackingCup != current {
		if !contains(trackingCup, ignoreList) {
			if trackingCup.value < current.value && (destinationCup == nil || trackingCup.value > destinationCup.value) {
				destinationCup = trackingCup
			}
		}
		trackingCup = trackingCup.right
	}
	return destinationCup
}

func contains(cup *Cup, list []*Cup) bool {
	for i := range list {
		if list[i].value == cup.value {
			return true
		}
	}
	return false
}

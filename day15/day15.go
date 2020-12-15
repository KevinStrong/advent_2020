package day15

func Solve(input []int) int {
	numbersInOrder := input
	// Make this 1 indexed so that zero means the number is absent from the list
	previousValue, indexOfLastOccurrence := setUpMap(numbersInOrder)
	numberToCountTo := 30000000
	for i := len(input); i < numberToCountTo; i++ {
		numbersInOrder = append(
			numbersInOrder,
			getNextNumber(previousValue, numbersInOrder, indexOfLastOccurrence),
		)
		indexOfLastOccurrence[previousValue] = i
		previousValue = numbersInOrder[len(numbersInOrder)-1]
	}
	return numbersInOrder[numberToCountTo-1]
}

func setUpMap(order []int) (int, map[int]int) {
	var numbersToLastOccurrence map[int]int = make(map[int]int)
	// Skip the last element; add it in later
	for i := 0; i < len(order)-1; i++ {
		// Make the map point to 1 based index so that returning a 0 means a map miss
		numbersToLastOccurrence[order[i]] = i + 1
	}
	return order[len(order)-1], numbersToLastOccurrence
}

func getNextNumber(previousValue int, numbers []int, indexOfLastOccurrence map[int]int) int {
	if indexOfLastOccurrence[previousValue] == 0 {
		return 0
	}
	// Both are 1 indexes so no adjustments are required to diff them
	return len(numbers) - indexOfLastOccurrence[previousValue]
}

package day6

import "fmt"

func Solve(answers []string) int {
	fmt.Printf("Part One Solution: %d", solvePartOne(answers))
	return solvePartTwo(answers)
}

func solvePartTwo(answers []string) int {
	yesCount := 0
	currentGroup := make([]string, 0)
	for _, answer := range answers {
		if answer != "" {
			fmt.Printf("Grouping this person: %s\n", answer)
			currentGroup = append(currentGroup, answer)
		} else {
			fmt.Print("End of group detected\n")
			yesCount += processGroup(currentGroup)
			currentGroup = currentGroup[:0] // Clear for next group
		}
	}
	return yesCount
}

func processGroup(group []string) int {
	encountered := map[string]int{}
	for _, singlePersonsAnswer := range group {
		fmt.Printf("Processing a single person %s\n", singlePersonsAnswer)
		for _, singleAnswer := range singlePersonsAnswer {
			encountered[string(singleAnswer)]++
		}
	}
	groupYesCount := 0
	for key := range encountered {
		if encountered[key] == len(group) {
			groupYesCount++
		}
	}
	fmt.Printf("This group has %d yes answers\n", groupYesCount)

	return groupYesCount
}

// solvePartOne unused to solve part 2
func solvePartOne(answers []string) int {
	totalCount := 0
	for _, answer := range answers {
		totalCount += len(removeDuplicates(answer))
	}
	return totalCount
}

func removeDuplicates(groupAnswers string) string {
	encountered := map[string]bool{}

	for _, c := range groupAnswers {
		encountered[string(c)] = true
	}

	var result = ""
	for key := range encountered {
		result += key
	}
	return result
}

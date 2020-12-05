package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type PotentialPassword struct {
	password          string
	requiredLetter    string
	requiredLetterMin int
	requiredLetterMax int
}

func createPotentialPassword(lineOfInput string) PotentialPassword {
	words := strings.Fields(lineOfInput)
	password := words[2]
	requiredLetter := string(words[1][0])
	requiredLetterMin, _ := strconv.Atoi(strings.Split(words[0], "-")[0])
	requiredLetterMax, _ := strconv.Atoi(strings.Split(words[0], "-")[1])
	fmt.Printf("Creating Password: %s; %s, %s-%s\n", password, requiredLetter, strconv.Itoa(requiredLetterMin), strconv.Itoa(requiredLetterMax))
	pp := PotentialPassword{
		password:          password,
		requiredLetter:    requiredLetter,
		requiredLetterMax: requiredLetterMax,
		requiredLetterMin: requiredLetterMin,
	}
	return pp
}

func (pp *PotentialPassword) isValidPartOne() bool {
	actualCount := strings.Count(pp.password, pp.requiredLetter)
	return pp.requiredLetterMin <= actualCount && actualCount <= pp.requiredLetterMax
}

func (pp *PotentialPassword) isValidPartTwo() bool {
	doesFirstIndexMatch := string(pp.password[pp.requiredLetterMin-1]) == pp.requiredLetter
	doesSecondIndexMatch := string(pp.password[pp.requiredLetterMax-1]) == pp.requiredLetter
	return (doesFirstIndexMatch && !doesSecondIndexMatch) || (!doesFirstIndexMatch && doesSecondIndexMatch)
}

func Solve(potentialPasswords []string) int {
	count := 0
	for _, potentialPassword := range potentialPasswords {
		password := createPotentialPassword(potentialPassword)
		if password.isValidPartTwo() {
			count++
		}
	}
	return count
}

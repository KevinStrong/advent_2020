package day8

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(input []string) int {
	for index := range input {
		accumulatorResult, doesExit := runProgram(input, index)
		// fmt.Printf("Flipped line %d, Acc:%d\n", index, accumulatorResult)
		if doesExit {
			fmt.Printf("Flipped line %d works!\n", index)
			return accumulatorResult
		}
	}
	return -1
}

func runProgram(input []string, commandToFlip int) (int, bool) {
	var accululator int = 0
	visitedLines := make(map[int]bool)
	currentLine := 0
	for !visitedLines[currentLine] && currentLine < len(input) {
		visitedLines[currentLine] = true
		instruction := strings.Split(input[currentLine], " ")[0]
		parameter, _ := strconv.Atoi(strings.Split(input[currentLine], " ")[1])
		// fmt.Printf("Instruction: %s, parameter: %d\n", instruction, parameter)
		switch instruction {
		case "nop":
			if currentLine != commandToFlip {
				currentLine++
			} else {
				currentLine += parameter
			}
		case "acc":
			currentLine++
			accululator += parameter
		case "jmp":
			if currentLine != commandToFlip {
				currentLine += parameter
			} else {
				currentLine++
			}
		}
	}

	return accululator, currentLine >= len(input)
}

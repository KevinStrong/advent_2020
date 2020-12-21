package main

import (
	"advent_2020/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	equations := input.ReadLines("day18/sample_input.txt")
	results := make([]int, len(equations))
	for i := range equations {
		results[i] = calculate(stripWhiteSpace(equations[i]))
		fmt.Println(results[i])
	}
	sum := sumResults(results)
	fmt.Println("Sum: ", sum)
}

func stripWhiteSpace(line string) string {
	return strings.ReplaceAll(line, " ", "")
}

func calculate(equation string) int {
	result, remaining := evaluateAnElement(equation)
	if remaining != "" {
		panic(remaining)
	}
	return result
}

// Convert the entire input to a single "element" by wrapping in parenthesis
func evaluateAnElement(equation string) (int, string) {
	lhs := 0
	remainingEquation := equation
	var nextChar = "try the first character please"
	for nextChar != "" {
		nextChar, remainingEquation = eatChar(remainingEquation)
		switch nextChar {
		case "+", "*":
			var rhs int
			rhs, remainingEquation = evaluateAnElement(remainingEquation)
			lhs = performOpp(lhs, nextChar, rhs)
		case "(":
			lhs, remainingEquation = evaluateAnElement(remainingEquation)
		case ")":
			return lhs, remainingEquation
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			return parseNumber(nextChar), remainingEquation
		default:
			panic(nextChar)
		}
	}
	panic("Wrap everything in a parenthesis")
}

func eatChar(equation string) (string, string) {
	if len(equation) == 0 {
		return "", ""
	}
	return string(equation[0]), equation[1:]
}

func parseNumber(char string) int {
	value, err := strconv.Atoi(char)
	if err != nil {
		panic(err)
	} else {
		return value
	}
}

func performOpp(lhs int, opp string, rhs int) int {
	switch opp {
	case "*":
		return lhs * rhs
	case "+":
		return lhs + rhs
	case "-":
		return lhs - rhs
	default:
		panic("Unknown opp: " + opp)
	}
}

func sumResults(results []int) int {
	sum := 0
	for i := range results {
		sum += results[i]
	}
	return sum
}

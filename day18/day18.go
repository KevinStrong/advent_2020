package main

import (
	"advent_2020/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	equations := input.ReadLines("day18/input.txt")
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
	result := 0
	element := make([]string, 0)
	nextChar, remainingEquation := eatChar(equation)
	for nextChar != "" {
		switch nextChar {
		case "+", "*":
			element = append(element, nextChar)
		case "(":
			result, remainingEquation = evaluateAnElement(remainingEquation)
			element = append(element, strconv.Itoa(result))
		case ")":
			return collapseElement(element), remainingEquation
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			element = append(element, nextChar)
		default:
			panic(nextChar)
		}
		nextChar, remainingEquation = eatChar(remainingEquation)
	}
	return result, remainingEquation
}
func collapseElement(element []string) int {
	onlyMultiplication := performAllAddition(element)
	return performAllMultiplication(onlyMultiplication)
}

func performAllMultiplication(element []string) int {
	product := 0
	for i := 0; i < len(element); i++ {
		switch element[i] {
		case "*":
			i++
			rhs, err := strconv.Atoi(element[i])
			if err != nil {
				panic(err)
			}
			product *= rhs
		default:
			lhs, err := strconv.Atoi(element[i])
			if err != nil {
				panic(err)
			}
			product = lhs
		}
	}
	return product
}

func performAllAddition(element []string) []string {
	result := make([]string, 0)
	for i := 0; i < len(element); i++ {
		switch element[i] {
		case "*":
			result = append(result, "*")
		case "+":
			previousValue := result[len(result)-1]
			result = result[:len(result)-1]
			lhs, err := strconv.Atoi(previousValue)
			if err != nil {
				panic(err)
			}
			i++
			rhs, err := strconv.Atoi(element[i])
			if err != nil {
				panic(err)
			}
			result = append(result, strconv.Itoa(lhs+rhs))
		default:
			result = append(result, element[i])
		}
	}
	return result
}

func eatChar(equation string) (string, string) {
	if len(equation) == 0 {
		return "", ""
	}
	return string(equation[0]), equation[1:]
}

func sumResults(results []int) int {
	sum := 0
	for i := range results {
		sum += results[i]
	}
	return sum
}

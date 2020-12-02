package main

import (
	"bufio"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLinesNumbers(location string) []int {
	stringValues := ReadLines(location)
	numberValues := make([]int, len(stringValues))
	for i, s := range stringValues {
		value, err := strconv.Atoi(s)
		check(err)
		numberValues[i] = value
	}
	return numberValues
}

func ReadLines(location string) []string {
	file, err := os.Open(location)
	check(err)
	defer check(file.Close())

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	check(scanner.Err())

	return lines
}

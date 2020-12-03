package day_3

import (
	"fmt"
	"strconv"
)

//Right 1, down 1. -> 84
//Right 3, down 1. -> 195
//Right 5, down 1. -> 70
//Right 7, down 1. -> 70
//Right 1, down 2. -> 47
// Product of all values -> 3772314000

var slopeRight int = 1
var slopeDown int = 2
var tree string = "#"

func Solve(hillside []string) int {
	currentX := 0
	treesHit := 0
	for hillRowIndex, hillRow := range hillside {
		if hillRowIndex%slopeDown == 0 {
			tempString := replaceAtIndex(hillRow, 'K', currentX)
			fmt.Printf("%s : %s\n", tempString, strconv.Itoa(currentX))
			if string(hillRow[currentX]) == tree {
				treesHit++
			}
			currentX += slopeRight
			currentX = currentX % len(hillRow)
		}
	}
	return treesHit
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

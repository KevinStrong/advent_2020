package main

import (
	"advent_2020/input"
	"fmt"
	"strconv"
)

func main() {
	// flipTiles(input.ReadLines("day24/sample_input.txt"))
	flipTiles(input.ReadLines("day24/input.txt"))
}

func flipTiles(lines []string) {
	flippedTiles := make(map[string]bool)
	for i := range lines {
		flipTile(flippedTiles, lines[i])
	}
	countFlippedTiles(flippedTiles)
}

func countFlippedTiles(flippedTiles map[string]bool) {
	numberOfFlippedTiles := 0
	for _, value := range flippedTiles {
		if value {
			numberOfFlippedTiles++
		}
	}
	fmt.Println("Number of tiles flipped: ", numberOfFlippedTiles)
}

func flipTile(tiles map[string]bool, directions string) {
	location := getTileLocation(directions)
	tiles[location] = !tiles[location]
}

func getTileLocation(directions string) string {
	parseLocation := new(int)
	east := new(int)
	northEast := new(int)
	for *parseLocation < len(directions) {
		moveOneTile(directions, parseLocation, east, northEast)
	}
	return strconv.Itoa(*east) + ":" + strconv.Itoa(*northEast)
}

func moveOneTile(directions string, parseLocation *int, east *int, northEast *int) {
	switch string(directions[*parseLocation]) {
	case "e", "w":
		moveEastOrWest(directions, parseLocation, east)
	case "n", "s":
		moveDiagonally(directions, parseLocation, northEast, east)
	default:
		panic("failed")
	}
}

func moveDiagonally(directions string, parseLocation *int, northEast *int, east *int) {
	switch directions[*parseLocation : *parseLocation+2] {
	case "ne":
		*northEast++
	case "nw":
		*northEast++
		*east--
	case "sw":
		*northEast--
	case "se":
		*northEast--
		*east++
	}
	*parseLocation += 2
}
func moveEastOrWest(directions string, parseLocation *int, east *int) {
	switch string(directions[*parseLocation]) {
	case "e":
		*east++
	case "w":
		*east--
	}
	*parseLocation++
}

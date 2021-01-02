package main

import (
	"advent_2020/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// tiles := buildStartingFloor(input.ReadLines("day24/sample_input.txt"))
	tiles := buildStartingFloor(input.ReadLines("day24/input.txt"))
	fmt.Println("Starting black tiles: ")
	countFlippedTiles(tiles)
	runCycles(tiles, 100)
}

func runCycles(tiles map[string]bool, cycleCount int) {
	for i := 0; i < cycleCount; i++ {
		tiles = runACycle(tiles)
		clearWhiteTiles(tiles)
		countFlippedTiles(tiles)
	}
}

func clearWhiteTiles(tiles map[string]bool) {
	for location, isBlack := range tiles {
		if !isBlack {
			delete(tiles, location)
		}
	}
}

func runACycle(tiles map[string]bool) map[string]bool {
	relevantTiles := getAllRelevantTiles(tiles)
	updatedTiles := make(map[string]bool, len(tiles))
	for location := range relevantTiles {
		updatedTiles[location] = shouldTileBeBlack(tiles, location)
	}
	return updatedTiles
}

func shouldTileBeBlack(tiles map[string]bool, location string) bool {
	adjacentBlackTiles := countAdjacentBlackTiles(tiles, location)
	if tiles[location] {
		return adjacentBlackTiles == 1 || adjacentBlackTiles == 2
	}
	return adjacentBlackTiles == 2
}

func countAdjacentBlackTiles(tiles map[string]bool, location string) int {
	blackNeighbors := 0
	neighbors := getNeighbors(location)
	for i := range neighbors {
		if tiles[neighbors[i]] {
			blackNeighbors++
		}
	}
	return blackNeighbors
}

func getAllRelevantTiles(tiles map[string]bool) map[string]bool {
	relevantTiles := make(map[string]bool, len(tiles))
	for location := range tiles {
		if tiles[location] {
			relevantTiles[location] = true
			addNeighbors(relevantTiles, location)
		}
	}
	return relevantTiles
}

func addNeighbors(tiles map[string]bool, location string) {
	for _, neighbor := range getNeighbors(location) {
		tiles[neighbor] = true
	}
}

func getNeighbors(location string) []string {
	neighbors := make([]string, 6)
	east, northEast := convertStringToCoordinates(location)
	for index, intPair := range getNeighborsOffset() {
		eastOffset := intPair[0]
		northEastOffset := intPair[1]
		neighbors[index] = convertCoordinatesToString(east+eastOffset, northEast+northEastOffset)
	}
	return neighbors
}

func getNeighborsOffset() [][]int {
	return [][]int{
		{1, 0},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{0, -1},
		{1, -1},
	}
}

func buildStartingFloor(lines []string) map[string]bool {
	flippedTiles := make(map[string]bool)
	for i := range lines {
		flipTile(flippedTiles, lines[i])
	}
	return flippedTiles
}

func countFlippedTiles(flippedTiles map[string]bool) {
	numberOfFlippedTiles := 0
	for _, value := range flippedTiles {
		if value {
			numberOfFlippedTiles++
		}
	}
	fmt.Println("Black Tiles: ", numberOfFlippedTiles)
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
	return convertCoordinatesToString(*east, *northEast)
}

func convertCoordinatesToString(east int, northEast int) string {
	return strconv.Itoa(east) + ":" + strconv.Itoa(northEast)
}

func convertStringToCoordinates(location string) (int, int) {
	split := strings.Split(location, ":")
	east, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	northEast, err2 := strconv.Atoi(split[1])
	if err2 != nil {
		panic(err2)
	}
	return east, northEast
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

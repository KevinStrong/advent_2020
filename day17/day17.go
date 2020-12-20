package main

import (
	"advent_2020/input"
	"fmt"
	"math"
)

type Location struct {
	x int
	y int
	z int
}

func main() {
	startingBoard := createBoard(input.ReadLines("day17/input.txt"))
	printBoard(startingBoard)
	endingBoard := solvePartOne(startingBoard)
	printBoard(endingBoard)
	countActiveCells(endingBoard)
}

func countActiveCells(board map[Location]bool) {
	activeCells := 0
	for location := range board {
		if board[location] {
			activeCells++
		}
	}
	fmt.Println("Result: ", activeCells)
}

func printBoard(board map[Location]bool) {
	minX, maxX := findRange(board, func(location Location) int { return location.x })
	minY, maxY := findRange(board, func(location Location) int { return location.y })
	minZ, maxZ := findRange(board, func(location Location) int { return location.z })
	printableBoard := make([][][]string, (maxZ+1)-minZ)
	for zIndex := range printableBoard {
		printableBoard[zIndex] = make([][]string, (maxY+1)-minY)
		for yIndex := range printableBoard[zIndex] {
			printableBoard[zIndex][yIndex] = make([]string, (maxX+1)-minX)
		}
	}
	for z := range printableBoard {
		for y := range printableBoard[z] {
			for x := range printableBoard[z][y] {
				printableBoard[z][y][x] = "."
			}
		}
	}

	for location := range board {
		charToUse := "."
		if board[location] {
			charToUse = "#"
		}
		printableBoard[location.z-minZ][location.y-minY][location.x-minX] = charToUse
	}

	printThreeDBoard(printableBoard)
}

func printThreeDBoard(printableBoard [][][]string) {
	for i := range printableBoard {
		fmt.Printf("Z level: %d \n\n", i)
		for i2 := range printableBoard[i] {
			for i3 := range printableBoard[i][i2] {
				fmt.Print(printableBoard[i][i2][i3])
			}
			fmt.Printf("\n")
		}
	}
}

func findRange(board map[Location]bool, f func(Location) int) (int, int) {
	min := math.MaxInt64
	max := math.MinInt64
	for location := range board {
		if f(location) < min {
			min = f(location)
		}
		if f(location) > max {
			max = f(location)
		}
	}
	return min, max
}

func solvePartOne(board map[Location]bool) map[Location]bool {
	var nextBoard = board
	for i := 0; i < 6; i++ {
		nextBoard = performOneCycle(nextBoard)
		printBoard(nextBoard)
	}
	return nextBoard
}

func performOneCycle(board map[Location]bool) map[Location]bool {
	nextCycle := make(map[Location]bool)
	for key := range findDisabledCellsToEnable(board) {
		nextCycle[key] = true
	}
	for key := range findEnabledCellsToEnable(board) {
		nextCycle[key] = true
	}
	return nextCycle
}

func findEnabledCellsToEnable(board map[Location]bool) map[Location]bool {
	nextCycle := make(map[Location]bool)
	for location := range board {
		if isLocationActive(board, location) {
			activeNeighbors := getActiveNeighbors(location, board)
			if activeNeighbors == 2 || activeNeighbors == 3 {
				nextCycle[location] = true
			}
		}
	}
	return nextCycle
}

func findDisabledCellsToEnable(board map[Location]bool) map[Location]bool {
	nextCycle := make(map[Location]bool)
	neighborCount := getAllCellsNeighborCount(board)
	for location := range neighborCount {
		if !isLocationActive(board, location) {
			if neighborCount[location] == 3 {
				nextCycle[location] = true
			}
		}
	}
	return nextCycle
}

func isLocationActive(board map[Location]bool, location Location) bool {
	return board[location]
}

func getAllCellsNeighborCount(board map[Location]bool) map[Location]int {
	neighbors := make(map[Location]int)
	for location := range board {
		for x := -1; x < 2; x++ {
			for y := -1; y < 2; y++ {
				for z := -1; z < 2; z++ {
					if !(z == 0 && y == 0 && x == 0) {
						neighbors[Location{x: x + location.x, y: y + location.y, z: z + location.z}]++
					}
				}
			}
		}
	}
	return neighbors
}

func getActiveNeighbors(location Location, board map[Location]bool) int {
	activeNeighbors := 0
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			for z := -1; z < 2; z++ {
				if !(z == 0 && y == 0 && x == 0) {
					if isLocationActive(
						board,
						Location{x: x + location.x, y: y + location.y, z: z + location.z},
					) {
						activeNeighbors++
					}
				}
			}
		}
	}
	return activeNeighbors
}

func createBoard(lines []string) map[Location]bool {
	board := make(map[Location]bool)
	for i := range lines {
		for i2 := range lines[i] {
			if string(lines[i][i2]) == "#" {
				board[Location{x: i2, y: i, z: 0}] = true
			}
		}
	}
	return board
}

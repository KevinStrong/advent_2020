package day11

import "fmt"

func Solve(input []string) int {
	var previous []string = make([]string, len(input))
	var current []string = input
	for !equal(current, previous) {
		printBoard(current)
		previous = current
		current = stepBoardForward(current)
	}
	return countFilledSeats(current)
}

func printBoard(current []string) {
	for _, row := range current {
		fmt.Println(row)
	}
	fmt.Println()
}

func countFilledSeats(current []string) int {
	filledSeats := 0
	for _, row := range current {
		for _, seatValue := range row {
			if string(seatValue) == "#" {
				filledSeats++
			}
		}
	}
	return filledSeats
}

func stepBoardForward(current []string) []string {
	var nextBoard = make([]string, len(current))
	for rowIndex, row := range current {
		nextRow := ""
		for columnIndex, seatValue := range row {
			switch string(seatValue) {
			case "L":
				if shouldFillSeat(rowIndex, columnIndex, current) {
					nextRow += "#"
				} else {
					nextRow += "L"
				}
			case "#":
				if shouldEmptySeat(rowIndex, columnIndex, current) {
					nextRow += "L"
				} else {
					nextRow += "#"
				}
			case ".":
				nextRow += "."
			}
		}
		nextBoard[rowIndex] = nextRow
	}
	return nextBoard
}

func shouldEmptySeat(rowIndex int, columnIndex int, current []string) bool {
	adjacentFilledSeats := countAdjacentSeats(rowIndex, columnIndex, current)
	return adjacentFilledSeats >= 4
}

func shouldFillSeat(rowIndex int, columnIndex int, current []string) bool {
	adjacentFilledSeats := countAdjacentSeats(rowIndex, columnIndex, current)
	return adjacentFilledSeats == 0
}

func countAdjacentSeats(rowToCount int, columnToStart int, current []string) int {
	adjacentFilledSeats := 0
	for rowOffset := -1; rowOffset < 2; rowOffset++ {
		for columnOffset := -1; columnOffset < 2; columnOffset++ {
			currentRow := rowToCount + rowOffset
			currentColumn := columnToStart + columnOffset
			if isOnBoard(currentColumn, currentRow, current) && checkAdjacentSeats(rowOffset, columnOffset) {
				if string(current[currentRow][currentColumn]) == "#" {
					adjacentFilledSeats++
				}
			}
		}
	}
	return adjacentFilledSeats
}

func checkAdjacentSeats(rowOffset int, columnOffset int) bool {
	return rowOffset != 0 || columnOffset != 0
}

func isOnBoard(columnIndex int, rowIndex int, board []string) bool {
	return columnIndex > -1 && rowIndex > -1 && columnIndex < len(board[0]) && rowIndex < len(board)
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

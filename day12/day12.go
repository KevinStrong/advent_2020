package day12

import (
	"fmt"
	"strconv"
)

func Solve(input []string) int {
	fmt.Printf("Day 12 Part 1: %d\n", solvePartOne(input))
	return solvePartTwo(input)
}

func solvePartTwo(input []string) int {
	shipY := 0
	shipX := 0
	wayPointY := 1
	wayPointX := 10
	for _, order := range input {
		fmt.Printf("Next Instruction %s\n", string(order[0]))
		switch string(order[0]) {
		case "F":
			numberOfTimesToJump := getMagnitude(order)
			for i := 0; i < numberOfTimesToJump; i++ {
				shipX += wayPointX
				shipY += wayPointY
			}
		case "L":
			numberOfTimesToRotate := getMagnitude(order) / 90
			for i := 0; i < numberOfTimesToRotate; i++ {
				wayPointX, wayPointY = rotate(wayPointX, wayPointY, counterclockwise)
			}
		case "R":
			numberOfTimesToRotate := getMagnitude(order) / 90
			for i := 0; i < numberOfTimesToRotate; i++ {
				wayPointX, wayPointY = rotate(wayPointX, wayPointY, clockwise)
			}
		case "N":
			wayPointY += getMagnitude(order)
		case "S":
			wayPointY -= getMagnitude(order)
		case "E":
			wayPointX += getMagnitude(order)
		case "W":
			wayPointX -= getMagnitude(order)
		}
		fmt.Printf("current ship location (%d,%d)\n", shipX, shipY)
		fmt.Printf("current wayp location (%d,%d)\n", wayPointX, wayPointY)
	}
	return abs(shipX) + abs(shipY)
}

type Direction int

const (
	counterclockwise Direction = iota
	clockwise
)

// This is implements the the linear transformation
// necessary to do a 90 degree clockwise or counterclockwise rotation
// For example, clockwise matrix below
// [  0 1 ]
// [ -1 0 ]
func rotate(wayPointX int, wayPointY int, direction Direction) (int, int) {
	switch direction {
	case clockwise:
		oldWayPointX := wayPointX
		wayPointX = wayPointY
		wayPointY = -oldWayPointX
	case counterclockwise:
		oldWayPointY := wayPointY
		wayPointY = wayPointX
		wayPointX = -oldWayPointY
	}
	return wayPointX, wayPointY
}

func solvePartOne(input []string) int {
	y := 0
	x := 0
	currentXDirection := 0 // 0 = east, 90 = north
	for _, order := range input {
		switch string(order[0]) {
		case "F":
			xChange, yChange := moveForward(currentXDirection, getMagnitude(order))
			x += xChange
			y += yChange
		case "L":
			currentXDirection += getMagnitude(order)
			currentXDirection %= 360
		case "R":
			currentXDirection -= getMagnitude(order)
			currentXDirection = add360UntilPositive(currentXDirection)
			currentXDirection %= 360
		case "N":
			y += getMagnitude(order)
		case "S":
			y -= getMagnitude(order)
		case "E":
			x += getMagnitude(order)
		case "W":
			x -= getMagnitude(order)
		}
	}
	return abs(x) + abs(y)
}

func moveForward(direction int, magnitude int) (x int, y int) {
	switch direction {
	case 0:
		return magnitude, 0
	case 90:
		return 0, magnitude
	case 180:
		return -magnitude, 0
	case 270:
		return 0, -magnitude
	}
	fmt.Printf("error!, direction: %d, magnitude: %d\n", direction, magnitude)
	return 0, 0
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func add360UntilPositive(direction int) int {
	for direction < 0 {
		direction += 360
	}
	return direction
}

func getMagnitude(order string) int {
	distance, _ := strconv.Atoi(order[1:])
	return distance
}

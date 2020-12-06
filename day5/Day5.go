package day5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solve(ids []string) int {
	seatIds := make([]int, 0)
	var highestID int = 0
	for _, id := range ids {
		row := convertToNumber(id[:7])
		column := convertToNumber(id[7:])
		fmt.Printf(
			"id: %s -- Row: %d -- Column: %d\n",
			id,
			row,
			column,
		)
		seatID := (row * 8) + column
		seatIds = append(seatIds, seatID)
		fmt.Printf("Value is: %d\n", seatID)
		if seatID > highestID {
			highestID = seatID
		}
	}
	return findGapInSeatIDs(seatIds)
}

func findGapInSeatIDs(seatIds []int) int {
	sort.Ints(seatIds)
	var mySeatID int = 0
	previousSeat := seatIds[0]
	for _, seatID := range seatIds {
		if seatID-previousSeat > 1 {
			mySeatID = seatID - 1
			fmt.Printf("Found a gap! Top of gap is %d and the gap is %d\n", seatID, seatID-previousSeat)
		}
		previousSeat = seatID
	}
	return mySeatID
}

func convertToNumber(id string) int {
	return convertToBinary(convertToBinaryString(id))
}

func convertToBinary(binaryString string) int {
	i, _ := strconv.ParseInt(binaryString, 2, 64)
	return int(i)
}

func convertToBinaryString(id string) string {
	replaceF := strings.ReplaceAll(id, "F", "0")
	replaceB := strings.ReplaceAll(replaceF, "B", "1")

	replaceR := strings.ReplaceAll(replaceB, "R", "1")
	replaceL := strings.ReplaceAll(replaceR, "L", "0")

	return replaceL
}

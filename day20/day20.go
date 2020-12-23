package main

import (
	"advent_2020/input"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	tiles := createTiles(input.ReadLines("day20/input.txt"))
	success, board := findValidBoard(makeEmptyBoard(calculateBoardSize(len(tiles))), tiles, 0, 0)
	if success {
		for i := range board {
			for i2 := range board[i] {
				fmt.Print(board[i][i2].id, ":", board[i][i2].sides[0], " ")
			}
			fmt.Println()
		}
		product := productOfCornerIds(board)
		fmt.Println(product)
	}
	fmt.Printf("Execution took %s", time.Since(start))
}

func productOfCornerIds(board [][]Tile) int {
	return board[0][0].id *
		board[0][len(board)-1].id *
		board[len(board)-1][0].id *
		board[len(board)-1][len(board)-1].id
}

func makeEmptyBoard(size int) [][]Tile {
	rows := make([][]Tile, size)
	for i := range rows {
		rows[i] = make([]Tile, size)
	}
	return rows
}

func calculateBoardSize(i int) int {
	sqrt := math.Sqrt(float64(i))
	roundSqrt := math.Round(sqrt)
	return int(roundSqrt)
}

func findValidBoard(currentBoard [][]Tile, tiles []Tile, x int, y int) (bool, [][]Tile) {
	// fmt.Println("Matching on:", x, ":", y)
	if y == (len(currentBoard)) {
		return true, currentBoard
	}
	for i := range tiles {
		doesTileMatch, updatedBoards := checkTileMatch(currentBoard, tiles[i], x, y)
		if doesTileMatch {
			updatedX, updatedY := moveToNextTile(x, y, currentBoard)
			for i2 := range updatedBoards {
				doesThisBoardWork, potentialAssembledPuzzle := findValidBoard(updatedBoards[i2], remove(tiles, i), updatedX, updatedY)
				if doesThisBoardWork {
					return true, potentialAssembledPuzzle
				}
			}
		}
	}
	return false, make([][]Tile, 0)
}

func checkTileMatch(board [][]Tile, tile Tile, x int, y int) (bool, [][][]Tile) {
	adjacentTiles := getAdjacentTiles(x, y, board)
	allPositionsForThisTile := generateAllTilePositions(tile)
	boardCopy := copyBoard(board)
	results := make([][][]Tile, 0)
	for i := range allPositionsForThisTile {
		if doesTileFit(adjacentTiles, allPositionsForThisTile[i]) {
			boardCopy[y][x] = allPositionsForThisTile[i]
			results = append(results, boardCopy)
			boardCopy = copyBoard(board)
		}
	}
	return len(results) > 0, results
}

func copyBoard(board [][]Tile) [][]Tile {
	boardCopy := make([][]Tile, len(board))
	for i := range board {
		boardCopy[i] = make([]Tile, len(board))
		for i2 := range board[i] {
			boardCopy[i][i2] = board[i][i2]
		}
	}
	return boardCopy
}

func doesTileFit(adjacentTiles []Tile, tile Tile) bool {
	northMatch := adjacentTiles[0].id == Empty.id || (tile.sides[0] == adjacentTiles[0].sides[2])
	eastMatch := adjacentTiles[1].id == Empty.id || (tile.sides[1] == adjacentTiles[1].sides[3])
	southMatch := adjacentTiles[2].id == Empty.id || (tile.sides[2] == adjacentTiles[2].sides[0])
	westMatch := adjacentTiles[3].id == Empty.id || (tile.sides[3] == adjacentTiles[3].sides[1])
	return northMatch && eastMatch && southMatch && westMatch
}

func generateAllTilePositions(tile Tile) []Tile {
	tiles := make([]Tile, 8)
	for i := 0; i < 4; i++ {
		tile = tile.rotate()
		tiles[i] = tile
	}
	tile = tile.flip()
	for i := 4; i < 8; i++ {
		tile = tile.rotate()
		tiles[i] = tile
	}
	return tiles
}

func getAdjacentTiles(x int, y int, board [][]Tile) []Tile {
	return []Tile{
		getNorthTile(x, y, board),
		getEastTile(x, y, board),
		getSouthTile(x, y, board),
		getWestTile(x, y, board),
	}
}

var Empty = Tile{
	fullTile: nil,
	sides:    nil,
	id:       0,
}

func getNorthTile(x int, y int, board [][]Tile) Tile {
	if y == 0 {
		return Empty
	}
	return board[y-1][x]
}

func getEastTile(x int, y int, board [][]Tile) Tile {
	if x == len(board)-1 {
		return Empty
	}
	return board[y][x+1]
}

func getSouthTile(x int, y int, board [][]Tile) Tile {
	if y == len(board)-1 {
		return Empty
	}
	return board[y+1][x]
}

func getWestTile(x int, y int, board [][]Tile) Tile {
	if x == 0 {
		return Empty
	}
	return board[y][x-1]
}

func moveToNextTile(x int, y int, board [][]Tile) (int, int) {
	x++
	if x == len(board) {
		x = 0
		y++
	}
	return x, y
}

func remove(slice []Tile, s int) []Tile {
	tiles := make([]Tile, len(slice))
	copy(tiles, slice)
	copy(tiles[s:], tiles[s+1:]) // Shift tiles[s+1:] left one index.
	tiles[len(tiles)-1] = Empty  // Erase last element (write zero value).
	tiles = tiles[:len(tiles)-1] // Truncate slice.
	return tiles
}

type Tile struct {
	fullTile [][]string
	sides    []int
	id       int
}

func (tile Tile) rotate() Tile {
	return Tile{
		fullTile: rotateTile(tile.fullTile),
		id:       tile.id,
		sides:    rotate(tile.sides),
	}
}

func rotateTile(tile [][]string) [][]string {
	rotatedTile := make([][]string, len(tile))
	for rowIndex := range tile {
		rotatedTile[rowIndex] = make([]string, len(tile))
		for columnIndex := range tile[rowIndex] {
			rotatedTile[rowIndex][columnIndex] = tile[len(tile)-columnIndex-1][rowIndex]
		}
	}
	return rotatedTile
}

func (tile Tile) flip() Tile {
	return Tile{
		fullTile: flipTile(tile.fullTile),
		id:       tile.id,
		sides:    flip(tile.sides),
	}
}

func flipTile(tile [][]string) [][]string {
	flippedTile := make([][]string, len(tile))
	for rowIndex := range tile {
		flippedTile[rowIndex] = make([]string, len(tile))
		for columnIndex := range tile[rowIndex] {
			flippedTile[rowIndex][columnIndex] = tile[rowIndex][len(tile)-columnIndex-1]
		}
	}
	return flippedTile
}

func flip(sides []int) []int {
	flipped := make([]int, len(sides))
	copy(flipped, sides)
	flipped[3], flipped[1] = flipped[1], flipped[3]
	flipped[0] = reverse(flipped[0])
	flipped[2] = reverse(flipped[2])
	return flipped
}

func rotate(nums []int) []int {
	// top becomes east directly
	// east becomes south with a shift
	// south becomes west directly
	// west becomes top with a shift
	r := len(nums) - 1
	nums = append(nums[r:], nums[:r]...)
	nums[0] = reverse(nums[0])
	nums[2] = reverse(nums[2])
	return nums
}

func reverse(i int) int {
	// convert to binary
	binary := strconv.FormatInt(int64(i), 2)
	reversedBinary := reverseString(binary)
	parseInt, err := strconv.ParseInt(reversedBinary, 2, 32)
	if err != nil {
		panic(err)
	}
	return int(parseInt)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func createTiles(lines []string) []Tile {
	tiles := make([]Tile, 0)
	var aTile = make([]string, 0)
	for i := range lines {
		if lines[i] != "" {
			aTile = append(aTile, lines[i])
		} else {
			tiles = append(tiles, makeTile(aTile))
			aTile = make([]string, 0)
		}
	}
	tiles = append(tiles, makeTile(aTile))
	return tiles
}

func makeTile(tile []string) Tile {
	return Tile{
		fullTile: convertToTwoDArray(tile[1:]), sides: buildSides(tile[1:]), id: parseID(tile[0]),
	}
}

func convertToTwoDArray(tile []string) [][]string {
	output := make([][]string, len(tile))
	for row := range tile {
		output[row] = make([]string, len(tile))
		for column := range tile[row] {
			output[row][column] = string(tile[row][column])
		}
	}
	return output
}

func parseID(s string) int {
	var idRegex = regexp.MustCompile(`^Tile (\d+):$`)
	idString := idRegex.FindStringSubmatch(s)
	id, err := strconv.Atoi(idString[1])
	if err != nil {
		panic(err)
	}
	return id
}

func buildSides(tile []string) []int {
	sides := make([]int, 4)
	sides[0] = buildTop(tile)
	sides[1] = buildEast(tile)
	sides[2] = buildSouth(tile)
	sides[3] = buildWest(tile)
	return sides
}

func buildEast(tile []string) int {
	side := ""
	for _, row := range tile {
		side += string(row[len(row)-1])
	}
	return convertToDecimal(side)
}

func buildWest(tile []string) int {
	side := ""
	for _, row := range tile {
		side += string(row[0])
	}
	return convertToDecimal(side)
}

func buildSouth(tile []string) int {
	return convertToDecimal(tile[len(tile)-1])
}

func buildTop(tile []string) int {
	return convertToDecimal(tile[0])
}

func convertToDecimal(sequence string) int {
	addZeros := strings.ReplaceAll(sequence, ".", "0")
	bothReplaced := strings.ReplaceAll(addZeros, "#", "1")
	parseInt, err := strconv.ParseInt(bothReplaced, 2, 32)
	if err != nil {
		panic(err)
	}
	return int(parseInt)
}

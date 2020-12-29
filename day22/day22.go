package main

import (
	"advent_2020/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	crabDeck, yourDeck := buildDecks(input.ReadLines("day22/input.txt"))
	playGame(crabDeck, yourDeck)
}

func playGame(deck []int, deck2 []int) {
	crabDeck, yourDeck := playRecursiveGame(deck, deck2, 0)
	_ = scoreDecks(crabDeck, yourDeck)

}

func playRecursiveGame(crabDeck []int, yourDeck []int, recurseLevel int) ([]int, []int) {
	cache := make(map[string]bool)
	for len(crabDeck) != 0 && len(yourDeck) != 0 {
		// fmt.Println("Crab Deck:", crabDeck)
		// fmt.Println("Your Deck", yourDeck)
		repeat := checkForRepeat(crabDeck, yourDeck, cache)
		if repeat {
			// fmt.Println("Repeat detected")
			// fmt.Println("Player 1 score: " + strconv.Itoa(scoreDeck(crabDeck)))
			// Return an fake value to signal that crab wins
			return make([]int, 1), make([]int, 0)
		}
		crabTop, yourTop := peekTop(crabDeck, yourDeck)
		crabDeck, yourDeck = removeTopCards(crabDeck, yourDeck)

		var crabWin bool
		if len(crabDeck) >= crabTop && len(yourDeck) >= yourTop {
			// fmt.Println("Recursing: ", recurseLevel + 1)
			recursiveCrabDeck, _ := playRecursiveGame(copyDeck(crabDeck[:crabTop]), copyDeck(yourDeck[:yourTop]), recurseLevel+1)
			crabWin = len(recursiveCrabDeck) > 0
		} else {
			crabWin = crabTop > yourTop
		}

		// Give cards to winner
		if crabWin {
			crabDeck = append(crabDeck, crabTop, yourTop)
		} else {
			yourDeck = append(yourDeck, yourTop, crabTop)
		}
	}
	return crabDeck, yourDeck
}

func checkForRepeat(deck []int, deck2 []int, cache map[string]bool) bool {
	hash := buildHashForDecks(deck, deck2)
	isFound := cache[hash]
	cache[hash] = true
	return isFound

}

func buildHashForDecks(deck []int, deck2 []int) string {
	var hash = ""
	for _, value := range deck {
		hash += strconv.Itoa(value) + ":"
	}
	hash += ";"
	for _, value := range deck2 {
		hash += strconv.Itoa(value) + ":"
	}
	return hash
}

func removeTopCards(deck []int, deck2 []int) ([]int, []int) {
	return deck[1:], deck2[1:]
}

func peekTop(deck []int, deck2 []int) (int, int) {
	return deck[0], deck2[0]
}

func copyDeck(original []int) []int {
	copyDeck := make([]int, len(original))
	copy(copyDeck, original)
	return copyDeck
}

func scoreDecks(crabDeck []int, yourDeck []int) int {
	var score int
	if len(crabDeck) > 0 {
		score = scoreDeck(crabDeck)
	} else {
		score = scoreDeck(yourDeck)
	}
	fmt.Println("Score: ", score)
	return score
}

func scoreDeck(deck []int) int {
	multiplier := 1
	deckLength := len(deck) - 1
	score := 0
	for i := range deck {
		score += deck[deckLength-i] * multiplier
		multiplier++
	}
	return score
}

func buildDecks(lines []string) ([]int, []int) {
	decks := make([][]int, 0)
	for i := range lines {
		if strings.HasPrefix(lines[i], "Player") {
			decks = append(decks, make([]int, 0))
		} else {
			card, err := strconv.Atoi(lines[i])
			if err != nil {
				panic(err)
			}
			decks[len(decks)-1] = append(decks[len(decks)-1], card)
		}
	}
	return decks[0], decks[1]
}

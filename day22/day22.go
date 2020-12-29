package main

import (
	"advent_2020/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	crabDeck, yourDeck := buildDecks(input.ReadLines("day22/input.txt"))
	i := 1
	for len(crabDeck) != 0 && len(yourDeck) != 0 {
		fmt.Println("Round: ", i)
		crabDeck, yourDeck = playOneRound(crabDeck, yourDeck)
		i++
	}
	fmt.Println(crabDeck)
	fmt.Println(yourDeck)
	if len(crabDeck) > 0 {
		fmt.Println("Crab Score: ", scoreDeck(crabDeck))
	} else {
		fmt.Println("Your Score: ", scoreDeck(yourDeck))
	}
	fmt.Println()
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

func playOneRound(deck1 []int, deck2 []int) ([]int, []int) {
	top1 := deck1[0]
	top2 := deck2[0]
	deck1 = deck1[1:]
	deck2 = deck2[1:]
	if top1 > top2 {
		deck1 = append(deck1, top1, top2)
	}
	if top2 > top1 {
		deck2 = append(deck2, top2, top1)
	}
	if top1 == top2 {
		panic(top1)
	}
	return deck1, deck2
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

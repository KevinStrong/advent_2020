package main

import (
	"fmt"
	"math"
)

func main() {
	// My values
	doorPublicKey := 12092626
	cardPublicKey := 4707356

	// Example values
	// doorPublicKey := 17807724
	// cardPublicKey := 5764801

	subject := 7
	fmt.Println("Searching for door loop size")
	doorLoopSize := findLoopSize(doorPublicKey, subject)
	fmt.Println("Searching for card loop size")
	cardLoopSize := findLoopSize(cardPublicKey, subject)

	encryptionKeyFromDoor := transform(cardPublicKey, doorLoopSize)
	encryptionKeyFromCard := transform(doorPublicKey, cardLoopSize)
	fmt.Println("Encryption Key from door: ", encryptionKeyFromDoor)
	fmt.Println("Encryption Key from card: ", encryptionKeyFromCard)
}

func findLoopSize(key int, subject int) int {
	value := 1
	for loopCount := 1; loopCount < math.MaxInt64; loopCount++ {
		value *= subject
		value %= 20201227
		if value == key {
			fmt.Println("Loop Size Found: ", loopCount)
			return loopCount
		}
	}
	panic("Failed to find loop size")
}

func transform(subject int, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subject
		value %= 20201227
	}
	return value
}

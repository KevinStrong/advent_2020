package main

import "fmt"

type Simple struct {
	value int
}

func main() {
	simple := Simple{value: 2}
	// Make a copy of the struct, place it in a new location.
	structCopy := simple
	// Point to the original struct
	simplePointer := &simple
	// The pointer to the original and the copy have different memory addresses
	fmt.Println(simplePointer == &structCopy)
	// The pointer to the original and the copy have the same value
	fmt.Println(*simplePointer == structCopy)
	// The pointer points to the same memory address as the original value
	fmt.Println(simplePointer == &simple)
}

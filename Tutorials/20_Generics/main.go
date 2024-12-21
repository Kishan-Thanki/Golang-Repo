package main

import "fmt"

// Stack[T comparable] is a generic stack implementation
// `T` is a type parameter that must be comparable (e.g., supports == or != operations).
type Stack[T comparable] struct {
	elements []T // Slice to store stack elements
}

// printSlice[T] is a generic function that prints elements of a slice.
// T is restricted to int, string, or bool for simplicity in this example.
func printSlice[T int | string | bool](items []T) {
	// Iterate through the slice and print each item
	for _, item := range items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func main() {
	// Example usage of printSlice with various types
	printSlice([]int{1, 2, 3})                              // Prints integers
	printSlice([]string{"Golang", "Java", "Python", "C++"}) // Prints strings
	printSlice([]bool{true, false, true, true})             // Prints booleans

	// Creating a generic stack of integers
	stack := Stack[int]{
		elements: []int{1, 2, 3}, // Initializing with some elements
	}

	// Print stack to verify its initialization
	fmt.Println(stack)
}

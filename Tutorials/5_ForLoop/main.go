package main

import "fmt"

func main() {
	// Example 1: While Loop (Go style)
	// In Go, a 'while' loop is implemented using a 'for' loop with a condition.
	i := 1       // Initialize the variable.
	for i <= 3 { // Loop while 'i' is less than or equal to 3.
		fmt.Println(i)
		i++ // Increment 'i'.
	}

	// Example 2: Infinite Loop
	// Use caution with infinite loops; they will run forever unless interrupted.
	// Uncomment to see this in action.
	/*
		for {
			fmt.Println(1) // Continuously prints '1'.
		}
	*/

	// Example 3: Classic For Loop
	// The classic for loop with initialization, condition, and increment.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Example 4: Break Loop
	// Use 'break' to exit a loop prematurely.
	for k := 0; k <= 3; k++ {
		if k == 2 {
			break // Exits the loop when k value is 2.
		}
		fmt.Println(k)
	}

	// Example 5: Continue Loop
	// Use 'continue' to skip the current iteration and proceed to the next.
	for l := 0; l <= 3; l++ {
		if l == 1 {
			continue // Skips the iteration when 'l' equals 1.
		}
		fmt.Println(l) // Prints all values except '1'.
	}

	// Example 6: Range Loop
	// Use 'range' to iterate over a collection or a slice.
	// In this example, we'll iterate over a slice of integers.
	numbers := []int{10, 20, 30} // A slice of integers.
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value) // Prints the index and value.
	}

	// Example 7: Iterating Over an Integer Range
	// Go does not support directly iterating over integers using 'range'.
	// To achieve this, we can create a slice of integers and use 'range'.
	for m := 1; m <= 10; m++ {
		fmt.Println(m) // Prints numbers from 1 to 10.
	}
}

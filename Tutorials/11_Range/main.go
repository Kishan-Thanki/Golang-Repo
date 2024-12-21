package main

import "fmt"

func main() {
	// Example 1: Iterating Over a Slice with a Classic For Loop
	nums := []int{6, 7, 8}
	for i := 0; i < len(nums); i++ { // Classic for loop with index-based access
		fmt.Println(nums[i])
	}

	// Example 2: Iterating Over a Slice with Range
	// Range is often used for concise iteration over slices
	sum := 0
	for _, num := range nums { // Use range to iterate over elements
		// The first value from range is the index (ignored here using _)
		// The second value is the element
		sum += num
	}
	fmt.Println(sum)

	// Example 3: Iterating Over a Map
	m := map[string]string{"fname": "John", "lname": "Doe"}
	for k, v := range m { // Use range to iterate over the map
		// k is the key, v is the value
		fmt.Println(k, v)
	}

	// Example 4: Iterating Over a String
	// Strings in Go are sequences of UTF-8 encoded bytes
	for i, c := range "Golang" { // Use range to iterate over a string
		// i is the index, c is the Unicode code point (rune) of each character
		fmt.Println(i, c, string(c)) // Print index, code point, and character
	}
}

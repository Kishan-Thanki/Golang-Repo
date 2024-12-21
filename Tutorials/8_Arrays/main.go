package main

import "fmt"

// numsArray demonstrates the usage of an integer array.
func numsArray() {
	// Declare an integer array of length 4. Uninitialized elements default to 0.
	var nums [4]int

	// Print the length of the array using the len() function.
	fmt.Println(len(nums))

	// Assign a value to the first index of the array.
	nums[0] = 1

	// Print the value at the first index.
	fmt.Println(nums[0])

	// Print the entire array. Uninitialized elements will show their default value (0).
	fmt.Println(nums) // NOTE: Uninitialized indexes' values will be 0 by DEFAULT.
}

// boolsArray demonstrates the usage of a boolean array.
func boolsArray() {
	// Declare a boolean array of length 4. Uninitialized elements default to false.
	var vals [4]bool

	// Assign a value (true) to the second index of the array.
	vals[1] = true

	// Print the entire array. Uninitialized elements will show their default value (false).
	fmt.Println(vals)
}

// namesArray demonstrates the usage of a string array.
func namesArray() {
	// Declare a string array of length 3. Uninitialized elements default to an empty string ("").
	var name [3]string

	// Assign a value to the third index of the array.
	name[2] = "golang"

	// Print the entire array. Uninitialized elements will show their default value (empty string).
	fmt.Println(name)
}

func main() {
	// Call the numsArray function to demonstrate integer array usage.
	numsArray()

	// Call the boolsArray function to demonstrate boolean array usage.
	boolsArray()

	// Call the namesArray function to demonstrate string array usage.
	namesArray()

	// Demonstrate declaration and initialization of an integer array with values.
	num := [3]int{1, 2, 3} // Declare and initialize an array of size 3 with specific values.
	fmt.Println(num)       // Print the entire array.

	// Demonstrate a two-dimensional integer array.
	twoDArray := [2][2]int{{3, 4}, {5, 6}} // A 2x2 matrix initialized with values.

	// Print the length of the 2D array (number of rows).
	fmt.Println(len(twoDArray))

	// Print the entire 2D array.
	fmt.Println(twoDArray)
}

package main

import "fmt"

func main() {
	// Example 1: Uninitialized Slice
	// A slice that is declared but not initialized is nil by default.
	var nums []int
	fmt.Println(nums == nil) // true: nums is nil
	fmt.Println(len(nums))   // 0: Length of an uninitialized slice is 0

	// Example 2: Slice Initialization Using make()
	// A slice can be created using make(). Syntax: make([]T, length, capacity)
	var nums2 = make([]int, 2, 5) // Slice with length 2 and capacity 5
	fmt.Println(nums2 == nil)     // false: nums2 is not nil
	fmt.Println(nums2)            // [0 0]: Initialized with default zero values
	fmt.Println(cap(nums2))       // 5: Capacity of the slice

	// Adding elements to the slice using append()
	nums2 = append(nums2, 1) // Append 1 to the slice
	nums2 = append(nums2, 2) // Append 2 to the slice
	nums2 = append(nums2, 3) // Append 3 to the slice
	nums2 = append(nums2, 4) // Append 4 to the slice
	fmt.Println(nums2)       // [0 0 1 2 3 4]: Slice with appended values
	fmt.Println(cap(nums2))  // 10: Capacity doubles when the slice grows beyond its current capacity
	fmt.Println(len(nums2))  // 6: Updated length after appending

	// Example 3: Slice Initialization Using Literal Syntax
	// Slices can be initialized directly using literals.
	nums3 := []int{10, 20, 30, 40, 50} // Slice with specified elements
	fmt.Println(nums3)                 // [10 20 30 40 50]
	fmt.Println(cap(nums3))            // 5: Capacity equals the number of elements
	fmt.Println(len(nums3))            // 5: Length equals the number of elements

	// Example 4: Slice Sub-slicing
	// You can create a new slice by slicing an existing slice.
	subSlice := nums3[1:4]     // Sub-slice from index 1 to 3 (excludes index 4)
	fmt.Println(subSlice)      // [20 30 40]
	fmt.Println(cap(subSlice)) // 4: Capacity is reduced but depends on the underlying array

	// Example 5: Copying Slices
	// Copy elements from one slice to another using the copy() function.
	src := []int{1, 2, 3}
	dest := make([]int, len(src)) // Create a destination slice of the same length as the source
	copy(dest, src)               // Copy elements from src to dest
	fmt.Println(dest)             // [1 2 3]: Destination slice now contains the elements of the source

	// Example 6: Iterating Over a Slice
	// Use a for loop or range to iterate over the slice.
	for i, value := range nums3 {
		fmt.Printf("Index: %d, Value: %d\n", i, value) // Prints index and value of each element
	}

	// Example 7: Modifying a Slice
	// Slices are reference types, so modifying one will affect others that share the same backing array.
	original := []int{5, 10, 15}
	modified := original  // Both slices share the same backing array
	modified[1] = 20      // Change the second element
	fmt.Println(original) // [5 20 15]: The change is reflected in the original slice

	// Example 8: Slices of Slices
	// You can create slices of slices (e.g., a 2D slice).
	twoD := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(twoD) // [[1 2 3] [4 5 6] [7 8 9]]

	// Example 9: Comparing Slices
	// Slices cannot be compared directly using == (except for nil checks).
	// To compare, iterate over elements or use reflect.DeepEqual().
	// Uncommenting the following will cause an error:
	// fmt.Println(nums3 == nums3) // Invalid

	// Example 10: Appending Multiple Elements
	// Append multiple elements at once.
	nums4 := []int{100, 200}
	nums4 = append(nums4, 300, 400, 500) // Appends multiple elements
	fmt.Println(nums4)                   // [100 200 300 400 500]

	// Example 11: Clearing a Slice
	// Reset a slice to an empty slice.
	nums4 = nums4[:0]  // Clears the slice
	fmt.Println(nums4) // []

	// Example 12: Nil vs Empty Slice
	// A nil slice has no backing array, while an empty slice does.
	var nilSlice []int             // Nil slice
	emptySlice := []int{}          // Empty slice
	fmt.Println(nilSlice)          // []
	fmt.Println(emptySlice)        // []
	fmt.Println(nilSlice == nil)   // true
	fmt.Println(emptySlice == nil) // false
}

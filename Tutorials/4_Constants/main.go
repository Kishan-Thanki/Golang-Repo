package main

import "fmt"

// Declare a constant outside the main function.
// Constants are values that cannot be changed once assigned.
const age = 30 // Declares an untyped constant.

func main() {
	// Example 1: String Constant
	const language string = "Golang" // Declares a string constant with an explicit type.
	fmt.Println(language)

	// Attempting to change a constant will result in a compilation error.
	// Example (uncomment to see the error): language = "Python"

	// Example 2: Integer Constant
	fmt.Println(age)

	// Example 3: Multiple Constants
	// Use the `const` block to declare multiple constants together.
	const (
		port = 5000        // Declares an integer constant for port.
		host = "localhost" // Declares a string constant for host.
	)

	fmt.Println(host, ":", port) // Combines and prints the host and port constants.

	// Example 4: Float Constant
	const pi = 3.14159 // Declares a float constant for the value of Pi.
	fmt.Println("Value of Pi:", pi)

	// Example 5: Boolean Constant
	const isGolangFun = true // Declares a boolean constant.
	fmt.Println("Is Golang fun?", isGolangFun)

	// Example 6: Untyped Constants
	// Constants without an explicit type (like 'age' and 'pi') are untyped and take on the type as needed.
	const untyped = 42       // Untyped constant.
	var number int = untyped // Untyped constant assigned to an int variable.
	fmt.Println("Untyped constant assigned to a variable:", number)
}

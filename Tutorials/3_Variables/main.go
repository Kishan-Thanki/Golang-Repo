package main

import "fmt"

func main() {
	// Example 1: String Variables
	// Declares a string variable with an explicit type and initializes it with "Golang".
	var lang1 string = "Golang"
	fmt.Println(lang1)

	// Declares a string variable with type inference and initializes it with "Python".
	var lang2 = "Python"
	fmt.Println(lang2)

	// Declares and initializes a string variable using shorthand syntax (type inferred).
	lang3 := "Java"
	fmt.Println(lang3)

	// Example 2: Integer Variables
	// Declares an integer variable with an explicit type and initializes it with 42.
	var number1 int = 42
	fmt.Println(number1)

	// Declares and initializes an integer variable using shorthand syntax.
	number2 := 100
	fmt.Println(number2)

	// Example 3: Float Variables
	// Declares a float variable with an explicit type and initializes it with 3.14159.
	var pi float64 = 3.14159
	fmt.Println(pi)

	// Declares and initializes a float variable using shorthand syntax.
	area := 45.67
	fmt.Println(area)

	// Example 4: Boolean Variables
	// Declares a boolean variable with an explicit type and initializes it with true.
	var isGoFun bool = true
	fmt.Println(isGoFun)

	// Declares and initializes a boolean variable using shorthand syntax.
	isPythonEasy := false
	fmt.Println(isPythonEasy)
}

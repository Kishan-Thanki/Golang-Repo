package main

import "fmt"

func main() {
	// Example 1: Basic if-else Ladder
	age := 18 // Declare and initialize the variable 'age'.

	if age >= 18 { // Checks if 'age' is 18 or greater.
		fmt.Println("Person is an Adult") // Executes if the condition is true.
	} else if age >= 12 { // Checks if 'age' is between 12 and 17 (inclusive).
		fmt.Println("Person is a Teenager") // Executes if the condition is true.
	} else { // Executes if all above conditions are false.
		fmt.Println("Person is a Kid")
	}

	// Example 2: Logical AND (&&) Condition
	var role = "Admin"        // Declare and initialize the variable 'role'.
	var hasPermissions = true // Declare and initialize the boolean 'hasPermissions'.

	if role == "Admin" && hasPermissions { // Checks if 'role' is "Admin" AND 'hasPermissions' is true.
		fmt.Println("Yes") // Executes if both conditions are true.
	}

	// Example 3: Short Variable Declaration in If Statement
	// You can declare a variable within an if-statement. The variable is scoped to the if-else block.
	if age := 15; age >= 18 { // Declares 'age' within the if-statement and checks if it's 18 or greater.
		fmt.Println("Person is an Adult")
	} else if age >= 59 { // Checks if the 'age' is 59 or greater.
		fmt.Println("Person is a Senior Citizen")
	} else { // Executes if all above conditions are false.
		fmt.Println("Person is under 18 and not a Senior Citizen")
	}
}

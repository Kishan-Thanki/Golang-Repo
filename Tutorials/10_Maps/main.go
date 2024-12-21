package main

import "fmt"

func main() {
	// Creating a map with string keys and integer values. The map is initialized with two key-value pairs.
	m1 := map[string]int{"price": 10, "phone": 1234567891}
	fmt.Println(m1) // Output: map[price:10 phone:1234567891]

	// Creating an empty map with string keys and string values using make().
	m := make(map[string]string)

	// Adding key-value pairs to the map.
	m["name"] = "golang"
	m["oops"] = "java"

	// Accessing values by key and printing them.
	fmt.Println(m["name"], m["oops"]) // Output: golang java

	// If the key does not exist, the value will be empty for strings, 0 for integers, and false for booleans.
	fmt.Println(m["phone"]) // Output: "" (empty string), because the key "phone" does not exist in the map.

	// Using the "comma ok" idiom to check if a key exists in the map.
	found, notFound := m["name"]

	// If the key exists, "notFound" will be false, and "found" will contain the value.
	if notFound {
		fmt.Println("Key not found")
	} else {
		// If the key is found, print the value.
		fmt.Printf("Key %s found", found) // Output: Key golang found
	}

	// Printing the number of key-value pairs in the map using len().
	fmt.Println(len(m)) // Output: 2, since "name" and "oops" are the only keys.

	// Deleting a key-value pair from the map.
	delete(m, "oops")

	// Printing the map after deletion. The key "oops" is removed.
	fmt.Println(m) // Output: map[name:golang]

	// Clearing the map using the clear function (from Go 1.21+). This removes all key-value pairs.
	clear(m)

	// Printing the map after clearing. It will be empty.
	fmt.Println(m) // Output: map[]
}

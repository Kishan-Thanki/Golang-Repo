package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, string, string) {
	return "Golang", "C++", "Java", "Python"
}

func processIt(fn func(a int) int) {
	fn(1)
}

func processIt2() func(a int) int {
	return func(a int) int {
		return 3
	}
}

func main() {
	result := add(3, 5)
	fmt.Println(result)

	fmt.Println(getLanguages())

	fn := func(a int) int {
		return 2
	}
	processIt(fn)

	fn1 := processIt2()
	fn1(8)
}

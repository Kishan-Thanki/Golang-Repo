package main

import "fmt"

// By value
func changeNumByValue(num int) {
	num = 5
	fmt.Println("In changeNumByValue", num)
}

func changeNumByReference(num *int) {
	*num = 5
	fmt.Println("In changeNumByReference", *num)
}

func main() {
	num := 1

	changeNumByValue(num)
	fmt.Println(num)

	changeNumByReference(&num)
	fmt.Println(num)
}

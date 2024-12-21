package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(1, 2, 34.56, 7, 8, "Hello", true)

	result := sum(1, 2, 3, 4)
	fmt.Println(result)

	nums := []int{3, 4, 5, 6}
	result = sum(nums...)
	fmt.Println(result)

}

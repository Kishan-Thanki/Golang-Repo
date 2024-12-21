package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple Switch
	i := 5

	switch i {
	case 1:
		fmt.Println("One")
		break

	case 2:
		fmt.Println("Two")
		break

	case 3:
		fmt.Println("Three")
		break

	case 4:
		fmt.Println("Four")
		break

	case 5:
		fmt.Println("Five")
		break

	default:
		fmt.Println("Other")
	}

	// Multiple-Condition Switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
		break

	default:
		fmt.Println("It's Working Day")
	}

	// Type Switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an Integer")

		case string:
			fmt.Println("It's a String")

		case bool:
			fmt.Println("It's a Boolean")

		default:
			fmt.Println("Other", t)
		}
	}
	whoAmI(55.9)
}

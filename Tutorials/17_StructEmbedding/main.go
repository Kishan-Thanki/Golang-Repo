package main

import (
	"fmt"
	"time"
)

type Customer struct {
	name    string
	contact string
}

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	Customer
}

func main() {

	customer1 := Customer{
		name:    "John",
		contact: "1234567891",
	}
	order1 := Order{
		id:       "101",
		amount:   79.99,
		status:   "Received",
		Customer: customer1,
	}

	fmt.Println("Order1 Struct", order1)
}

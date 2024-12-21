package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func newOrder(id string, amount float32, status string) *Order {
	createOrder := Order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &createOrder
}

func (O *Order) changeOrderStatus(status string) {
	O.status = status
}

func (O *Order) getOrderAmount() float32 {
	return O.amount
}

func main() {
	order1 := Order{
		id:        "101",
		amount:    50.59,
		status:    "Dispatch",
		createdAt: time.Now(),
	}
	fmt.Println("Order1 Struct", order1)
	order1.changeOrderStatus("Delivered")
	fmt.Println("Order1 Struct Status", order1.status)

	order2 := Order{
		id:        "102",
		amount:    79.99,
		status:    "Delivered",
		createdAt: time.Now(),
	}
	fmt.Println("Order2 Struct", order2)
	fmt.Println("Order2 Struct Amount", order2.getOrderAmount())

	order3 := newOrder("103", 100.99, "Received")
	fmt.Println("Order3 Struct", order3)

	language := struct {
		name   string
		isGood bool
	}{"Golang", true}
	fmt.Println(language)

}

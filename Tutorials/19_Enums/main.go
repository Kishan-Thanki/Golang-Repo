package main

import "fmt"

// Enumerated Types
type OrderStatus string

const (
	Delivered  OrderStatus = "Delivered"
	Dispatched             = "Dispatched"
	Confirmed              = "Confirmed"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changeOrderStatus(Delivered)
}

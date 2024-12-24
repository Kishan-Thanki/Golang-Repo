package main

import (
	"fmt"
	"sync"
	"time"
)

func validateOrder(orderID string) string {
	time.Sleep(2 * time.Second) // Simulate validation
	fmt.Println("Order", orderID, "validated.")
	return "Validated"
}

func shipOrder(orderID string, status string) {
	if status == "Validated" {
		time.Sleep(1 * time.Second) // Simulate shipping
		fmt.Println("Order", orderID, "shipped.")
	}
}

func processOrder(wg *sync.WaitGroup, orderID string) {
	defer wg.Done()

	status := validateOrder(orderID)
	shipOrder(orderID, status)
}

func main() {
	var wg sync.WaitGroup
	orderIDs := []string{"12345", "67890", "11223", "44556"}

	for _, orderID := range orderIDs {
		wg.Add(1)
		go processOrder(&wg, orderID)
	}

	wg.Wait()
}

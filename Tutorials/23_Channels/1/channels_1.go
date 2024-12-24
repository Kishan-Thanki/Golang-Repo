package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing number", num)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	numChan := make(chan int)
	go processNum(numChan)
	for i := 0; i <= 10; i++ {
		numChan <- rand.Intn(100)
	}
	time.Sleep(time.Second * 1)
	fmt.Println("Process completed")

	// messageChan := make(chan string)
	// messageChan <- "Ping" //Blocking
	// msg := <-messageChan
	// fmt.Println(msg)

}

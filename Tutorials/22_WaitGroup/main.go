package main

import (
	"fmt"
	"sync"
)

// task simulates a task identified by its ID.
// It uses WaitGroup to signal when the task is done.
func task(id int, w *sync.WaitGroup) {
	defer w.Done() // Marks this task as completed when it exits
	fmt.Println("Doing task:", id)
}

func main() {
	// WaitGroup is used to wait for all Goroutines to complete
	var wg sync.WaitGroup

	// Launch 10 tasks concurrently using Goroutines
	for i := 1; i <= 10; i++ {
		wg.Add(1) // Increment the counter for each Goroutine
		go task(i, &wg)
	}

	wg.Wait() // Block the main Goroutine until all tasks are done
	fmt.Println("All tasks completed")
}

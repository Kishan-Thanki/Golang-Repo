package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (P *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		P.mu.Unlock()
	}()

	P.mu.Lock()
	P.views++
}

func main() {
	var wg sync.WaitGroup

	myPost := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)
}

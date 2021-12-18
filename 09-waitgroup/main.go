package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	start := make(chan struct{})
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(count int) {
			defer wg.Done()
			<-start // wait for the start channel to be closed
			fmt.Println(count)
		}(i)
	}

	time.Sleep(3 * time.Second)

	// at this point, all goroutines are ready to go - we just need to
	// tell them to start by closing the start channel
	close(start)

	wg.Wait()
}

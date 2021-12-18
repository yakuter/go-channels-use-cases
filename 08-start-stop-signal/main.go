package main

import (
	"fmt"
	"time"
)

func main() {
	start := make(chan struct{})
	for i := 0; i < 10000; i++ {
		go func(k int) {
			<-start // wait for the start channel to be closed
			fmt.Println(k)
		}(i)
	}

	time.Sleep(3 * time.Second)

	close(start)
}

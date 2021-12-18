package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(5500 * time.Millisecond)
	ch <- "The secret of the world: 42"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}

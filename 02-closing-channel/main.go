package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer(ch)

	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

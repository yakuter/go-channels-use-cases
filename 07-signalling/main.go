package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

func main() {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{})

	// The sorting goroutine
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		// Notify sorting is done.
		fmt.Println("tamamlandi")
		done <- struct{}{}
	}()

	// do some other things ...

	<-done // waiting here for notification
	fmt.Println(values[0], values[len(values)-1])
}

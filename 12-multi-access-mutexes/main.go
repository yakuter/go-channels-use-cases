package main

import (
	"log"
	"math/rand"
	"time"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {
	log.Print("customer#", c, " enters the bar")
	seat := <-bar // need a seat to drink
	log.Print("++ customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("-- customer#", c, " frees seat#", seat)
	bar <- seat // free seat and leave the bar
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// the bar has 10 seats.
	bar24x7 := make(Bar, 3)
	// Place seats in an bar.
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		// None of the sends will block.
		bar24x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		go bar24x7.ServeCustomer(customerId)
	}

}

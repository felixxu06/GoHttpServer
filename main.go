package main

import (
	"fmt"
	"time"
)

func main() {
	quantity := 5
	requests := make(chan int, quantity)

	go func() {
		for i := 0; i < quantity*2; i++ {

			requests <- i
		}

		close(requests)
	}()

	limiter := time.Tick(time.Microsecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(2000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequest <- i
	}

	close(burstyRequest)
	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request 2 ", req, time.Now())
	}
}

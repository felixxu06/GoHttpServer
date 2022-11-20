package main

import (
	"fmt"
	"time"
)

func worker(signal chan bool) {
	fmt.Println("worker is working")
	time.Sleep(time.Second)
	fmt.Println("done")

	<-signal
}

func main() {
	signal := make(chan bool, 1)
	go worker(signal)
	<-signal
}

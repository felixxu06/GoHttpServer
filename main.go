package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "add message to c1"
	}()

	select {
	case msg := <-c1:
		fmt.Println("massage is ", msg)
	case <-time.After(time.Second):
		fmt.Println("timeout after two second")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "second message to c2"
	}()

	select {
	case msg := <-c2:
		fmt.Println("message", msg)
	case <-time.After(time.Second * 3):
		fmt.Println("time out after three second")
	}
}

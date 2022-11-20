package main

import "fmt"

func main() {
	stringChannel := make(chan string)

	go func() {
		stringChannel <- "signal"
	}()
	fmt.Println("before get message")
	msg := <-stringChannel
	fmt.Println("after get message")

	fmt.Println("message from channel", msg)
}

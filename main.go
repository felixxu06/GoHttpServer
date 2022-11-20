package main

import "fmt"

func ping(msg string, channel chan string) {
	fmt.Println("ping a message", msg)
	channel <- msg
}

func pong(ping chan string, pong chan string) {
	msg := <-ping
	fmt.Println("get a message", msg)
	pong <- msg
}

func main() {
	pingC := make(chan string, 1)
	pongC := make(chan string, 1)

	ping("a message", pingC)
	pong(pingC, pongC)

	fmt.Println("end")
}

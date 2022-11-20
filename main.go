package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	f("direct2")

	go func(msg string) {
		fmt.Println("message:", msg)
	}("going")
	fmt.Println("after going")
	time.Sleep(time.Second)

	fmt.Println("end")
}

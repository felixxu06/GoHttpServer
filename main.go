package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string, 2)

	myChannel <- "first"
	myChannel <- "second"

	fmt.Println("the message from chan is ", <-myChannel)
	fmt.Println("the message from chan is ", <-myChannel)

	fmt.Println("end")
}

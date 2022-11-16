package main

import "fmt"

func vals() (int, int) {
	return 3, 5
}

func main() {

	a, b := vals()
	fmt.Println("a is ", a, "b is ", b)

	_, c := vals()
	fmt.Println("c is ", c)
}

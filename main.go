package main

import "fmt"

func main() {
	i := 1
	for i < 3 {
		fmt.Println("i is ", i)
		i++
	}

	i = 1
	for {
		fmt.Println("i is", i)

		if i > 6 {
			break
		}
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println("j is ", j)
	}
}

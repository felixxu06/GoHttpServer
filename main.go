package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty array:", a)

	a[4] = 100
	fmt.Println("Set:", a)

	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [6]int{1, 2, 3, 4, 5, 8}
	fmt.Println("dcl:", b)

	for index, value := range b {
		fmt.Println("index is ", index, " value ", value)
	}

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d is ", twoD)
}

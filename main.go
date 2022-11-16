package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3, 4, 57, 88)

	numbers := []int{1, 2, 3, 4, 5, 9, 6, 8}
	sum(numbers...)
}

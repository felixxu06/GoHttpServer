package main

import "fmt"

func main() {
	nums := []int{2, 4, 7}

	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 4 {
			fmt.Println("current index is", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for index, value := range "this is a string" {
		fmt.Println("index is ", index, " value is ", value)
	}
}

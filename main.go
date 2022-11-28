package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"a", "c", "Anita", "Felix", "Nancy", "J&J"}

	sort.Strings(strings)

	fmt.Println("Strings:", strings)

	ints := []int{8, 10, 1}

	sort.Ints(ints)
	fmt.Println(ints)
}

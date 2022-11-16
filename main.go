package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["felix"] = 37
	m["anita"] = 35

	fmt.Println("map", m)
	v1 := m["felix"]
	fmt.Println("v1 is", v1)

	fmt.Println("len:", len(m))

	delete(m, "felix")
	fmt.Println("map after delete", m)

	_, prs := m["anita"]
	fmt.Println("prs:", prs)

	n := map[string]int{"xinxin": 7, "xincheng": 5}
	fmt.Println("map n", n)
}

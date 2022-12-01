package main

import (
	"encoding/json"
	"fmt"
)

type MyResponse struct {
	Page   int
	Fruits []string
}

type JsonResponse struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &MyResponse{
		Page: 1,
		Fruits: []string{
			"apple", "peach", "pear",
		},
	}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(res1B)

}

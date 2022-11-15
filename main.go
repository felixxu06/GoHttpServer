package main

import (
	"fmt"
	"math"
)

const s string = "const s value"

func main() {
	fmt.Println("const s is ", s)

	const n = 50000

	const d = 3e20 / n
	fmt.Println("const d is ", d)

	fmt.Println("int64(d)", int64(d))

	fmt.Println("math.Sin(n)", math.Sin(n))
}

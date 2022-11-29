package main

import "fmt"

type point struct {
	x, y int
}

func main() {

	p := point{x: 1, y: 2}
	fmt.Printf("strut1: %v \n", p)

	fmt.Printf("struct2: %+v\n", p)

	fmt.Printf("struct3: %#v\n", p)

	fmt.Printf("type: %T\n", p)

	fmt.Printf("bool: %t\n", true)

	fmt.Printf("int %d\n", 123)

	fmt.Printf("bin: %d\n", 14)

	fmt.Printf("char: %c\n", 33)

	fmt.Printf("hex: %x\n", 456)

	fmt.Printf("float1: %f\n", 78.9)
}

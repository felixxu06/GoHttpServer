package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 10
	fmt.Println("initial value of i is", i)

	zeroval(i)
	fmt.Println("after zeroval i is", i)

	zeroptr(&i)
	fmt.Println("after zeroptr i is", i)
}

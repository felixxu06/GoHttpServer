package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.2345", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("1212", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	p := fmt.Println

	u, _ := strconv.ParseUint("55665", 0, 64)
	p(u)

	k, _ := strconv.Atoi("135")
	p(k)

	_, e := strconv.Atoi("wat")
	p(e)
}

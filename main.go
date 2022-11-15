package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty is ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "e")
	s = append(s, "f", "g")
	fmt.Println("new s is ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("new l is ", l)

	fmt.Println("s[:5]", s[:5])

	fmt.Println("s[2:]", s[2:])
}

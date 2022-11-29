package main

import (
	"fmt"
	"regexp"
)

func main() {
	const expression = "p([a-z]+)ch"
	match, _ := regexp.MatchString(expression, "peach")
	fmt.Println("is match:", match)

	r, _ := regexp.Compile(expression)

	fmt.Println(r.MatchString("padadsch"))

	fmt.Println(r.FindString("asdads, peasdfasdch"))

	fmt.Println("idx:", r.FindStringIndex("aldfj, padfach"))

	fmt.Println(r.FindStringSubmatch("paldfkach adflkalsldk "))

}

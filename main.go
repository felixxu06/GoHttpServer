package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "foovariable")

	fmt.Println("Foo environment variable is", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println("======================================================")

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println("environment ", pair[0], " value is", pair[1])
	}
}

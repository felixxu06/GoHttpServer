package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numberPtr := flag.Int("numb", 42, "an int")

	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "svard", "a string var")

	flag.Parse()

	p := fmt.Println

	p("word:", *wordPtr)
	p("number:", *numberPtr)
	p("fork:", *forkPtr)
	p("svar:", svar)
	p("tail:", flag.Args())
}

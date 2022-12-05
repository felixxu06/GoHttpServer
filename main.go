package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "a int value level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	p := fmt.Println

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		p("subcommand foo")
		p(" enable", *fooEnable)
		p(" fooName", *fooName)
		p(" tail:", fooCmd.Args())

	case "bar":
		barCmd.Parse(os.Args[2:])
		p("subcommand bar")
		p(" level", *barLevel)

	default:
		p("not expected command ", os.Args[1])
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for index, value := range s {
		fmt.Printf("%#U starts at %d\n", value, index)
	}

	fmt.Println("\n Using DecodeRuneInstrings")

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Println("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'a' {
		fmt.Println("fount so sua")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func c(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	p := fmt.Printf

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("c:\\temp\\gotext.txt", d1, 0644)
	c(err)

	f, err := os.Create("c:\\temp\\gotext2.txt")
	c(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	c(err)
	p("wrote %v bytes \n", n2)

	n3, err := f.WriteString("write\n")
	c(err)
	p("wrote %d bytes for n3", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffed\n")
	c(err)
	p("wrote %d bytes for n4 \n", n4)

	w.Flush()
}

package main

import "os"

func main() {

	panic("a problem")

	_, error := os.Create("c:\temp\file.txt")
	if error != nil {
		panic(error)
	}
}

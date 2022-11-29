package main

import (
	"fmt"
	"os"
)

func createFile(path string) *os.File {
	fmt.Println("Create file:", path)
	f, error := os.Create(path)
	if error != nil {
		panic(error)
	}
	return f
}

func writeFile(file *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(file, "data")
}

func closeFile(file *os.File) {
	fmt.Println("closing")

	error := file.Close()

	if error != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", error)
	}
}

func main() {

	f := createFile("c:\\temp\\defer.txt")

	defer closeFile(f)

	writeFile(f)
}

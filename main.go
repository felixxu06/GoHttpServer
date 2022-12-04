package main

import "embed"

var fileString string
var fileByte []byte
var folder embed.FS

func main() {
	print(fileString)
	print(string(fileByte))

	content, _ := folder.ReadFile("folder/file1.hash")
	print(string(content))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}

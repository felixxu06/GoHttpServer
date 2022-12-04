package main

import (
	"os"
	"time"
)

func c(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const dir string = "subdir"
	err := os.Mkdir(dir, 0755)
	c(err)

	defer func() {
		time.Sleep(time.Second * 10)
		os.RemoveAll(dir)
	}()

	createEmptyFile := func(name string) {
		d := []byte("")
		c(os.WriteFile(name, d, 0644))
	}

	createEmptyFile(dir + "/file1")

	err = os.MkdirAll(dir+"/parent/child", 0755)
	c(err)

	createEmptyFile(dir + "/parent/file2")
	createEmptyFile(dir + "/parent/file3")
	createEmptyFile(dir + "/parent/child/file4")

}

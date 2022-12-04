package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func c(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp("", "sample")
	c(err)

	p := fmt.Println

	p("temp file name ", f.Name())

	defer func() {
		time.Sleep(time.Second * 10)
		os.Remove(f.Name())
	}()

	_, err = f.Write([]byte{1, 2, 3, 4})
	c(err)

	dname, err := os.MkdirTemp("", "sampledir")
	c(err)
	p("temp dir name:", dname)

	defer func() {
		time.Sleep(10 * time.Second)
		os.Remove(dname)
	}()

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	c(err)

}

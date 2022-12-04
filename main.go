package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	j := filepath.Join
	path := j("dir1", "\\dir2", "/dir4")

	p := fmt.Println

	p(path)

	p("dir(path)", filepath.Dir(path))

	p("base(path)", filepath.Base(path))

	p(filepath.IsAbs("dir/file"))

	p(filepath.IsAbs("c:\\temp\\abc.txt"))

	const fileName string = "config.json"

	ext := filepath.Ext(fileName)
	p(ext)

	p(strings.TrimSuffix(fileName, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	p(rel)
}

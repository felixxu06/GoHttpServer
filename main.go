package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)

	fmt.Println("Unix time is ", now.Unix())
	fmt.Println("Unix Milli is ", now.UnixMilli())
	fmt.Println("Unix Nano is", now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}

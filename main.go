package main

import (
	"fmt"
	"time"
)

func main() {
	const TimeFormat string = "2006-01-02 15:04:05"
	fmt.Println("1 time: ", time.Now().Format(TimeFormat))
	time1 := time.NewTimer(2 * time.Second)

	fmt.Println("2 time: ", time.Now().Format(TimeFormat))
	<-time1.C
	fmt.Println("Timer 1 fired")
	fmt.Println("3 time: ", time.Now().Format(TimeFormat))

	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("Timer 2 fired")
		fmt.Println("4 time: ", time.Now().Format(TimeFormat))
	}()

	stop2 := time2.Stop()

	if stop2 {
		fmt.Println("time 2 stop")
	}

	time.Sleep(time.Second * 4)
}

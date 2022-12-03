package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()

	p(now)

	then := time.Date(2019, 11, 11, 11, 50, 56, 111212, time.UTC)
	p(then)

	p(then.Year())

	p(then.Weekday())

	p(then.Before(now))

	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
}

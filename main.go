package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) increase(key string, n int) {
	for i := 0; i < n; i++ {
		c.mu.Lock()
		c.counters[key]++
		c.mu.Unlock()
	}
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup
	increaseKey := func(key string, n int) {
		for i := 0; i < n; i++ {
			c.increase(key, n)
		}
		wg.Done()
	}

	wg.Add(3)

	increaseKey("a", 1000)
	increaseKey("a", 1000)
	increaseKey("b", 1000)

	wg.Wait()
	fmt.Println("result", c.counters)
}

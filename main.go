package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func woker(taskId int) {
	fmt.Println("worker id is ", taskId, " started at", time.Now())
	waitTime := rand.Int63n(5) + 1
	time.Sleep(time.Second * time.Duration(waitTime))
	fmt.Println("worker id is ", taskId, " end at", time.Now())
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		j := i
		go func() {
			woker(j)
			wg.Done()
		}()

	}

	wg.Wait()

}

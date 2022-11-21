package main

import (
	"fmt"
	"time"
)

func work(workerId int, jobs <-chan int, result chan<- string) {
	for job := range jobs {
		fmt.Println("work id:", workerId, " start work on job", job)
		time.Sleep(time.Second)
		fmt.Println("work id:", workerId, " finish work on job", job)

		result <- fmt.Sprintln("work id:", workerId, " for job", job)
	}
}

func main() {
	jobNums := 5
	jobs := make(chan int, jobNums)
	result := make(chan string, jobNums)

	for i := 0; i < 3; i++ {
		go work(i, jobs, result)
	}

	for i := 0; i < jobNums; i++ {
		jobs <- i * 5
	}

	close(jobs)

	for i := 0; i < jobNums; i++ {
		fmt.Println("result is :", <-result)
	}

}

package main

import (
	"fmt"
	"sync"
	"time"
)

// Implement a worker pool in Go with the following requirements:

// 🔹 Input
// A list of files:
// files := []int{10, 20, 5, 15, 30, 25}
// (Each number = file size in MB)
// 🔹 Requirements
// Create N workers (e.g., 3 workers)
// Use:
// jobs channel → to send file sizes
// results channel → to collect results
// Each worker should:
// Pick a file
// Simulate processing using:
// time.Sleep(time.Duration(fileSize) * time.Millisecond)
// Return:
// fmt.Sprintf("File %d processed", fileSize)

func workerpool() {

	files := []int{10, 20, 5, 15, 30, 25}
	var wg sync.WaitGroup
	jobs := make(chan int)
	res := make(chan string)
	fmt.Println("spawning workers")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := range jobs {
				fmt.Println("worker", i, " picked a job", j)
				time.Sleep(time.Duration(j) * time.Millisecond)
				res <- fmt.Sprintf("worker %d doing work %d", i, j)
			}

		}(i)
	}

	go func() {
		for _, file := range files {
			fmt.Println("send job", file)
			jobs <- file
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		fmt.Println(r)
	}

}

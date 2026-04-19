package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	ID    int
	job   func(int) (float64, error)
	retry int
}

type res struct {
	ID    int
	value float64
	err   error
}

func queue() {

	n := 20
	taskqueue := make(chan task, 4)
	resultq := make(chan res, n)
	maxretry := 3
	var wg sync.WaitGroup
	var workers sync.WaitGroup

	for i := 0; i < 3; i++ {
		workers.Add(1)
		go func(i int) {
			defer workers.Done()
			for currentTask := range taskqueue {
				fmt.Println("worker: ", i, "doing task: ", currentTask.ID)
				val, err := currentTask.job(currentTask.ID)
				if err != nil {
					if currentTask.retry < maxretry {
						currentTask.retry++

						fmt.Println("reassigning the task again on the queue")

						go func(task task) {
							taskqueue <- task
						}(currentTask)

						continue

					}
					result := res{
						ID:    currentTask.ID,
						value: val,
						err:   err,
					}
					resultq <- result
					wg.Done()
					continue

				}
				result := res{
					ID:    currentTask.ID,
					value: val,
					err:   nil,
				}
				resultq <- result
				time.Sleep(100 * time.Millisecond)
				wg.Done()
			}
		}(i)
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		task := task{
			ID:  i,
			job: divide,
		}
		taskqueue <- task
		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
	close(taskqueue)
	workers.Wait()
	close(resultq)

	for res := range resultq {
		if res.err != nil {
			fmt.Println("result for ", res.ID, "failed:", res.err)
			continue
		}
		fmt.Println("result for ", res.ID, "is", res.value)
	}

}

package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	ID  int
	job func(int) int
}

func queue() {

	taskqueue := make(chan task, 4)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		go func(i int) {
			for task := range taskqueue {
				fmt.Println("worker: ", i, "doing task: ", task.ID)
				res := task.job(task.ID)
				fmt.Println("result of task: ", task.ID, res)
				time.Sleep(100 * time.Millisecond)
				wg.Done()
			}
		}(i)
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		task := task{
			ID:  i,
			job: factorial,
		}
		taskqueue <- task
		time.Sleep(100 * time.Millisecond)
	}

	close(taskqueue)

	wg.Wait()

}

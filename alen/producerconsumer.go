package main

import (
	"fmt"
	"sync"
	"time"
)

// Problem Statement

// You need to build a system where:

// There are multiple producers generating numbers
// There are multiple consumers processing those numbers
// All communication happens via channels
// You must ensure:
// No data loss
// No deadlocks
// Proper synchronization using sync.WaitGroup

func producerconsumer() {

	ch := make(chan string)
	var pwg sync.WaitGroup
	var cwg sync.WaitGroup
	// msgsent := 0
	// msgrec := 0

	var wg sync.WaitGroup
	// var mu sync.Mutex

	msgchan := make(chan string)
	wg.Add(1)
	// done := make(chan struct{})
	go func() {
		defer wg.Done()
		// for {
		// 	select {
		// 	case <-done:
		// 		return
		// 	default:
		// 		time.Sleep(1 * time.Second)
		// 		mu.Lock()
		// 		fmt.Println("📊 sent:", msgsent, "received:", msgrec)
		// 		mu.Unlock()
		// 	}
		// }
		msgsent := 0
		msgrec := 0

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case msg, ok := <-msgchan:
				if !ok {
					fmt.Println("sent :", msgsent, "received :", msgrec)
					return
				}
				if msg == "sent" {
					msgsent++
				} else if msg == "rec" {
					msgrec++
				}

			case <-ticker.C:
				fmt.Println("sent :", msgsent, "received :", msgrec)
			}
		}

	}()
	for i := 1; i <= 3; i++ {
		pwg.Add(1)
		go func(i int) {
			defer pwg.Done()
			for j := 1; j <= 5; j++ {
				work := fmt.Sprintf("from producer %d work : %d sent", i, j)
				ch <- work
				msgchan <- "sent"
				// mu.Lock()
				// msgsent++
				// mu.Unlock()
				time.Sleep(2 * time.Second)
			}
		}(i)
	}

	for i := 1; i < 3; i++ {
		cwg.Add(1)
		go func(i int) {
			defer cwg.Done()
			for work := range ch {
				worker := fmt.Sprintf("worker %d received : %s", i, work)
				msgchan <- "rec"
				// mu.Lock()
				// msgrec++
				// mu.Unlock()
				fmt.Println(worker)
				time.Sleep(2 * time.Second)
			}
		}(i)
	}
	go func() {
		pwg.Wait()
		close(ch)
	}()
	cwg.Wait()
	// close(done)
	close(msgchan)
	wg.Wait()
	fmt.Println("total msg recived")
}

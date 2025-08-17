package main

import (
	"fmt"
	"time"
)

// Golang Practice Set 1: Channels & Buffered Channels (No WaitGroups/Mutexes)

func channel() {

	ch := make(chan int)

	// 1. Send and receive a single message via channel
	go func() {
		ch <- 42 // Sending a value to the channel
	}()

	msg := <-ch // Receiving the value from the channel
	fmt.Println("Received message:", msg)

	// 2. Use buffered channel to send/receive without goroutine
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch2)
	}

	// 3. Send multiple messages through channel
	ch3 := make(chan int)
	done := make(chan bool)

	go func() {
		defer close(ch3)
		fmt.Println("sending the numbers")
		for i := 4; i < 10; i++ {
			ch3 <- i
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("all the signals are sent")
	}()

	go func() {
		for x := range ch3 {
			fmt.Println(x)
		}
		done <- true
	}()
	<-done
	fmt.Println("sending and receiving is done")

	// 4. Ping-pong between two goroutines
	ping := make(chan string)
	pong := make(chan string)

	//ping goroutine
	go func() {

		for i := 0; i < 5; i++ {
			// send the first message
			if i == 0 {
				ping <- "ping"
			} else {
				// receive from pong
				fmt.Println(<-pong)

				// send to pong
				ping <- "ping"
			}

		}

	}()

	//pong goroutine
	go func() {

		for i := 0; i < 5; i++ {

			// receive from ping
			fmt.Println(<-ping)
			// send to ping
			if i == 4 {
				break
			}
			pong <- "pong"
		}
		done <- true
	}()
	<-done

	// 5. Check if channel is empty using default
	ch4 := make(chan int)

	select {
	case msg := <-ch4:
		fmt.Println(msg)
	default:
		fmt.Println("empty")
	}

	// 6. Channel to sum numbers

	ch7 := make(chan int)
	res := make(chan int)

	go func() {
		sum := 0
		for n := range ch7 {
			sum += n
		}
		res <- sum
	}()

	for i := 0; i < 5; i++ {
		ch7 <- i
	}
	close(ch7)
	sum := <-res
	fmt.Println("result")
	fmt.Println(sum)

	// 7. Fan-out pattern
	work := make(chan int)

	for i := 0; i < 3; i++ {
		go func(workerId int) {
			for t := range work {
				fmt.Printf(" task : %d, worker : %d\n", t, workerId)
				time.Sleep(100 * time.Millisecond)
			}
			done <- true
		}(i)
	}
	for i := 0; i < 5; i++ {
		work <- i
	}
	close(work)
	<-done

	// 8. Buffered channel example
	buf := make(chan int, 5)

	for i := 0; i < 5; i++ {
		buf <- i
	}
	close(buf)

	for x := range buf {
		fmt.Println(x)
	}

	// 9. Channel direction (send-only/receive-only)

	//bidirectional channel
	ch8 := make(chan int)

	// send only channel
	var sendonly chan<- int

	// receive only channel
	var reconly <-chan int

	sendonly <- 2
	<-reconly
	ch8 <- 2
	<-ch8

	// 10. Signal completion with a channel
	// done above

	// 11. Infinite loop with select and channel

}

// 12. Use close to detect completion

// 13. Channel to square numbers

// 14. Producer-consumer

// 15. Select between two channels

// 16. Timer with channel

// 17. Ticker with channel

// 18. Buffered channel as queue

// 19. Closing a channel and using ok idiom

// 20. Simple pipeline

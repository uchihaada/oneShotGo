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
	msg1 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		msg1 <- 1
		time.Sleep(1 * time.Second)
		msg1 <- 2
		time.Sleep(1 * time.Second)
		msg1 <- 3
		done <- true
	}()

loop:
	for {
		select {
		case m := <-msg1:
			fmt.Println(m)
		case <-done:
			fmt.Println("done")
			break loop
		}
	}

	// 12. Use close to detect completion
	// close already used above

	// 13. Channel to square numbers
	num1 := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			num1 <- i
		}
		close(num1)
	}()

loop1:
	for {
		select {
		case num, ok := <-num1:
			if !ok {
				break loop1
			}
			fmt.Println(num)
		default:
			fmt.Println("done")

		}
	}

	// 14. Producer-consumer
	// already done

	// 15. Select between two channels
	messages := make(chan string)
	numbers := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		messages <- "Hello"
		time.Sleep(2 * time.Second)
		messages <- "World"
		close(messages)
	}()

	// Producer 2: sends numbers
	go func() {
		time.Sleep(500 * time.Millisecond)
		numbers <- 42
		time.Sleep(1500 * time.Millisecond)
		numbers <- 100
		time.Sleep(1 * time.Second)
		numbers <- 200
		close(numbers)
	}()

	activechannel := 2

	for activechannel > 0 {
		select {
		case msg, ok := <-messages:
			if !ok {
				messages = nil
				activechannel--
			} else {
				fmt.Println(msg)
			}
		case num, ok := <-numbers:
			if !ok {
				numbers = nil
			} else {
				fmt.Println(num)
			}

		}
	}

	// 16. Timer with channel

	timer := time.After(2 * time.Second)

loop3:
	for {
		select {
		case msg, ok := <-messages:
			if !ok {
				fmt.Println("All messages received!")
				break loop3
			}
			fmt.Printf("Received: %s\n", msg)

		case <-timer:
			fmt.Println("Timer expired! Taking too long...")
			// Reset timer for another 2 seconds
			timer = time.After(2 * time.Second)
		}
	}

	// 17. Ticker with channel

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Important: stop ticker to free resources

	tickCount := 0
loop4:
	for {
		select {
		case msg, ok := <-messages:
			if !ok {
				fmt.Println("All tasks completed!")
				break loop4
			}
			fmt.Printf("✅ %s\n", msg)

		case <-ticker.C:
			tickCount++
			fmt.Printf("Heartbeat #%d (system running...)\n", tickCount)
		}
	}

	// 18. Buffered channel as queue
	// done previously

	// 19. Closing a channel and using ok idiom
	// done previously

	// 20. Simple pipeline
	c1 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c1 <- i
		}
		close(c1)
	}()

	c2 := make(chan int)

	go func() {
		for n1 := range c1 {
			c2 <- n1 * n1
		}
		close(c2)
	}()

	for result := range c2 {
		fmt.Printf("%d\n", result)
	}
}

package main

import (
	"fmt"
	"time"
)

// Synchronus Programming
// Synchronous programming is a programming model where tasks are performed one after another, in a specific order. Each operation must finish before the next one starts. In Go (and many other languages), this means that code executes line by line, and if a function takes time to complete (like reading a file or making a network request), the program waits for it to finish before moving on

// Sequential Programming
// Sequential programming is a subset of synchronous programming where tasks are executed in a strict sequence, one after another. In Go, this means that each line of code is executed in the order it appears in the program. If one task takes a long time, the entire program waits for that task to``

// Concurrent Programming
// Concurrent programming is a programming model where multiple tasks can be executed independently and potentially at the same time.

// Parallel Programming
// Parallel programming is a type of concurrent programming where multiple tasks are executed simultaneously, often on multiple processors or cores.

// Concuqrency in Go
// In Go, concurrency is achieved through goroutines and channels. Goroutines are lightweight threads managed by the Go runtime, allowing you to run functions concurrently.
// this cannot return a value.

// func sayHello(ch chan string) {
// 	// fmt.Println("Hello from goroutine!")
// 	ch <- "Hello from goroutine!" // Send a message to the channel
// }

// func goroutine() {
// 	go sayHello()               // Start sayHello in a new goroutine
// 	time.Sleep(1 * time.Second) // Give goroutine time to finish
// 	fmt.Println("Main function finished")
// }

// channels are a way for goroutines to communicate with each other. They allow you to send and receive values between goroutines, enabling synchronization and data sharing.

// A channel is a typed conduit through which you can send and receive values with the channel operator, <-. The direction of the arrow indicates the direction of data flow: sending or receiving.

// func channel() {
// 	ch := make(chan string)
// 	sayHello(ch)
// 	msg := <-ch
// 	println(msg)
// }

// Channels can be buffered or unbuffered. Buffered channels have a capacity, allowing them to hold a certain number of values before blocking. Unbuffered channels block until both the sender and receiver are ready.

// Example of unbuffered channel
// func unbufferedChannelExample() {
// 	ch := make(chan string) // Unbuffered channel

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch <- "unbuffered: sent after 1 second"
// 	}()

// 	msg := <-ch // This will block until the goroutine sends a value
// 	fmt.Println(msg)
// }

// Example of buffered channel
// func bufferedChannelExample() {
// 	ch := make(chan string, 2) // Buffered channel with capacity 2

// 	ch <- "buffered: first message"
// 	ch <- "buffered: second message"
// 	// The above two sends do not block because the buffer has space

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch <- "buffered: third message" // This will block until there is space
// 	}()

// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// Call these functions in main() to see the difference
// func main() {
// 	unbufferedChannelExample()
// 	bufferedChannelExample()
// }
// Blocking and deadlocks happens when the go functions doesnt exist or dont have any value to return but is trying to read the returned value. The following code is blocked as there is no value returned.

// func block() {
// 	ch := make(chan int)
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 	}
// 	print(<-ch)
// }

// tokens in channel

// func worker(id int, tokens chan struct{}) {
// 	// Try to receive a token before proceeding
// 	<-tokens
// 	fmt.Printf("Worker %d is working\n", id)
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("Worker %d is done\n", id)
// 	// Return the token after work is done
// 	tokens <- struct{}{}
// }

// func token() {
// 	// Create a channel to hold 2 tokens (buffered channel)
// 	tokens := make(chan struct{}, 2)
// 	// Fill the channel with 2 tokens
// 	tokens <- struct{}{}
// 	tokens <- struct{}{}

// 	// Start 5 workers, but only 2 can work at a time
// 	for i := 1; i <= 5; i++ {
// 		go worker(i, tokens)
// 	}

// 	// Wait for all workers to finish
// 	time.Sleep(6 * time.Second)
// }

// This example uses a buffered channel as a "token bucket" to limit concurrency.
// The channel holds a fixed number of empty structs (tokens).
// Each worker must acquire a token from the channel before starting work.
// If no tokens are available, the worker blocks until one is returned.
// After finishing, the worker puts the token back, allowing another worker to proceed.
// This pattern is useful for controlling access to limited resources (e.g., limiting the number of concurrent goroutines).

// CLOSE A CHANNEL// Example: Closing a channel in Go

// func closeChannel() {
// 	ch := make(chan int)

// 	// Sender goroutine
// 	go func() {
// 		for i := 1; i <= 3; i++ {
// 			ch <- i
// 		}
// 		close(ch) // Close the channel after sending all values
// 	}()

// 	// Receiver loop
// 	for val := range ch {
// 		fmt.Println(val)
// 	}
// 	fmt.Println("Channel closed, done receiving.")
// }

// checking if a channel is closed
// V,OK := <-ch
// if channel is buffered it will be true if the channel is closed and there are no values left to read.

// use range to read values from a buffered channel or dynamically generated values from a unbuffered channel

// select

// func selectExample() {
// 	chInt := make(chan int)
// 	chStr := make(chan string)
// 	select {
// 	case i, ok := <-chInt:
// 		if ok {
// 			fmt.Printf("Received int: %d\n", i)
// 		} else {
// 			fmt.Println("chInt closed")
// 		}
// 	case s, ok := <-chStr:
// 		if ok {
// 			fmt.Printf("Received string: %s\n", s)
// 		} else {
// 			fmt.Println("chStr closed")
// 		}
// 	default:
// 		fmt.Println("No value received")
// 	}

// }

// ticker

// tickerExample demonstrates how to use a ticker to perform an action at regular intervals.
func tickerExample() {
	ticker := time.NewTicker(1 * time.Second) // Create a ticker that ticks every second
	defer ticker.Stop()                       // Always stop the ticker to release resources

	done := make(chan bool)

	// Goroutine to stop the ticker after 5 ticks
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	fmt.Println("Ticker started. It will tick every second:")
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		case <-done:
			fmt.Println("Ticker stopped.")
			return
		}
	}
}

// Call tickerExample() in main() to see it in action
// func main() {
// 	tickerExample()
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

// Golang Practice Set 2: WaitGroups & Mutexes

// Use a WaitGroup to wait for the goroutine to finish
var wg sync.WaitGroup

func waitgroupMutex() {

	// 1. Wait for one goroutine
	wg.Add(1)
	go func() {
		defer wg.Done() // Decrement the counter when the goroutine completes
		println("Goroutine is running")
	}()

	wg.Wait() // Wait for the goroutine to finish
	println("Goroutine finished execution")

	// 2. Wait for multiple goroutines
	wg.Add(2) // Add two to the WaitGroup counter

	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				time.Sleep(100 * time.Millisecond) // Simulate work
				println("Even number:", i)
			}
		}
	}()
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			if i%2 != 0 {
				time.Sleep(100 * time.Millisecond) // Simulate work
				println("Odd number:", i)
			}
		}
	}()

	wg.Wait()
	println("Both goroutines finished execution")

	// 3. Use mutex to protect shared variable
	var mu sync.Mutex
	counter := 0

	wg.Add(2) // Add two to the WaitGroup counter

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			mu.Lock() // Lock the mutex before accessing the shared variable
			counter++
			fmt.Println("Counter incremented by goroutine 1:", counter)
			mu.Unlock()                       // Unlock the mutex after accessing
			time.Sleep(50 * time.Millisecond) // Simulate work
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			mu.Lock()
			counter++
			fmt.Println("Counter incremented by goroutine 2:", counter)
			mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Wait()
	println("Final counter value:", counter)

	// lets create a struct to demonstrate mutex usage
	type SafeCounter struct {
		mu    sync.Mutex
		count int
	}

	sc := SafeCounter{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			sc.mu.Lock()
			sc.count++
			fmt.Println("SafeCounter incremented by goroutine 1:", sc.count)
			sc.mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			sc.mu.Lock()
			sc.count++
			fmt.Println("SafeCounter incremented by goroutine 2:", sc.count)
			sc.mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Wait()
	println("Final SafeCounter value:", sc.count)

	// 4. Safe concurrent map write with mutex
	type SafeMap struct {
		mu    sync.Mutex
		store map[string]int
	}

	sm := SafeMap{store: make(map[string]int)}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			sm.mu.Lock()
			sm.store[fmt.Sprintf("key%d", i)] = i
			fmt.Println("SafeMap updated by goroutine 1:", sm.store)
			sm.mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 5; i < 10; i++ {
			sm.mu.Lock()
			sm.store[fmt.Sprintf("key%d", i)] = i
			fmt.Println("SafeMap updated by goroutine 2:", sm.store)
			sm.mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Wait()
	println("Final SafeMap value:", sm.store)

	// 5. Read-write counter
	type ReadWriteCounter struct {
		mu    sync.RWMutex
		count int
	}
	rwc := ReadWriteCounter{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			rwc.mu.Lock() // Lock for writing
			rwc.count++
			fmt.Println("ReadWriteCounter incremented by goroutine 1:", rwc.count)
			rwc.mu.Unlock() // Unlock after writing
			time.Sleep(50 * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			rwc.mu.RLock() // Lock for reading
			fmt.Println("ReadWriteCounter read by goroutine 2:", rwc.count)
			rwc.mu.RUnlock() // Unlock after reading
			time.Sleep(50 * time.Millisecond)
		}
	}()
	wg.Wait()
	println("Final ReadWriteCounter value:", rwc.count)

	// 6. Multiple goroutines incrementing shared var
	type sharedVar struct {
		count int
		mu    sync.Mutex
	}

	scv := sharedVar{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			scv.mu.Lock()
			scv.count++
			fmt.Println("SharedVar incremented by goroutine 1:", scv.count)
			scv.mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			scv.mu.Lock()
			scv.count++
			fmt.Println("SharedVar incremented by goroutine 2:", scv.count)
			scv.mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			scv.mu.Lock()
			scv.count++
			fmt.Println("SharedVar incremented by goroutine 3:", scv.count)
			scv.mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()
	wg.Wait()
	println("Final SharedVar value:", scv.count)

	// 7. Print in sequence using WaitGroups
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		println("First goroutine is running")
		time.Sleep(100 * time.Millisecond) // Simulate work
		ch1 <- true
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ch1
		println("Second goroutine is running")
		time.Sleep(100 * time.Millisecond) // Simulate work
		ch2 <- true
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ch2
		println("Third goroutine is running")
		time.Sleep(100 * time.Millisecond) // Simulate work
	}()
	wg.Wait()
	println("All goroutines finished execution")

	// 8. Protecting counter using Mutex
	type Counter struct {
		mu    sync.Mutex
		count int
	}
	c := Counter{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			c.mu.Lock()
			c.count++
			fmt.Println("Counter incremented by goroutine 1:", c.count)
			c.mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			c.mu.Lock()
			c.count++
			fmt.Println("Counter incremented by goroutine 2:", c.count)
			c.mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()
	wg.Wait()
	println("Final Counter value:", c.count)

	// 9. Race condition demo (without mutex)
	raceCounter := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			raceCounter++
			fmt.Println("1:", raceCounter)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			raceCounter++
			fmt.Println("2:", raceCounter)
		}
	}()
	wg.Wait()
	println("Final raceCounter value (without mutex):", raceCounter)

	// 10. Fix race condition with mutex
	raceCounterMutex := 0
	var raceMu sync.Mutex
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			raceMu.Lock()
			raceCounterMutex++
			fmt.Println("1:", raceCounterMutex)
			raceMu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			raceMu.Lock()
			raceCounterMutex++
			fmt.Println("2:", raceCounterMutex)
			raceMu.Unlock()
		}
	}()
	wg.Wait()
	println("Final raceCounter value (with mutex):", raceCounterMutex)

	// 11. Nested goroutines with sync
	wg.Add(1)
	go func() {
		defer wg.Done()
		println("Outer goroutine is running")
		wg.Add(1)
		go func() {
			defer wg.Done()
			println("Inner goroutine is running")
			time.Sleep(100 * time.Millisecond) // Simulate work
		}()
		println("Inner goroutine finished execution")
	}()
	wg.Wait() // Wait for the outer goroutine to finish
	println("All nested goroutines finished execution")

	// 12. One-time print using sync.Once
	var once sync.Once
	wg.Add(1)
	go func() {
		defer wg.Done()
		once.Do(func() {
			println("This will only print once, even if called multiple times")
		})
	}()
	wg.Wait()
	println("One-time print completed")

	// show me that even if we call the once.Do multiple times, it will only execute once
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			once.Do(func() {
				println("This will still only print once, even in a loop")
			})
			time.Sleep(50 * time.Millisecond) // Simulate work
		}
	}()
	wg.Wait()
	println("One-time print in loop completed")

	// 13. WaitGroup in a loop
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			println("Goroutine", i, "is running")
			time.Sleep(100 * time.Millisecond) // Simulate work
		}(i)
	}
	wg.Wait()
	println("All goroutines in loop finished execution")

	// 14. Simultaneous readers/writers
	type RWCounter struct {
		mu    sync.RWMutex
		count int
	}
	rwc1 := RWCounter{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			rwc1.mu.Lock() // Lock for writing
			rwc1.count++
			fmt.Println("RWCounter incremented by goroutine 1:", rwc1.count)
			rwc1.mu.Unlock() // Unlock after writing
			time.Sleep(50 * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			rwc1.mu.RLock() // Lock for reading
			fmt.Println("RWCounter read by goroutine 2:", rwc1.count)
			rwc1.mu.RUnlock() // Unlock after reading
			time.Sleep(50 * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			rwc1.mu.RLock() // Lock for reading
			fmt.Println("RWCounter read by goroutine 3:", rwc1.count)
			rwc1.mu.RUnlock() // Unlock after reading
			time.Sleep(50 * time.Millisecond)
		}
	}()
	wg.Wait()
	println("Final RWCounter value:", rwc1.count)

	// 15. Parallel squares
	type Square struct {
		mu      sync.Mutex
		squares []int
	}
	sq := Square{squares: make([]int, 10)}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sq.mu.Lock() // Lock the mutex before accessing the shared variable
			sq.squares[i] = i * i
			fmt.Printf("Square of %d is %d\n", i, sq.squares[i])
			sq.mu.Unlock() // Unlock the mutex after accessing
		}(i)
	}
	wg.Wait()
	println("All squares calculated:", sq.squares)

	// 17. WaitGroup with function reuse
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			println("Reusable goroutine", i, "is running")
			time.Sleep(100 * time.Millisecond) // Simulate work
		}(i)
	}
	wg.Wait()
	println("All reusable goroutines finished execution")

	// 18. Sum of even and odd concurrently
	type Sum struct {
		mu   sync.Mutex
		even int
		odd  int
	}
	s := Sum{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				s.mu.Lock()
				s.even += i
				fmt.Println("Even sum updated by goroutine:", s.even)
				s.mu.Unlock()
			}
			time.Sleep(50 * time.Millisecond) // Simulate work
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i%2 != 0 {
				s.mu.Lock()
				s.odd += i
				fmt.Println("Odd sum updated by goroutine:", s.odd)
				s.mu.Unlock()
			}
			time.Sleep(50 * time.Millisecond) // Simulate work
		}
	}()
	wg.Wait()
	println("Final even sum:", s.even, "and odd sum:", s.odd)

	// 19. Chained WaitGroup execution
	wg.Add(1)
	go func() {
		defer wg.Done()
		println("Chained goroutine 1 is running")
		time.Sleep(100 * time.Millisecond) // Simulate work
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		println("Chained goroutine 2 is running")
		time.Sleep(100 * time.Millisecond) // Simulate work
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		println("Chained goroutine 3 is running")
		time.Sleep(100 * time.Millisecond) // Simulate work
	}()
	wg.Wait()

	// 20. Multiple counters combined
	wg.Add(2)

	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait()
	println("Final execution of waitgroupMutex completed")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrease counter by 1 when done
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

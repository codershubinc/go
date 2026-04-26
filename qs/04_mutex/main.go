// Level 4: Thread-Safe State (Mutex)
//
// Task: We have a telemetry agent trying to count system errors. 
// 1000 Goroutines are trying to increment the exact same map key simultaneously.
// If you run this code right now, it will crash or give the wrong number.
//
// 1. Add a `sync.Mutex` to the `SafeCounter` struct.
// 2. In the `Inc` method, Lock the mutex before modifying the map, and Unlock it after.
// 3. In the `Value` method, Lock the mutex before reading the map, and Unlock it after.

package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	// YOUR CODE HERE: Add a mutex (e.g., mu sync.Mutex)
	
	counts map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	// YOUR CODE HERE: Lock the mutex, increment the map, then Unlock
	
	c.counts[key]++
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	// YOUR CODE HERE: Lock the mutex, read the map, then Unlock
	
	return c.counts[key]
}

func main() {
	// We initialize our counter
	c := SafeCounter{counts: make(map[string]int)}
	
	// WaitGroup waits for a collection of Goroutines to finish
	var wg sync.WaitGroup

	// Launch 1000 concurrent workers
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("system_errors")
		}()
	}

	// Wait for all 1000 workers to finish
	wg.Wait()
	
	fmt.Println("Total Errors:", c.Value("system_errors"))
}

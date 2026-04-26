// Level 3: The Select Statement and Timeouts
//
// Task: You are querying a system service that is sometimes fast, but sometimes hangs.
// You need to implement a strict 2-second timeout to protect your orchestrator.
//
// 1. `slowQuery` simulates a system call that sends a result to a channel after a random delay.
// 2. In `main`, use a `select` block to wait for EITHER the result from `dataChan`
//    OR a 2-second timeout using `time.After()`.
// 3. Print the result if it succeeds, or "Timeout: Query took too long!" if it fails.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func slowQuery(c chan string) {
	// Simulates a query taking between 1 and 4 seconds
	delay := time.Duration(rand.Intn(3000)+1000) * time.Millisecond
	time.Sleep(delay)
	c <- "System Telemetry: [CPU: 45%, RAM: 2GB]"
}

func main() {
	// Create the channel
	dataChan := make(chan string)

	// Launch the query
	go slowQuery(dataChan)

	// YOUR CODE HERE:
	// Write a `select` statement that listens to `dataChan` AND `time.After(2 * time.Second)`.
	// Print the data if it arrives in time, otherwise print the timeout message.
	select {
	case res := <-dataChan:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("Slow query")

	}

}

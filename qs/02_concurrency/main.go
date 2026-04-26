// Level 2: Concurrent Slice Sum
//
// Task: We have a large slice of numbers. To speed up the calculation,
// we want to split the slice in half and have two separate Goroutines
// calculate the sum of their respective halves simultaneously.
//
// 1. Complete the `sumHalf` function to calculate the sum of a slice and
//    send the result into the provided channel.
// 2. In `main`, launch two Goroutines using `sumHalf`: one for the first
//    half of `numbers`, and one for the second half.
// 3. Receive the two results from the channel, add them together, and print the total.

package main

import "fmt"

func sumHalf(s []int, c chan int) {
	// YOUR CODE HERE: Calculate sum of 's' and send it to 'c'
	sum := 0
	for _, num := range s {
		sum += num
	}
	c <- sum
}

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80}

	// We create a channel of integers
	c := make(chan int)

	// YOUR CODE HERE:
	// 1. Launch Goroutine 1 for the first half of the slice
	// 2. Launch Goroutine 2 for the second half of the slice
	// 3. Read the two results from the channel and calculate the final total
	// 4. Print the final total
	mid := len(numbers) / 2
	go sumHalf(numbers[mid:], c)
	go sumHalf(numbers[:mid], c)
	sum1 := <-c
	sum2 := <-c
	fmt.Println("Total sum is", sum1+sum2)

}

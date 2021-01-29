package main

import "fmt"

func upTo(from, to int) chan int {
	var c chan int = make(chan int)
	go func() {
		for i := from; i < to; i++ {
			c <- i
		}
		close(c) // Using defer close() does not prevent the leak.
	}()
	return c
}

func main() {
	fmt.Println("Trying to leak channels...")
	// Try to leak channels.
	for {
		for a := range upTo(2, 5) {
			// Pretend that we are using 'a' to avoid the error: "a declared but
			// not used".
			if false {
				a++
			}

			// By breaking here, the channel returned by upTo will never be
			// closed.
			break
		}
	}

	// The channel created within upTo is never closed and the goroutine remains
	// suspended. This quickly causes gigabytes of memory to leak in a tight
	// loop.
}

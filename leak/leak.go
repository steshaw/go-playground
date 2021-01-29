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
	maybeLeakChannel := func() {
		for a := range upTo(2, 5) {
			a++ // Use 'a' because we can't use '_'.

			// By breaking here, the channel returned by upTo will never be
			// closed.
			break
		}
	}
	// Try to leak channels.
	for {
		maybeLeakChannel()
	}
}

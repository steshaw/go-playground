package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	msgs := make(chan string, 5)

	// Preload bufferred channel.
	for i := 1; i <= 5; i++ {
		msgs <- fmt.Sprintf("outer-%d", i)
	}

	go func() {
		defer close(msgs)
		for i := 0; i < 10; i++ {
			num := rand.Intn(10)
			msg := fmt.Sprintf("inner-%d", num)
			fmt.Printf("sending '%s'\n", msg)
			msgs <- msg
			fmt.Printf("sent '%s'\n", msg)
			sleep := time.Duration(rand.Intn(1000)) * time.Millisecond
			fmt.Println("sleeping for", sleep)
			time.Sleep(sleep)
			fmt.Println("slept")
		}
	}()

	for msg := range msgs {
		fmt.Println("received", msg)
	}
}

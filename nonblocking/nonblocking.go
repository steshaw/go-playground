package main

import (
	"fmt"
	"time"
)

func nonBlockingRecv() {
	fmt.Println("Non-blocking recv")
	fmt.Println("-----------------")
	ch := make(chan string, 1)
	fmt.Println("ch", ch)

	ch <- "hi"

	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch:
			fmt.Println("msg", msg)
		default:
			fmt.Println("no message on ch")
		}
	}
	fmt.Println("---")
}

func nonBlockingSend() {
	fmt.Println("Non-blocking send")
	fmt.Println("-----------------")

	ch := make(chan int)
	go func() {
		i := <-ch
		fmt.Println("recv from ch", i)
	}()

	// Wait for go routine to start...
	time.Sleep(50 * time.Millisecond)

	// Hmm, a non-blocking send seems weird to do with a select block...
	// But it's nice that send is symmetric to recv.

	// This one can send.
	select {
	case ch <- 1:
		fmt.Println("Able to send to ch")
	default:
		fmt.Println("Unable to send to ch")
	}

	// These cannot send as no-one is receiving.
	select {
	case ch <- 1:
		fmt.Println("Able to send to ch")
	default:
		fmt.Println("Unable to send to ch")
	}
	select {
	case ch <- 1:
		fmt.Println("Able to send to ch")
	default:
		fmt.Println("Unable to send to ch")
	}

	fmt.Println("---")
}

func main() {
	nonBlockingRecv()
	nonBlockingSend()
}

package main

import (
	"fmt"
	"time"
)

// unit, the "empty" tuple.
type unit struct{}

const sleepDuration = 1 * time.Second

func main0() {
	// Off to the races!
	for i := 0; i < 9; i++ {
		ch := make(chan string)
		go func() {
			time.Sleep(sleepDuration)
			ch <- "done"
		}()

		timeoutCh := make(chan unit)
		go func() {
			// Using this method of timing out, sometimes ch wins (i.e. gets to
			// deliver it's value).
			time.Sleep(sleepDuration)
			timeoutCh <- unit{}
		}()

		select {
		case v := <-ch:
			fmt.Println("main0: v", v)
		case t := <-timeoutCh:
			fmt.Println("main0: t", t)
		}
	}
}

func main1() {
	// Off to the races!
	for i := 0; i < 9; i++ {
		ch := make(chan string)
		go func() {
			time.Sleep(sleepDuration)
			ch <- "done"
		}()

		// Using `After` to timeout here with the same duration, causes the
		// operation to very rarely timeout. So, After must have a somewhat
		// different implementation to the above in `main0`.
		select {
		case v := <-ch:
			fmt.Println("main1: v", v)
		case t := <-time.After(sleepDuration):
			fmt.Println("main1: t", t)
		}
	}
}

func main2() {
	// Off to the races!
	for i := 0; i < 9; i++ {
		ch := make(chan string)
		go func() {
			time.Sleep(500 * time.Millisecond)
			ch <- "done"
		}()

		// Sleep an extra ms in the timeout to see if the "done" can win the
		// race.
		select {
		case v := <-ch:
			fmt.Println("main2: v", v)
		case t := <-time.After(501 * time.Millisecond):
			fmt.Println("main2: t", t)
		}
	}
}

var c chan int

func handle(int) {}

func eg() {
	select {
	case m := <-c:
		handle(m)
	case <-time.After(1 * time.Second):
		fmt.Println("timed out")
	}
}

func main() {
	fmt.Printf("sleepDuration.type = %T\n", sleepDuration)
	fmt.Printf("sleepDuration = %v\n", sleepDuration)
	eg()
	main0()
	main1()
	main2()
}

package main

import (
	"fmt"
	"time"
)

// unit, the "empty" tuple.
type unit struct{}

func main() {
	// Off to the races!
	for i := 0; i < 3; i++ {
		ch := make(chan string)
		go func() {
			time.Sleep(100 * time.Millisecond)
			ch <- "done"
		}()

		timeoutCh := make(chan unit)
		go func() {
			time.Sleep(100 * time.Millisecond)
			timeoutCh <- unit{}
		}()

		select {
		case v := <-ch:
			fmt.Println("v", v)
		case t := <-timeoutCh:
			fmt.Println("t", t)
		}
	}
}

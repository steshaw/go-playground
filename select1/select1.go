package main

import (
	"fmt"
	"time"
)

func main() {
	ca := make(chan rune)
	cb := make(chan rune)

	go func() {
		defer close(ca)
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			ca <- '✅'
		}
	}()
	go func() {
		defer close(cb)
		for i := 0; i < 10; i++ {
			time.Sleep(200 * time.Millisecond)
			cb <- '🙂'
		}
	}()

	for {
		select {
		case c := <-ca:
			fmt.Println("on ca", c)
			fmt.Printf("received on ca: %c\n", c)
		case c := <-cb:
			fmt.Println("on cb", c)
			fmt.Printf("received on cb: %c\n", c)
		default:
			fmt.Println("default")
		}
	}
}

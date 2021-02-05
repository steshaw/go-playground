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

	numItems := 0
	fmt.Printf("Items: ")
	for {
		select {
		case c, ok := <-ca:
			if ok {
				numItems++
				fmt.Printf("%c", c)
			} else {
				ca = nil
			}
		case c, ok := <-cb:
			if ok {
				numItems++
				fmt.Printf("%c", c)
			} else {
				cb = nil
			}
		}
		if ca == nil && cb == nil {
			break
		}
	}
	fmt.Printf("\n")
	fmt.Println("total items =", numItems)
}

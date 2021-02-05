package main

import "fmt"

func nonblockingReceive() {
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
}

func main() {
	nonblockingReceive()
}

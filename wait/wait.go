package main

import "fmt"

func main() {
	done := make(chan bool)
	go func() {
		fmt.Println("hi")
		done <- false // value is irrelevant
	}()

	shouldWait := true
	if shouldWait {
		result := <-done
		fmt.Println("result =", result)

	}
	fmt.Println("finished!")
}

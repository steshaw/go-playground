package main

import (
	"fmt"
)

func upTo(from, to int) func(func(int)) {
	return func(f func(int)) {
		for i := from; i < to; i++ {
			f(i)
		}
	}
}

func main() {
	fmt.Println("non-leaky iteration")
	for {
		upTo(2, 5)(func(a int) {
			// Cannot exit iteration early with "break".
			// break
		})
	}
}

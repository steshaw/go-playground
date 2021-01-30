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

func iter1() {
	for {
		upTo(2, 5)(func(i int) {
			fmt.Println(i)
			// Cannot exit iteration early with "break".
			// break
		})
	}
}

func iter2_aux() {
	upTo(2, 5)(func(i int) {
		fmt.Println(i)
		// Try to break iteration with `return`.
		// Unfortunately, here `return` only returns from the anonymous function
		// and not from `iter2_aux`.
		return
	})
}

func iter2() {
	for {
		iter2_aux()
	}
}

func main() {
	fmt.Println("non-leaky iteration")
	if false {
		iter1()
	}
	iter2()
}

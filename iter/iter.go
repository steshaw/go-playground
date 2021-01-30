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

func upTov2(from, to int) func(func(int) bool) {
	return func(f func(int) bool) {
		for i := from; i < to; i++ {
			quit := f(i)
			if quit {
				break
			}
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

func iter3() {
	upTov2(2, 5)(func(i int) bool {
		fmt.Println(i)
		// Explicitly end iteration by returning `quit = true`.
		return true
	})
}

func main() {
	fmt.Println("non-leaky iteration")
	if false {
		iter1()
	}
	if false {
		iter2()
	}

	iter3()
}

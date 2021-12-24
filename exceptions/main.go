package main

import (
	"fmt"
)

func foo(i int) int {
	// Try
	result, err := func() (ii int, e error) {
		defer func() {
			if except := recover(); except != nil {
				e = fmt.Errorf("caught panic: %v", except)
			}
		}()

		if i == 8 {
			panic("I don't like 8")
		} else {
			return i + 1, nil
		}
	}()

	// Catch the "exception" and return normally.
	if err != nil {
		fmt.Printf("Intercepted error: [[%v]]\n", err)
		return i
	} else {
		return result
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("foo(%d) = %d\n", i, foo(i))
	}
}

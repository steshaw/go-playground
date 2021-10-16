package main

import (
	. "example.com/option/option"
	"fmt"
)

func div(n, m int) Option {
	if m == 0 {
		return None
	} else {
		return Some(n / m)
	}
}

func main() {
	fmt.Println(div(42, 7))
	fmt.Println(div(42, 6))
	fmt.Println(div(3, 0))
}

package main

import (
	"fmt"
)

func Eq[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	fmt.Println(Eq(1, 1))
	fmt.Println(Eq(2, 1))
	fmt.Println(Eq("foo", "foo"))
	fmt.Println(Eq("foo", "fred"))
	fmt.Println(Eq('a', 'a'))
	fmt.Println(Eq('a', 'b'))
}

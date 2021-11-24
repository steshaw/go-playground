package main

import (
	"fmt"
)

func p[A any](as ...A) {
	for _, a := range as {
		fmt.Println(a)
	}
}

func main() {
	p("Hello", "there")
	p(42, 3, 9)
	p[interface{}]('a', 3, "foo")
}

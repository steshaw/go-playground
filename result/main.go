package main

import (
	"fmt"

	. "example.com/result/results"
)

func div(n, m int) Result {
	if m == 0 {
		return Err("Divide by zero!")
	} else {
		return Ok(n / m)
	}
}

func checkEq(a, b Result) {
	fmt.Printf("a=%v b=%v a.Eq(b)=%v\n", a, b, a.Eq(b))
	fmt.Printf("a=%v b=%v b.Eq(a)=%v\n", a, b, b.Eq(a))
}

func main() {
	fmt.Println(div(42, 7))
	fmt.Println(div(42, 6))
	fmt.Println(div(3, 0))

	checkEq(Ok(1), Ok(1))
	checkEq(Ok(1), Ok(2))
	checkEq(Ok('a'), Ok('b'))
	checkEq(Ok('a'), Ok('a'))
	checkEq(Ok("foo"), Ok("bar"))
	checkEq(Ok("hi"), Ok("hi"))
}

// TODO: Man or boy test.
// TODO: Upwards funarg example.
// TODO: Downwards funarg example.

package main

import "fmt"

func simple() {
	var i int = 0
	next := func() int {
		i++
		return i
	}

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

func nextFrom(i int) func() int {
	next := i
	return func() int {
		result := next
		next++
		return result
	}
}

func next() {
	g0 := nextFrom(0)
	g1 := nextFrom(1000)
	fmt.Println("g0 =", g0())
	fmt.Println("g1 =", g1())
	fmt.Println("g0 =", g0())
	fmt.Println("g1 =", g1())
	fmt.Println("g0 =", g0())
	fmt.Println("g1 =", g1())
	fmt.Println("g1 =", g1())
	fmt.Println("g1 =", g1())
	fmt.Println("g0 =", g0())
	fmt.Println("g0 =", g0())
}

func main() {
	simple()
	next()
}

package main

import "fmt"

func main() {
	var i int = 0
	next := func() int {
		i++
		return i
	}

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

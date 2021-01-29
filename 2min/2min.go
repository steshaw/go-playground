package main

import (
	"fmt"
)

func main() {
	print("Hi")
	fmt.Println("Hi")

	{
		var i int
		fmt.Println(i)
	}

	{
		var s string
		fmt.Println(s)
	}
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")

	{
		var i int
		fmt.Println(i)
	}

	{
		var s string
		fmt.Println(s)
	}

	{
		for i := 0; i < 5; i++ {
			fmt.Println("i = %s", i)
		}
	}
}

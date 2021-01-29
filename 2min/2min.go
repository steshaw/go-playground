package main

import (
	"fmt"
)

func PrintRange(is []int) {
	fmt.Println("range {")
	for x, i := range is {
		fmt.Printf("  %d : %d\n", x, i)
	}
	fmt.Println("}")
}

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
		for i := 10; i < 15; i++ {
			fmt.Printf("i = %d\n", i)
		}
	}

	/*
		{
			for x, i := 10; i < 15; i++ {
				fmt.Println(x, ": ", i)
			}
		}
	*/

	var as = [5]int{10, 11, 12, 13, 14}
	fmt.Println(as)
	{
		for a := range as {
			fmt.Printf("a = %d\n", a)
		}
	}
	{
		for i, a := range as {
			fmt.Printf("%d: %d\n", i, a)
		}
	}
	PrintRange([]int{})
	PrintRange([]int{2})
	PrintRange([]int{2, 4})
	PrintRange([]int{2, 4, 6})
}

package main

import (
	"fmt"
)

func PrintRange(name string, is []int) {
	fmt.Printf("%s = range {\n", name)
	for x, i := range is {
		fmt.Printf("  %d : %d\n", x, i)
	}
	fmt.Println("}")
}

func main() {
	fmt.Println("Hi")
	fmt.Println("Hi" + " " + "there")
	fmt.Println("6 * 7 = ", 6*7)
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println(true || false)
	fmt.Println(true && true)
	fmt.Println(1 == 2)
	fmt.Println(0 == 0)
	fmt.Println(3.0 / 9.0)
	fmt.Println(3.0 == 3.0)
	fmt.Println(3.0 <= 3.0)

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
	PrintRange("i0", []int{})
	PrintRange("i1", []int{2})
	PrintRange("i2", []int{2, 4})
	PrintRange("i3", []int{2, 4, 6})
	var bs []int = as[0:1]
	PrintRange("bs", bs)
	PrintRange("as", as[0:])
}

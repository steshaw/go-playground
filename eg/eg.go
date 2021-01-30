package main

import (
	"fmt"
	"strings"
)

func PrintRange(name string, is []int) {
	fmt.Printf("%s = range {\n", name)
	for x, i := range is {
		fmt.Printf("  %d : %d\n", x, i)
	}
	fmt.Println("}")
}

func channels() {
	fmt.Printf("\n\n")
	fmt.Println("Channels")
	fmt.Println("--------")
	c := make(chan string)
	go func() {
		c <- "one"
		c <- "two"
		c <- "hi"
		c <- " "
		c <- "there"
		// Channel must be closed to use range to avoid deadlock.
		// Channel must be closed in this groutine rather than in the receiving
		// routine to avoid panic "send on closed channel".
		close(c)
		if false {
			// Don't close a channel twice to avoid panic!
			close(c)
		}
	}()

	m1 := <-c
	m2 := <-c
	fmt.Printf("m1 = %s\n", m1)
	fmt.Printf("m2 = %s\n", m2)

	// Do not close channel here to avoid panic "send on closed channel".
	if false {
		close(c)
	}
	fmt.Println("msgs {")
	for msg := range c {
		fmt.Println("  ", msg)
	}
	fmt.Println("}")
}

func upTo(from, to int) chan int {
	var c chan int = make(chan int)
	go func() {
		for i := from; i < to; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func upto() {
	fmt.Println("UpTo(10, 15) = {")
	var ix int
	for i := range upTo(10, 15) {
		fmt.Println("  ", ix, ": ", i)
		ix++
	}
	fmt.Println("}")
}

// Go doesn't have a ternary operator.
func ternary1(i int) string {
	var size string
	if i > 100 {
		size = "big"
	} else {
		size = "small"
	}
	return size
}
func ternary2(i int) string {
	if i > 100 {
		return "big"
	} else {
		return "small"
	}
	// Should be unreachable.
	return "unreachable"
}

func main() {
	fmt.Println("Hi")
	fmt.Println("Hi" + " " + "there")
	fmt.Println("6 * 7 = ", 6*7)
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println(true || false)
	fmt.Println(1 == 2)
	fmt.Println(0 == 0)
	fmt.Println(3.0 / 9.0)
	fmt.Println(3.0 == 3.0)
	fmt.Println(3.0 <= 3.0)

	{
		var i int
		fmt.Println(i)
	}

	header := func(title string) {
		fmt.Printf("\n\n")
		fmt.Println(title)
		fmt.Println(strings.Repeat("-", len(title)))
	}

	{
		header("Strings")
		var s string
		fmt.Printf("s = %s\n", s)
		s = "fred"
		fmt.Printf("s = %s\n", s)
		const s1 string = "constant"
		fmt.Printf("s1 = %s\n", s1)
		fmt.Println()
	}

	{
		for i := 10; i < 15; i++ {
			fmt.Printf("i = %d\n", i)
		}
	}

	upto()

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

	channels()

	if num := 7; num < 0 {
		fmt.Println("< 0")
	} else if num > 0 {
		fmt.Println("> 0")
	} else if num == 0 {
		fmt.Println("== 0")
	}
	// `num` is inaccessible here
	//fmt.Println(num)

	fmt.Println(ternary1(32))
	fmt.Println(ternary2(142))
}

package main

import (
	"fmt"
	"strconv"
)

func classify0(i int) string {
	switch {
	case i > 0:
		return "above"
	case i < 0:
		return "below"
	default:
		return "zero"
	}
}
func classify1(i int) string {
	switch a := i > 0; a {
	case true:
		return strconv.FormatBool(a) + " " + "above"
	case false:
		return strconv.FormatBool(a) + " " + "hmmm"
	default:
		return strconv.FormatBool(a) + " " + "unreachable"
	}
}

func do(classify func(int) string) {
	fmt.Println("result = ", classify(3))
	fmt.Println("result = ", classify(-32))
	fmt.Println("result = ", classify(0))
}
func main() {
	do(classify0)
	do(classify1)
}

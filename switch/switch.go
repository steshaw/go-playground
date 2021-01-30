package main

import "fmt"

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
	switch i > 0 {
	case true:
		return "above"
	case false:
		return "hmm"
	default:
		return "unreachable"
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

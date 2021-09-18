package main

import "fmt"

func hmm(a interface{}) {
	fmt.Printf("%v\n", a)
}

func main() {
	hmm(42)
	hmm("asdf")
}

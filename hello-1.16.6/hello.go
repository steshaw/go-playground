package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Printf("Hello = %s\n", quote.Hello())
	fmt.Printf("Go = %s\n", quote.Go())
	fmt.Printf("Glass = %s\n", quote.Glass())
	fmt.Printf("Opt = %s\n", quote.Opt())
}

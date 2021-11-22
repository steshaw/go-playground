package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("i = %d\n", i)
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

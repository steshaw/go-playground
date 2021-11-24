package main

import (
	"fmt"
	"log"
)

func main() {
	println("println'ed") // stderr
	fmt.Println("fmt'ed") // stdout
	log.Println("log'ed") // stderr
}

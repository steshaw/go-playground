package main

import (
	"fmt"
)

func main() {
	fn := func() { fmt.Print(1) }
	defer fn()
	fn = func() { fmt.Print(2) }
	defer fn()
}

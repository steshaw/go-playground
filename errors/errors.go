package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("An example error")
	fmt.Println(e)
}

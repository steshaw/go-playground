package main

import "fmt"

type Fooer interface {
	Foo() string
}

type Container struct {
	Fooer
}

func doFoo(f Fooer) {
	fmt.Println(f.Foo())
}

func main() {
	doFoo(Container{})
}

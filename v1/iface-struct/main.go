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

type foo struct{}

func (foo) Foo() string {
	return "foo!"
}

type bar struct{}

func (bar) Foo() string {
	return "bar!"
}

func panicky() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered r=%v\n", r)
		}
	}()
	doFoo(Container{})
}

func main() {
	panicky()
	c := Container{Fooer: foo{}}
	doFoo(c)
	c.Fooer = bar{}
	doFoo(c)
}

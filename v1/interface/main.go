package main

import "fmt"

type Helloer interface {
	Hello()
}

type A struct{}

func (a *A) Hello() {
	fmt.Printf("Hi from A, a = %v\n", a)
}

type Embedded1 struct {
	A
	i int
}

type Embedded2 struct {
	A
	i int
}

func (e2 *Embedded2) Hello() {
	fmt.Printf("Hi from Embedded2, e2 = %v\n", e2)
}

func main() {
	a := &A{}
	a.Hello()

	e1 := Embedded1{i: 42}
	e1.Hello() // Using Helloer impl from the embedded A!

	e2 := Embedded2{i: 3}
	e2.Hello()
	e2.A.Hello()
}

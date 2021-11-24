// https://groups.google.com/g/golang-dev/c/r1djftRyEDc

package main

import "fmt"

type Fooer interface {
	Foo()
}

type Foo bool

func (f Foo) Foo() {}

type Bar[t any] t

func bar[t any](b Bar[t]) {
	fmt.Println("I can get a higher rank type...")
}
func (b Bar[t]) String() string {
	return fmt.Sprint("...but I can't make it a fmt.Stringer")
}

func main() {
	b := Bar[Fooer](Foo(true))
	fmt.Printf("%T %v\n", b, b)
	bar[Fooer](b)
}

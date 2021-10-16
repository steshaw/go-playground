package main

import (
	"fmt"
)

type Show interface {
  Show() string
}

type String string

func (a String) Show() string {
  return string(a)
}

type Int int

func (a Int) Show() string {
  return fmt.Sprint(a)
}

type Rune rune

func (a Rune) Show() string {
  return fmt.Sprint(a)
}

func show[A Show](as ...A) {
	for _, a := range as {
		fmt.Println(a.Show())
	}
}

func main() {
	show(String("Hello"), String("there"))
	show(Int(42), Int(3), Int(9))
	show(Show(Rune('a')), Show(Int(3)), Show(String("foo")))
}

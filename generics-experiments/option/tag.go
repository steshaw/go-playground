// https://go2goplay.golang.org/p/_KxeaTbSkDs

package main

import "fmt"

type Option[T any] struct {
	ok bool
	x T
}

func (o Option[T]) Get() (T, bool) {
	return o.x, o.ok
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func Some[T any](x T) Option[T] {
	return Option[T]{
		ok: true,
		x: x,
	}
}

func inspect[T any](option Option[T]) {
	if x, ok := option.Get(); ok {
		fmt.Printf("Yes %v\n", x)
	} else {
		fmt.Println("No")
	}

}

func main() {
	v := Some(1)
	n := None[int]()
	inspect(v)
	inspect(n)
}

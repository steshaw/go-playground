package main

import 	"fmt"


type Option[T any] struct {
	p interface{}
}

func (o Option[T]) Get() (T, bool) {
	if o.p != nil {
		return o.p.(T), true
	}
	return *new(T), false
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func Some[T any](x T) Option[T] {
	return Option[T]{
		p: x,
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

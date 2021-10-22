// https://go2goplay.golang.org/p/_KxeaTbSkDs

package main

import "fmt"

type Option[T any] struct {
	isNone bool
	a      T
}

type Unit struct{}

func unit() Unit {
	return Unit{}
}

func Match[t, r any](o Option[t], n func() r, s func(t) r) r {
	if o.isNone {
		return n()
	}
	return s(o.a)
}

func None[a any]() Option[a] {
	return Option[a]{isNone: true}
}

func Some[A any](a A) Option[A] {
	return Option[A]{
		isNone: false,
		a:      a,
	}
}

func inspect[A any](option Option[A]) {
	Match(option, func() Unit {
		fmt.Println("No")
		return unit()
	}, func(a A) Unit {
		fmt.Printf("Yes %v\n", a)
		return unit()
	})
}

func main() {
	v := Some(1)
	n := None[int]()
	inspect(v)
	inspect(n)
}

// https://go2goplay.golang.org/p/_KxeaTbSkDs

package main

import "fmt"

type option[A any] struct {
	isnone bool
	a      A
}

type Unit struct{}

func unit() Unit {
	return Unit{}
}

func Match[A, B any](o option[A], n func() B, s func(A) B) B {
	if o.isnone {
		return n()
	}
	return s(o.a)
}

func none[a any]() option[a] {
	return option[a]{isnone: true}
}

func some[A any](a A) option[A] {
	return option[A]{
		isnone: false,
		a:      a,
	}
}

func (o option[a]) String() string {
	return Match(o, func() string {
		return "None"
	}, func(a a) string {
		return fmt.Sprintf("Some(%v)", a)
	})
}

func printOption[A any](option option[A]) {
	Match(option, func() Unit {
		fmt.Println("No")
		return unit()
	}, func(a A) Unit {
		fmt.Printf("Yes %v\n", a)
		return unit()
	})
}

func ps(optString option[string]) {
	fmt.Println(optString)
}

func main() {
	fmt.Println(some("fred"))
	fmt.Println(none[string]())
	ps(some("wilma"))
	ps(none[string]())

	printOption(some(1))
	printOption(some(2))
	printOption(some("hi"))
	printOption(none[rune]())
}

package main

import (
	"fmt"
)

// The playground now uses square brackets for type parameters. Otherwise,
// the syntax of type parameter lists matches the one of regular parameter
// lists except that all type parameters must have a name, and the type
// parameter list cannot be empty. The predeclared identifier "any" may be
// used in the position of a type parameter constraint (and only there);
// it indicates that there are no constraints.

type folder[A][B] func(b B, a A) B

/*
type FoldLefter[A] interface {
  FoldLeft[
}
*/

type List[A] interface {
	foldLeft[B](f folder[A, B], b B) B
}

type Nil[A] struct{}
type Cons[A] struct {
	car A
	cdr List[A]
}

func (as Nil[A]) foldLeft[B any](f folder[A, B], b B) B {
	return b
}

func nil[A any]() List[A] { return Nil[A]{} }

func cons[A any](a A, as List[A]) List[A] {
	return Cons[A]{car: a, cdr: as}
}

// (b -> a -> b) -> b -> t a -> b

func printAll[A any](as List[A]) {
	as.foldLeft(
		func(b, a) { return b + a },
		0,
		xs,
	)
}

func Print[A any](as []A) {
	for _, a := range as {
		fmt.Println(a)
	}
	println()
}

func main() {
	var xs = cons(
		1, cons(
			2, cons(
				3, nil[int]())))

	Print([]string{"Hello", "playground"})
	Print([]int{1, 2, 3})
}




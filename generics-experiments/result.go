package main

import (
	"fmt"
)

// The Result algebra.
type resultAlg[T comparable, E comparable] struct {
	ok  func(T)
	err func(E)
}

type Result[T comparable, E comparable] func(resultAlg[T, E])

func Err[T comparable, E comparable](e E) Result[T, E] {
	return func(resultAlg resultAlg[T, E]) {
		resultAlg.err(e)
	}
}

func Ok[T comparable, E comparable](a T) Result[T, E] {
	return func(resultAlg resultAlg[T, E]) {
		resultAlg.ok(a)
	}
}

func (self Result[T, E]) String() string {
	var result string
	self(resultAlg[T, E]{
		ok:  func(a T) { result = fmt.Sprintf("Ok(%v)", a) },
		err: func(a E) { result = fmt.Sprintf("Err(%v)", a) },
	})
	return result
}

func konst[T any, B any](b B, bp *B) func(T) {
	return func(_ T) { *bp = b }
}

func (r Result[T, E]) Eq(r2 Result[T, E]) bool {
	var result bool
	r(resultAlg[T, E]{
		ok: func(a T) {
			r2(resultAlg[T, E]{
				ok:  func(b T) { result = a == b },
				err: konst[E, bool](false, &result),
			})
		},
		err: func(e E) {
			r2(resultAlg[T, E]{
				ok:  konst[T, bool](false, &result),
				err: func(e2 E) { result = e == e2 },
			})
		},
	})
	return result
}

func div(n int, m int) Result[int, string] {
	if m == 0 {
		return Err[int, string]("Divide by zero!")
	} else {
		return Ok[int, string](n / m)
	}
}

func checkEq[T comparable, E comparable](a, b Result[T, E]) {
	fmt.Printf("a=%v b=%v a.Eq(b)=%v\n", a, b, a.Eq(b))
	fmt.Printf("a=%v b=%v b.Eq(a)=%v\n", a, b, b.Eq(a))
}

func checkEqs[T comparable, E comparable](rs []Result[T, E]) {
	fmt.Printf("Checking %v\n", rs)
	for _, r1 := range rs {
		for _, r2 := range rs {
			checkEq(r1, r2)
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println(div(42, 7))
	fmt.Println(div(42, 6))
	fmt.Println(div(3, 0))
	fmt.Println()

	checkEqs([]Result[int, string]{
		Ok[int, string](1),
		Ok[int, string](2),
		Err[int, string]("ouf!"),
		Err[int, string]("nah!"),
	})

	checkEqs([]Result[int, string]{
		Ok[int, string]('a'),
		Ok[int, string]('b'),
		Err[int, string]("ouch!"),
		Err[int, string]("argh!"),
	})

	checkEqs([]Result[string, error]{
		Ok[string, error]("foo"),
		Ok[string, error]("bar"),
		Err[string, error](fmt.Errorf("naf!")),
		Err[string, error](fmt.Errorf("nagh!")),
	})
}

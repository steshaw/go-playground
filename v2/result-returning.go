// This is really out of control.

package main

import "fmt"

// The Result algebra.
type resultAlg[T comparable, E comparable, B any] struct {
	ok  func(T) B
	err func(E) B
}

type Result[T comparable, E comparable, B any] func(resultAlg[T, E, B]) B

func Err[T comparable, E comparable, B any](e E) Result[T, E, B] {
	return func(resultAlg resultAlg[T, E, B]) B {
		return resultAlg.err(e)
	}
}

func Ok[T comparable, E comparable, B any](a T) Result[T, E, B] {
	return func(resultAlg resultAlg[T, E, B]) B {
		return resultAlg.ok(a)
	}
}

func toString[T comparable, E comparable](r Result[T, E, string]) string {
	return r(resultAlg[T, E, string]{
		ok: func(a T) string {
			return fmt.Sprintf("Ok(%v)", a)
		},
		err: func(e E) string {
			return fmt.Sprintf("Err(%v)", e)
		},
	})
}

func konst[T any, B any](b B) func(T) B {
	return func(_ T) B { return b }
}

func Eq[T comparable, E comparable](r Result[T, E, bool], r2 Result[T, E, bool]) bool {
	return r(resultAlg[T, E, bool]{
		ok: func(a T) bool {
			return r2(resultAlg[T, E, bool]{
				ok: func(b T) bool { return a == b },
				//err: konst[E, bool](false),
				err: func(e E) bool { return false },
			})
		},
		err: func(e E) bool {
			return r2(resultAlg[T, E, bool]{
				ok:  konst[T, bool](false),
				err: func(e2 E) bool { return e == e2 },
			})
		},
	})
}

func div[B any](n int, m int) Result[int, string, B] {
	if m == 0 {
		return Err[int, string, B]("Divide by zero!")
	} else {
		return Ok[int, string, B](n / m)
	}
}

func checkEq[T comparable, E comparable](
	aBool, bBool Result[T, E, bool],
	aString, bString Result[T, E, string],
) {
	fmt.Printf("a=%v b=%v Eq(a, b)=%v\n",
		toString(aString),
		toString(bString),
		Eq(aBool, bBool),
	)
	fmt.Printf("a=%v b=%v Eq(b, a)=%v\n",
		toString(aString),
		toString(bString),
		Eq(bBool, aBool),
	)
}

func checkEqs[T comparable, E comparable](rsB []Result[T, E, bool], rsS []Result[T, E, string]) {
	//fmt.Printf("Checking %v\n", rsS)
	// strings := rsS.map(toString)
	strings := []string{}
	for _, r := range rsS {
		strings = append(strings, toString(r))
	}
	fmt.Printf("Checking %v\n", strings)

	for i, r1 := range rsB {
		for j, r2 := range rsB {
			r1s := rsS[i]
			r2s := rsS[j]
			checkEq[T, E](r1, r2, r1s, r2s)
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println(toString[int, string](Err[int, string, string]("oh!")))

	fmt.Println(toString(div[string](42, 7)))
	fmt.Println(toString(div[string](42, 6)))
	fmt.Println(toString(div[string](3, 0)))
	fmt.Println()

	checkEqs([]Result[int, string, bool]{
		Ok[int, string, bool](1),
		Ok[int, string, bool](2),
		Err[int, string, bool]("ouf!"),
		Err[int, string, bool]("nah!"),
	}, []Result[int, string, string]{
		Ok[int, string, string](1),
		Ok[int, string, string](2),
		Err[int, string, string]("ouf!"),
		Err[int, string, string]("nah!"),
	})

	/*
		checkEqs([]Result[int, string, bool]{
			Ok[int, string, bool]('a'),
			Ok[int, string, bool]('b'),
			Err[int, string, bool]("ouch!"),
			Err[int, string, bool]("argh!"),
		})

		checkEqs([]Result[string, error, bool]{
			Ok[string, error, bool]("foo"),
			Ok[string, error, bool]("bar"),
			Err[string, error, bool](fmt.Errorf("naf!")),
			Err[string, error, bool](fmt.Errorf("nagh!")),
		})
	*/
}

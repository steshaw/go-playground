package result

import (
	"fmt"
)

type any interface{}

// The Result algebra.
type resultAlg struct {
	err func(any)
	ok  func(any)
}

type Result func(resultAlg)

func Err(a any) Result {
	return func(resultAlg resultAlg) {
		resultAlg.err(a)
	}
}
func Ok(a any) Result {
	return func(resultAlg resultAlg) {
		resultAlg.ok(a)
	}
}

func (self Result) String() string {
	var result string
	self(resultAlg{
		err: func(a any) { result = fmt.Sprintf("Err(%v)", a) },
		ok:  func(a any) { result = fmt.Sprintf("Ok(%v)", a) },
	})
	return result
}

func (r Result) Eq(r2 Result) bool {
	var result bool
	konst := func(b bool) func(any) {
		return func(_ any) { result = b }
	}
	r(resultAlg{
		err: func(a any) {
			r2(resultAlg{
				err: konst(true),
				ok:  konst(false),
			})
		},
		ok: func(a any) {
			r2(resultAlg{
				err: konst(false),
				ok:  func(b any) { result = a == b },
			})
		},
	})
	return result
}

package option

import (
	"fmt"
)

type any interface{}

// The Option algebra.
type optionAlg struct {
	none func()
	some func(any)
}

type Option func(optionAlg)

var (
	None = Option(func(alg optionAlg) { alg.none() })
)

func Some(a any) Option {
	return func(optionAlg optionAlg) {
		optionAlg.some(a)
	}
}

func (self Option) String() string {
	var result string
	self(optionAlg{
		none: func() { result = "None" },
		some: func(a any) { result = fmt.Sprintf("Some(%v)", a) },
	})
	return result
}

func (r Option) Eq(r2 Option) bool {
	var result bool
	t := func() { result = true }
	f := func() { result = false }
	konst := func(b bool) func(any) {
		return func(_ any) { result = b }
	}
	r(optionAlg{
		none: func() {
			r2(optionAlg{
				none: t,
				some: konst(false),
			})
		},
		some: func(a any) {
			r2(optionAlg{
				none: f,
				some: func(b any) { result = a == b },
			})
		},
	})
	return result
}

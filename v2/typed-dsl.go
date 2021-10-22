// Argh! This is currently broken on gotip.

// This file implements a deep embedding of a typed-DSL in Go. The
// representation is type-safe (we cannot construct ill-typed terms)
// and accepts multiple interpretations. The type system of the target
// language is identity-mapped to the Go type system such that type
// checking of the DSL is hoisted up to type-checking the Go code that
// contains the target language expression.
//
// Normally this requires either GADTs or higher-rank types. I show
// that it is possible to encode it in Go, a language which doesn't
// have GADTs (nor regular ADTs for that matter), nor higher-rank
// types. I exploit the duality between universal and existential
// quantification and encode sum types using the existential dual of the
// Boehm-Berarducci isomorphism. Unlike the Boehm-Berarducci encoding,
// my encoding is not universally-quantified, but existentially
// quantified, and does not require higher-rank polymorphism capitalizing
// on the fact that Go interfaces are existential types.
//
// Just like an algebraic data type, my encoding is closed, its usage
// is type safe, and the match operations are checked at compile-time
// for exhaustivness.
//
// A related, alternative encoding would be to encode the GADT into
// tagless-final style. This requires polymorphic terms, which in Go
// can only be functions, which are not serializable. My encoding
// is bidirectionally serializable.
//
// As presented, the encoding is closed because I want to show that
// I can encode every GADT property. It is also possible, and perhaps
// desirable to make the encoding open, which solves the Expression
// Problem.
//
// Although GADT invariants are encoded using Go generics (which are
// a form of universal quantification) the encoding does not require
// that the Curry–Howard isomorphism holds, it is purely existential.
// In fact, if one only wants plain ADTs, generics are not needed at
// all. Go can encode sum types, and was always able to.
//
// We will encode the following GADT (expressed in Haskell)
//
// data Exp a where
//     Lit :: a -> Exp a
//     Add :: Exp Int -> Exp Int -> Exp Int
//     Or  :: Exp Bool -> Exp Bool -> Exp Bool
//     Lt  :: Exp Int -> Exp Int -> Exp Bool
//     If  :: Exp Bool -> Exp a -> Exp a -> Exp a

package main

import (
	"fmt"
)

// An Exp represents a serializable DSL expression. A sum type is
// nothing but the simplest non-degenerate abstract data type. It can
// only be unpacked, and I make this explicit.
type Exp interface {
	// MatchExp unpacks the ADT and continues processing. ExpBoolVars
	// and ExpIntVars are sets of all possible continuations.
	// ExpKind is not strictly required by the encoding, but it's
	// convenient, it tells us the kind of continuation that was
	// followed in case we need this information later.
	MatchExp(ExpBoolVars, ExpIntVars) ExpKind
}

// ety is a constraint that fixes the internal DSL's denotational
// semantics. Note that the semantics given by an arbitrary external
// interpreter are not constrained by this. This has no consequence
// for the consumer, but it can improve the safety for the producer,
// us, the human compiler of the ADT. If Go had bounded universal
// quantification for methods, this safety guarantee would come for
// free.
type ety interface {
	type bool, int
}

// exp[t] is typed expression of DSL-type t, which is also the Go type
// t. Unlike Exp, exp[t] is not serializable. The reason we can keep
// Exp serializable while still having a typed representation is that
// t is a phantom type. It exists purely at compile-time to participate
// in type-checking, but it doesn't leave any record at runtime. The
// encoding is tagless. As such, we can hide it behind an existential
// and have heterogenous lists of expressions.
type exp[t ety] interface {
	Exp
	is(exp[t])
}

// ExpBoolVars are sets of continuations for the possible processings
// of the ADT. They are the projections of the ADT constructors over
// its denotations reclasified by the denotations. Notice the type
// refinement.
//
// For this to work, the set of classes of internal denotations must
// be finite. Not shown here, but the constructors could still be
// polymorphic, potentially quantifying over an infinite set of
// denotations. Only the set of classes (kinds) need to be finite.

// ExpBoolVars are the possible continuations after boolean expression
// matching.
type ExpBoolVars interface {
	Lit(bool)
	Or(exp[bool], exp[bool])
	Lt(exp[int], exp[int])
	If(exp[bool], exp[bool], exp[bool])
}

// ExpIntVars are the possible continuations after integer expression
// matching.
type ExpIntVars interface {
	Lit(int)
	Add(exp[int], exp[int])
	If(exp[bool], exp[int], exp[int])
}

// ExpKind is an ADT (yes, another ADT) that classifies expressions by
// their type.
type ExpKind interface {
	MatchKind(KindVars)
}

// KindVars are the possible continuations after matching the result
// of matching an expression.
type KindVars interface {
	Bool(ExpBoolVars)
	Int(ExpIntVars)
}

// boolVal captures the continuation of matching a boolean expression.
// It implements ExpKind.
type boolVal struct{ ev ExpBoolVars }

// intVal captures the continuation of matching an integer expression.
// It implements ExpKind.
type intVal struct{ ev ExpIntVars }

func (b boolVal) MatchKind(v KindVars) { v.Bool(b.ev) }
func (n intVal) MatchKind(v KindVars)  { v.Int(n.ev) }

// MatchIntExp is a convenience function that allows one to reuse the
// captured data of a continuation provided we know the expression
// type is integer without having to re-match on the expression's kind.
func MatchIntExp[tv ExpIntVars](e exp[int], v tv) tv {
	e.MatchExp(nil, v)
	return v
}

// MatchBoolExp is a convenience function that allows one to reuse the
// captured data of a continuation provided we know the expression
// type is boolean without having to re-match on the expression's kind.
func MatchBoolExp[tv ExpBoolVars](e exp[bool], v tv) tv {
	e.MatchExp(v, nil)
	return v
}

// lit[t], add, or, lt, and if_ encode typed expressions. If an ADT
// is interpreted as an abstract type, then these are the concrete
// types of its values. These types implement exp[t], the type of typed
// expressions.

type lit[t ety] t
type add struct{ a, b exp[int] }
type or struct{ a, b exp[bool] }
type lt struct{ a, b exp[int] }
type if_[t ety] struct {
	c exp[bool]
	t exp[t]
	f exp[t]
}

// The MatchExp methods implement exp[t]. Each function selects the
// appropiate continuations, continues it, then returns an ADT containing
// the kind of the continuation set. These functions must be injective.

func (l lit[t]) MatchExp(bv ExpBoolVars, nv ExpIntVars) ExpKind {
	// We use a type switch to get around the fact that we do not
	// have bounded universal quantification for methods. This
	// type switch is safe, because we know the bound for t.
	// Unfortunately, the Go compiler can't yet prove that this
	// switch is exhaustive, hence the default case.
	//
	// Note that we are acting as a compiler here, the consumer
	// never has to write any type switches.
	switch v := interface{}(l).(type) {
	case lit[bool]:
		bv.Lit(bool(v))
		return boolVal{bv}
	case lit[int]:
		nv.Lit(int(v))
		return intVal{nv}
	default:
		panic("unreachable")
	}
}
func (n add) MatchExp(_ ExpBoolVars, v ExpIntVars) ExpKind {
	v.Add(n.a, n.b)
	return intVal{v}
}
func (b or) MatchExp(v ExpBoolVars, _ ExpIntVars) ExpKind {
	v.Or(b.a, b.b)
	return boolVal{v}
}
func (b lt) MatchExp(v ExpBoolVars, _ ExpIntVars) ExpKind {
	v.Lt(b.a, b.b)
	return boolVal{v}
}
func (e if_[t]) MatchExp(bv ExpBoolVars, nv ExpIntVars) ExpKind {
	switch interface{}(*new(t)).(type) {
	case bool:
		bv.If(e.c, e.t.(exp[bool]), e.f.(exp[bool]))
		return boolVal{bv}
	case int:
		nv.If(e.c, e.t.(exp[int]), e.f.(exp[int]))
		return intVal{nv}
	default:
		panic("unreachable")
	}
}

// The is functions enforce that an expression is not invariant over
// its type. Ideally we would return an exp[t] by converting the
// receiver to it. Then, by the virtue of the Curry–Howard corespondence
// this would be a compile-time proof that the receiver is in fact
// what we claim it is. Unfortunately, due to a bug in Go we can't yet
// do that, see issue 47887.

func (l lit[t]) is(exp[t]) {}
func (n add) is(exp[int])  {}
func (b or) is(exp[bool])  {}
func (b lt) is(exp[bool])  {}
func (e if_[t]) is(exp[t]) {}

// We implement fmt.Stringer for our value constructors. This is a
// fixed Exp interpretation (more on that when we get to Eval).

func (l lit[t]) String() string { return fmt.Sprintf("(lit %v)", t(l)) }
func (n add) String() string    { return fmt.Sprintf("(add %v %v)", n.a, n.b) }
func (b or) String() string     { return fmt.Sprintf("(or %v %v)", b.a, b.b) }
func (b lt) String() string     { return fmt.Sprintf("(lt %v %v)", b.a, b.b) }
func (e if_[t]) String() string { return fmt.Sprintf("(if %v %v %v)", e.c, e.t, e.f) }

// We implement the value constructors as functions instead of letting
// the user create the types directly because type-inference in Go
// works better this way. Unfortunately, it still doesn't work for If,
// which is very dissapointing.
//
// If Go used Hindley–Milner type-inference, this would not be a
// problem.

func Lit[t ety](v t) lit[t] { return lit[t](v) }
func Add(a, b exp[int]) add { return add{a, b} }
func Or(a, b exp[bool]) or  { return or{a, b} }
func Lt(a, b exp[int]) lt   { return lt{a, b} }
func If[t ety](c exp[bool], tr, fl exp[t]) if_[t] {
	return if_[t]{c, tr, fl}
}

// Everything that was written above this line was the producer's side,
// the side of the human compiler if you will. Below this line is the
// side of the consumer. In particular, this could be implemented in
// a different package. The code that follows is completely oblivious
// to the internal representation of the ADT. Our sum type is an
// abstract data type.

// We want to write an interpretation for the ADT that evalues it in
// big-step operational semantics style. This is *our* choice. The
// producer of the ADT doesn't force eval on us (like it forced
// fmt.Stringer). The interpretation is not fixed, for example, we
// could choose to make another interpreter that's a simplifer, or one
// that's a compiler to machine code. All these have a different type,
// and all are possible.
//
// 	    eval :: Exp a -> a
// 	simplify :: Exp a -> Exp a
// 	 compile :: Exp a -> [Word]
//
// We only implement eval, the rest are left as an exercise to the
// reader.

// EvalInt is a set of mutally recursive continuations (implemented
// as an object satisfying ExpIntVars) that capture the denotation of
// the expression as an int.
//
// EvalInt provides a big-step operational interpretation for an
// exp[int] that returnes its evaluated value when participating in
// matching.
//
// EvalInt implements ExpIntVars, so the compiler checks that we
// implement all required continuations (match cases). We cannot forget
// to implement a case. It won't compile.
type EvalInt int

// EvalBool is a set of mutally recursive continuations (implemented
// as an object satisfying ExpBoolVars) that capture the denotation of
// the expression as an bool.
//
// EvalBool provides a big-step operational interpretation for an
// exp[bool] that returnes its evaluated value when participating in
// matching.
//
// EvalBool implements ExpBoolVars, so the compiler checks that we
// implement all required continuations (match cases). We cannot forget
// to implement a case. It won't compile.
type EvalBool bool

func (ev *EvalInt) Lit(n int) { *ev = EvalInt(n) }
func (ev *EvalInt) Add(a, b exp[int]) {
	*ev = evalInt(a) + evalInt(b)
}
func (ev *EvalInt) If(c exp[bool], t, f exp[int]) {
	if evalBool(c) {
		*ev = evalInt(t)
		return
	}
	*ev = evalInt(f)
}

// evalInt evaluates an exp[int] returning its value as an EvalInt.
func evalInt(e exp[int]) EvalInt {
	// We know we have an exp[int], so we can use the simpler
	// MatchIntExp instead of Exp.MatchExp. Note that the compiler
	// still checks this. We can't pass an exp[bool], for example,
	// nor can we call MatchBoolExp by mistake. Here, on the
	// consumer side, the interface is impossible to misuse.
	return *MatchIntExp(e, new(EvalInt))
}

func (ev *EvalBool) Lit(b bool) { *ev = EvalBool(b) }
func (ev *EvalBool) Or(a, b exp[bool]) {
	*ev = evalBool(a) || evalBool(b)
}
func (ev *EvalBool) Lt(a, b exp[int]) {
	// Note that we call evalInt here!
	if evalInt(a) < evalInt(b) {
		*ev = EvalBool(true)
		return
	}
	*ev = EvalBool(false)
}
func (ev *EvalBool) If(c, t, f exp[bool]) {
	if evalBool(c) {
		*ev = evalBool(t)
		return
	}
	*ev = evalBool(f)
}
func evalBool(e exp[bool]) EvalBool {
	return *MatchBoolExp(e, new(EvalBool))
}

// Eval provides all the continuations required for matching our ADT.
// Eval also implements KindVars, so we can use it as a continuation
// for matching the result of matching an expression.
//
// Eval implements ExpBoolVars, ExpIntVars, and ExpKind. The compiler
// checks that we implement all required continuations (match cases).
// We cannot forget to implement a case. It won't compile.
type Eval struct {
	EvalInt
	EvalBool
	val interface{}
}

// eval evaluates an expression. Unlike evalBool and evalInt, which
// take an exp[t], this one takes an Exp allowing it to work on
// heterogenous expression lists. An Exp contains enough information
// to process itself. exp[t] purely exists to facilitate type-checking,
// but t is a phantom type that doesn't participate in any runtime
// decision.
//
// Because the only thing we do with the result of evaluating an
// expression is to print it, we choose to return it as an interface{},
// but we don't have to do that, we have full type-information, in
// fact we make use of it in MatchKind.
func eval(e Exp) interface{} {
	var ev Eval

	// After this call, we know that all the right decisions
	// (continuations) have been taken, but we don't know which
	// one was taken.
	ek := e.MatchExp(&ev.EvalBool, &ev.EvalInt)

	// But after this call, we know. By matching on the returned
	// kind, we can put the right result where we want it.
	ek.MatchKind(&ev)
	return ev.val
}
func (ev *Eval) Bool(ExpBoolVars) {
	// Now we know for sure the expression we just processed
	// was an exp[bool], so use the boolean denotation.
	ev.val = ev.EvalBool
}
func (ev *Eval) Int(ExpIntVars) {
	// Now we know for sure the expression we just processed
	// was an exp[int], so use the integer denotation.
	ev.val = ev.EvalInt
}

// We can create and compose expressions. All of these have concrete
// types that are all exp[t]. We can only compose expressions if they
// type-check.

var e00 = Lit(false)
var e01 = Lit(42)
var e02 = Add(e01, Lit(22))
var e03 = Or(e00, Lit(true))
var e04 = Lt(e02, e01)
var e05 = If[int](e04, e01, e02)
var e06 = If[bool](e03, e03, e00)

// We can store expressions of different type as a heterogeneous list
// of programs of type Exp because an exp[t] is an Exp.

var progs = []Exp{
	e01,
	e02,
	e03,
	e04,
	If[int](e04, e01, e02),
	If[bool](e03, e03, e00),
}

func main() {
	// we can process expressions uniformely, without having to
	// know their type.
	for _, p := range progs {
		fmt.Printf("%v=%v\n", eval(p), p)
	}
	/* this prints:
	42=(lit 42)
	64=(add (lit 42) (lit 22))
	true=(or (lit false) (lit true))
	false=(lt (add (lit 42) (lit 22)) (lit 42))
	64=(if (lt (add (lit 42) (lit 22)) (lit 42)) (lit 42) (add (lit 42) (lit 22)))
	true=(if (or (lit false) (lit true)) (or (lit false) (lit true)) (lit false))
	*/
}

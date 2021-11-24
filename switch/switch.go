package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func classify0(i int) string {
	switch {
	case i > 0:
		return "above"
	case i < 0:
		return "below"
	default:
		return "zero"
	}
}
func classify1(i int) string {
	switch a := i > 0; a {
	case true:
		return strconv.FormatBool(a) + " " + "above"
	case false:
		return strconv.FormatBool(a) + " " + "hmmm"
	default:
		return strconv.FormatBool(a) + " " + "unreachable"
	}
}

func do(classify func(int) string) {
	fmt.Println("result = ", classify(3))
	fmt.Println("result = ", classify(-32))
	fmt.Println("result = ", classify(0))
}

func ty(object interface{}) {
	fmt.Println()
	fmt.Printf("%T\n", object)
	fmt.Println("type.Kind ", reflect.TypeOf(object).Kind())
	fmt.Println("type.Name ", reflect.TypeOf(object).Name())

	fmt.Print("switch -> ")
	switch object.(type) {
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	case int8:
		fmt.Println("int8")
	case int16:
		fmt.Println("int16")
	case int32:
		fmt.Println("int32")
	case int64:
		fmt.Println("int64")
	default:
		fmt.Println("other")
	}
}

type Foo = struct{}

func main() {
	do(classify0)
	do(classify1)

	var i int = 1
	ty(i)
	var i8 int8 = 1
	ty(i8)
	var i16 int8 = 1
	ty(i16)
	var i32 int8 = 1
	ty(i32)
	var i64 int8 = 1
	ty(i64)

	ty(32i)
	ty("hi")
	ty(false)
	ty(4.5)
	ty(make(chan int))
	ty([]int{1, 2, 3})
	ty(Foo{})
}

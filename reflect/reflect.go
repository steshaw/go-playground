package main

import (
	"fmt"
	"reflect"
)

type Foo = struct {
	I   int
	Msg string
}

func main() {
	foo := Foo{
		I:   42,
		Msg: "asdf",
	}
	fmt.Printf("%+v\n", foo)

	var a1 reflect.Value = reflect.ValueOf(&foo).Elem().FieldByName("Msg")
	fmt.Printf("%+v\n", a1)

	var a interface{} = reflect.ValueOf(&foo).Elem().FieldByName("Msg").Interface()
	fmt.Printf("%+v\n", a)

	var b interface{} = reflect.Indirect(reflect.ValueOf(foo)).FieldByName("Msg")
	fmt.Printf("%+v\n", b)
}

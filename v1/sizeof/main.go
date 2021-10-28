package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func sizeof(i interface{}) {
	tyName := reflect.TypeOf(i).Name()
	fmt.Printf("f.reflect.sizeof(%s) = %d\n", tyName, reflect.TypeOf(i).Size())
	fmt.Printf("f.reflect.sizeof(%s) = %d\n", tyName, unsafe.Sizeof(i))
}

func main() {
	fmt.Println()
	var i int
	fmt.Println("  reflect.sizeof(int) =", reflect.TypeOf(i).Size())
	fmt.Println("   unsafe.sizeof(int) =", unsafe.Sizeof(i))
	sizeof(i)

	fmt.Println()
	var i8 int8
	fmt.Println("  reflect.sizeof(int8) =", reflect.TypeOf(i8).Size())
	fmt.Println("   unsafe.sizeof(int8) =", unsafe.Sizeof(i8))
	sizeof(i8)

	fmt.Println()
	var i16 int16
	fmt.Println("  reflect.sizeof(int16) =", reflect.TypeOf(i16).Size())
	fmt.Println("   unsafe.sizeof(int16) =", unsafe.Sizeof(i16))
	sizeof(i16)

	fmt.Println()
	var i32 int32
	fmt.Println("  reflect.sizeof(int32) =", reflect.TypeOf(i32).Size())
	fmt.Println("   unsafe.sizeof(int32) =", unsafe.Sizeof(i32))
	sizeof(i32)

	fmt.Println()
	var intr interface{} = 1
	fmt.Println("  reflect.sizeof(intr) =", reflect.TypeOf(intr).Size())
	fmt.Println("   unsafe.sizeof(intr) =", unsafe.Sizeof(intr))
	sizeof(intr)

	fmt.Println()
	hello := "hello"
	fmt.Println("   reflect.sizeof(hello) =", reflect.TypeOf(hello).Size())
	fmt.Println("    unsafe.sizeof(hello) =", unsafe.Sizeof(hello))
	sizeof(hello)
}

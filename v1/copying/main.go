package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type Base struct {
	Slice     []string
	StringPtr *string
	BasePtr   *Base
	Map       map[int]bool
}

func newBase() *Base {
	s := "test"
	m := make(map[int]bool)
	m[1] = true
	m[3] = true
	m[5] = true
	return &Base{
		StringPtr: &s,
		Slice:     []string{"a", "b", "c"},
		BasePtr: &Base{
			Slice: []string{"1", "2", "3"},
		},
		Map: m,
	}
}

func withCopier() {
	fmt.Println("withCopier")

	a := newBase()
	fmt.Println("original:", a)
	fmt.Println("original BasePtr:", a.BasePtr)

	var b Base
	copier.Copy(&b, &a)
	b.Slice = []string{"c", "b", "a"}
	//b.Map = make(map[int]bool)
	b.Map[3] = false
	fmt.Printf("b.StringPtr: %p\n", b.StringPtr)
	fmt.Println("b.BasePtr:", b.BasePtr)
	b.BasePtr = &Base{
		Slice: []string{"3", "2", "1"},
	}
	fmt.Println("a after updating b:", a)
	fmt.Println("BasePtr after updating b:", a.BasePtr)
}

func withBuiltinCopy() {
	fmt.Println("withBuiltinCopy")

	a := newBase()
	fmt.Println("original:", a)
	fmt.Println("original BasePtr:", a.BasePtr)

	c := a
	c.Slice = []string{"c", "b", "a"}
	c.Map[3] = false
	fmt.Println("c.StringPtr:", c.StringPtr)
	fmt.Println("c.BasePtr:", c.BasePtr)
	c.StringPtr = nil
	c.BasePtr.Slice = []string{"3", "2", "1"}

	fmt.Println("a after updating c:", a)
	fmt.Println("BasePtr after updating c:", a.BasePtr)
}

func main() {
	withCopier()
	fmt.Println()
	withBuiltinCopy()
}

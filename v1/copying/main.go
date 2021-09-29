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
	return &Base{
		StringPtr: &s,
		Slice:     []string{"a", "b", "c"},
		BasePtr: &Base{
			Slice: []string{"1", "2", "3"},
		},
		Map: map[int]bool{
			1: true,
			3: true,
			5: true,
		},
	}
}

func inspectBase(msg string, base *Base) {
	fmt.Printf("%s: %#v\n", msg, base)
	fmt.Printf("%s.BasePtr: %#v\n", msg, base.BasePtr)
}

func inspectStringSlice(msg string, slice []string) {
	fmt.Printf("%s pointer=%p len=%d cap=%d: %#v\n",
		msg, slice, len(slice), cap(slice), slice,
	)
}

func withCopier() {
	fmt.Println("withCopier")

	a := newBase()
	inspectBase("a", a)
	println()

	fmt.Println("Creating b with copier.Copy")
	var b Base
	copier.Copy(&b, a)
	inspectBase("b", &b)

	inspectStringSlice("a.Slice", a.Slice)
	inspectStringSlice("b.Slice", b.Slice)

	fmt.Println("Append d to b.Slice")
	b.Slice = append(b.Slice, "d")
	inspectStringSlice("a.Slice", a.Slice)
	inspectStringSlice("b.Slice", b.Slice)
	fmt.Printf("a.Slice.p = %p\n", a.Slice)
	fmt.Printf("b.Slice.p = %p\n", b.Slice)
	fmt.Printf("a.Slice = %#v\n", a.Slice)
	fmt.Printf("b.Slice = %#v\n", b.Slice)

	newSlice := []string{"c", "b", "a"}
	fmt.Printf("Clobbering b.Slice with a new slice %#v\n", newSlice)
	b.Slice = newSlice
	fmt.Printf("a.Slice.p = %p\n", a.Slice)
	fmt.Printf("b.Slice.p = %p\n", b.Slice)
	fmt.Printf("a.Slice = %#v\n", a.Slice)
	fmt.Printf("b.Slice = %#v\n", b.Slice)

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

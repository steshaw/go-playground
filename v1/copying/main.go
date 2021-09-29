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
	inspectStringSlice(msg+".Slice", base.Slice)
	fmt.Printf("%s.Map pointer=%p\n", msg, base.Map)
	fmt.Printf("%s.BasePtr: %#v\n", msg, base.BasePtr)
}

func inspectStringSlice(msg string, slice []string) {
	fmt.Printf("%s pointer=%p len=%d cap=%d: %#v\n",
		msg, slice, len(slice), cap(slice), slice,
	)
}

func withCopier() {
	fmt.Println("withCopier")

	println()
	fmt.Println("Creating b with copier.Copy")
	a := newBase()
	var b Base
	copier.Copy(&b, a)
	inspectBase("a", a)
	inspectBase("b", &b)

	println()
	fmt.Println("Updating b.Slice[1] to X")
	b.Slice[1] = "X"
	inspectBase("a", a)
	inspectBase("b", &b)

	println()
	fmt.Println("Append d to b.Slice")
	b.Slice = append(b.Slice, "d")
	inspectBase("a", a)
	inspectBase("b", &b)

	newSlice := []string{"c", "b", "a"}
	println()
	fmt.Printf("Clobbering b.Slice with a new slice %#v\n", newSlice)
	b.Slice = newSlice
	inspectBase("a", a)
	inspectBase("b", &b)

	println()
	fmt.Println("Updating b.Map[3] to false")
	b.Map[3] = false
	inspectBase("a", a)
	inspectBase("b", &b)

	println()
	fmt.Println("Clobber b.BasePtr")
	b.BasePtr = &Base{
		Slice: []string{"3", "2", "1"},
	}
	inspectBase("a", a)
	inspectBase("b", &b)
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

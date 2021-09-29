package main

import (
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
)

var separator = strings.Repeat("-", 80)

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

func withCopy(copy func(dest *Base, source *Base)) {
	a := newBase()
	var b Base
	copy(&b, a)
	inspectBase("a", a)
	inspectBase("b", &b)

	fmt.Println()
	fmt.Println("Updating b.Slice[1] to X")
	b.Slice[1] = "X"
	inspectBase("a", a)
	inspectBase("b", &b)

	fmt.Println()
	fmt.Println("Append d to b.Slice")
	b.Slice = append(b.Slice, "d")
	inspectBase("a", a)
	inspectBase("b", &b)

	newSlice := []string{"c", "b", "a"}
	fmt.Println()
	fmt.Printf("Clobbering b.Slice with a new slice %#v\n", newSlice)
	b.Slice = newSlice
	inspectBase("a", a)
	inspectBase("b", &b)

	fmt.Println()
	fmt.Println("Updating b.Map[3] to false")
	b.Map[3] = false
	inspectBase("a", a)
	inspectBase("b", &b)

	fmt.Println()
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
	withCopy(func(dest *Base, source *Base) {
		fmt.Println("Copying with builtin assignment")
		fmt.Println()
		*dest = *source
	})

	fmt.Println()
	fmt.Println(separator)
	withCopy(func(dest *Base, source *Base) {
		fmt.Println("Copying with copier")
		fmt.Println()
		copier.Copy(dest, source)
	})

	fmt.Println()
	fmt.Println(separator)
	withCopy(func(dest *Base, source *Base) {
		fmt.Println("Copying with copier DeepCopy=true")
		fmt.Println()
		copier.CopyWithOption(dest, source, copier.Option{DeepCopy: true})
	})
}

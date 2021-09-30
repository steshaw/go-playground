package main

import (
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
)

var separator = strings.Repeat("-", 80)

type Baz struct {
	hello string
}

type Bar struct {
	Slice  []int8
	IntPtr *int32
	BazPtr *Baz
	Map    map[int]bool
}

type Foo struct {
	Slice     []string
	StringPtr *string
	FooPtr    *Foo
	BarPtr    *Bar
	Bar       Bar
	Map       map[int]bool
}

func newFoo() *Foo {
	s := "test"
	var barInt1 int32 = 42
	var barInt2 int32 = 84
	return &Foo{
		StringPtr: &s,
		Slice:     []string{"a", "b", "c"},
		FooPtr: &Foo{
			Slice: []string{"1", "2", "3"},
		},
		BarPtr: &Bar{
			Slice:  []int8{1, 2, 3},
			IntPtr: &barInt1,
			BazPtr: &Baz{
				hello: "hi",
			},
			Map: map[int]bool{
				1: false,
				3: true,
				9: true,
			},
		},
		Bar: Bar{
			Slice:  []int8{9, 9, 9},
			IntPtr: &barInt2,
			BazPtr: &Baz{hello: "toot!"},
			Map:    map[int]bool{0: false},
		},
		Map: map[int]bool{
			1: true,
			3: true,
			5: true,
		},
	}
}

func inspectFoo(msg string, base *Foo) {
	fmt.Printf("%s: %#v\n", msg, base)
	inspectStringSlice(msg+".Slice", base.Slice)
	fmt.Printf("%s.Map pointer=%p\n", msg, base.Map)
	fmt.Printf("%s.FooPtr: %#v\n", msg, base.FooPtr)
}

func inspectStringSlice(msg string, slice []string) {
	fmt.Printf("%s pointer=%p len=%d cap=%d: %#v\n",
		msg, slice, len(slice), cap(slice), slice,
	)
}

func withCopy(copy func(dest *Foo, source *Foo)) {
	a := newFoo()
	var b Foo
	copy(&b, a)
	inspectFoo("a", a)
	inspectFoo("b", &b)

	fmt.Println()
	fmt.Println("Updating b.Slice[1] to X")
	b.Slice[1] = "X"
	inspectFoo("a", a)
	inspectFoo("b", &b)

	fmt.Println()
	fmt.Println("Append d to b.Slice")
	b.Slice = append(b.Slice, "d")
	inspectFoo("a", a)
	inspectFoo("b", &b)

	newSlice := []string{"c", "b", "a"}
	fmt.Println()
	fmt.Printf("Clobbering b.Slice with a new slice %#v\n", newSlice)
	b.Slice = newSlice
	inspectFoo("a", a)
	inspectFoo("b", &b)

	fmt.Println()
	fmt.Println("Updating b.Map[3] to false")
	b.Map[3] = false
	inspectFoo("a", a)
	inspectFoo("b", &b)

	fmt.Println()
	fmt.Println("Clobber b.FooPtr")
	b.FooPtr = &Foo{
		Slice: []string{"3", "2", "1"},
	}
	inspectFoo("a", a)
	inspectFoo("b", &b)
}

func withBuiltinCopy() {
	fmt.Println("withBuiltinCopy")

	a := newFoo()
	fmt.Println("original:", a)
	fmt.Println("original FooPtr:", a.FooPtr)

	c := a
	c.Slice = []string{"c", "b", "a"}
	c.Map[3] = false
	fmt.Println("c.StringPtr:", c.StringPtr)
	fmt.Println("c.FooPtr:", c.FooPtr)
	c.StringPtr = nil
	c.FooPtr.Slice = []string{"3", "2", "1"}

	fmt.Println("a after updating c:", a)
	fmt.Println("FooPtr after updating c:", a.FooPtr)
}

func main() {
	withCopy(func(dest *Foo, source *Foo) {
		fmt.Println("Copying with builtin assignment")
		fmt.Println()
		*dest = *source
	})

	fmt.Println()
	fmt.Println(separator)
	withCopy(func(dest *Foo, source *Foo) {
		fmt.Println("Copying with copier")
		fmt.Println()
		copier.Copy(dest, source)
	})

	fmt.Println()
	fmt.Println(separator)
	withCopy(func(dest *Foo, source *Foo) {
		fmt.Println("Copying with copier DeepCopy=true")
		fmt.Println()
		copier.CopyWithOption(dest, source, copier.Option{DeepCopy: true})
	})
}

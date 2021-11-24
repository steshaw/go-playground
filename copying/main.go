package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/kortschak/utter"
)

var separator = strings.Repeat("-", 80)

var ut *utter.ConfigState

const verbose = false
const allOperations = false

func init() {
	ut = utter.NewDefaultConfig()
	ut.CommentPointers = true
}

type Baz struct {
	hello string
}

type Bar struct {
	Slice  []int8
	Map    map[int]bool
	IntPtr *int32
	BazPtr *Baz
}

type Foo struct {
	Slice     []string
	Map       map[int]bool
	StringPtr *string
	FooPtr    *Foo
	BarPtr    *Bar
	Bar       Bar
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
	if verbose {
		fmt.Printf("%s = %s\n", msg, ut.Sdump(base))
		inspectStringSlice(msg+".Slice", base.Slice)
		fmt.Printf("%s.Map pointer=%p\n", msg, base.Map)
		fmt.Printf("%s.FooPtr: %#v\n", msg, base.FooPtr)
		fmt.Println()
	}
}

func inspectStringSlice(msg string, slice []string) {
	fmt.Printf("%s pointer=%p len=%d cap=%d: %#v\n",
		msg, slice, len(slice), cap(slice), slice,
	)
}

// pointerOf works for maps and slices (as well as pointers).
// Doesn't work for strings.
func pointerOf(a interface{}) uintptr {
	return reflect.ValueOf(a).Pointer()
}

func compare(a *Foo, b *Foo) {
	inspectFoo("a", a)
	inspectFoo("b", b)

	// Using unsafe.Pointer() on maps and slices does not work.
	//fmt.Println("Foo.Map same", unsafe.Pointer(a.Map) == unsafe.Pointer(b.Map))
	//fmt.Println("Foo.Slice same", unsafe.Pointer(a.Slice) == unsafe.Pointer(b.Slice))

	fmt.Println("Foo same", pointerOf(a) == pointerOf(b))
	fmt.Println()

	fmt.Println("Foo.Slice same", pointerOf(a.Slice) == pointerOf(b.Slice))
	fmt.Println("Foo.Map same", pointerOf(a.Map) == pointerOf(b.Map))
	fmt.Println("Foo.StringPtr same", pointerOf(a.StringPtr) == pointerOf(b.StringPtr))
	fmt.Println("Foo.FooPtr same", pointerOf(a.FooPtr) == pointerOf(b.FooPtr))
	fmt.Println("Foo.BarPtr same", pointerOf(a.BarPtr) == pointerOf(b.BarPtr))
	fmt.Println()

	fmt.Println("Foo.BarPtr.Slice same",
		pointerOf(a.BarPtr.Slice) == pointerOf(b.BarPtr.Slice),
	)
	fmt.Println("Foo.BarPtr.Map same",
		pointerOf(a.BarPtr.Map) == pointerOf(b.BarPtr.Map),
	)
	fmt.Println("Foo.BarPtr.IntPtr same",
		pointerOf(a.BarPtr.IntPtr) == pointerOf(b.BarPtr.IntPtr),
	)
	fmt.Println("Foo.BarPtr.BazPtr same",
		pointerOf(a.BarPtr.BazPtr) == pointerOf(b.BarPtr.BazPtr),
	)
	fmt.Println()

	fmt.Println("Foo.Bar.Slice same",
		pointerOf(a.Bar.Slice) == pointerOf(b.Bar.Slice),
	)
	fmt.Println("Foo.Bar.Map same",
		pointerOf(a.Bar.Map) == pointerOf(b.Bar.Map),
	)
	fmt.Println("Foo.Bar.IntPtr same",
		pointerOf(a.Bar.IntPtr) == pointerOf(b.Bar.IntPtr),
	)
	fmt.Println("Foo.Bar.BazPtr same",
		pointerOf(a.Bar.BazPtr) == pointerOf(b.Bar.BazPtr),
	)
}

func withCopy(copy func(dest *Foo, source *Foo)) {
	a := newFoo()
	var b Foo
	copy(&b, a)
	compare(a, &b)

	if allOperations {
		fmt.Println()
		fmt.Println("=== Updating b.Slice[1] to X")
		fmt.Println()
		b.Slice[1] = "X"
		compare(a, &b)

		fmt.Println()
		fmt.Println("=== Append d to b.Slice")
		fmt.Println()
		b.Slice = append(b.Slice, "d")
		compare(a, &b)

		newSlice := []string{"c", "b", "a"}
		fmt.Println()
		fmt.Printf("=== Clobbering b.Slice with a new slice %#v\n", newSlice)
		fmt.Println()
		b.Slice = newSlice
		compare(a, &b)

		fmt.Println()
		fmt.Println("=== Updating b.Map[3] to false")
		fmt.Println()
		b.Map[3] = false
		compare(a, &b)

		fmt.Println()
		fmt.Println("=== Clobber b.FooPtr")
		fmt.Println()
		b.FooPtr = &Foo{
			Slice: []string{"3", "2", "1"},
		}
		compare(a, &b)
	}
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
		fmt.Println(separator)
		fmt.Println("Copying with builtin assignment")
		fmt.Println(separator)
		fmt.Println()
		*dest = *source
	})

	fmt.Println()
	withCopy(func(dest *Foo, source *Foo) {
		fmt.Println(separator)
		fmt.Println("Copying with copier")
		fmt.Println(separator)
		fmt.Println()
		copier.Copy(dest, source)
	})

	fmt.Println()
	withCopy(func(dest *Foo, source *Foo) {
		fmt.Println(separator)
		fmt.Println("Copying with copier DeepCopy=true")
		fmt.Println(separator)
		fmt.Println()
		copier.CopyWithOption(dest, source, copier.Option{DeepCopy: true})
	})
}

package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type Base struct {
	Slice    []string
	Pointer  *string
	SPointer *Base
	Map      map[int]bool
}

func main() {
	s := "test"
	m := make(map[int]bool)
	m[1] = true
	m[3] = true
	m[5] = true
	a := Base{
		Pointer: &s,
		Slice:   []string{"a", "b", "c"},
		SPointer: &Base{
			Slice: []string{"1", "2", "3"},
		},
		Map: m,
	}
	fmt.Println("original:", a)
	fmt.Println("original SPointer:", a.SPointer)

	fmt.Println()
	fmt.Println()

	var b Base
	copier.Copy(a, &b)
	b.Slice = []string{"c", "b", "a"}
	b.Map = make(map[int]bool)
	b.Map[3] = false
	fmt.Println("b.Pointer:", b.Pointer)
	fmt.Println("b.SPointer:", b.SPointer)
	b.SPointer = &Base{
		Slice: []string{"3", "2", "1"},
	}
	fmt.Println("a after updating b:", a)
	fmt.Println("SPointer after updating b:", a.SPointer)
	fmt.Println()
	fmt.Println()

	c := a
	c.Slice = []string{"c", "b", "a"}
	c.Map[3] = false
	fmt.Println("c.Pointer:", c.Pointer)
	fmt.Println("c.SPointer:", c.SPointer)
	c.Pointer = nil
	c.SPointer.Slice = []string{"3", "2", "1"}

	fmt.Println("a after updating c:", a)
	fmt.Println("SPointer after updating c:", a.SPointer)
}

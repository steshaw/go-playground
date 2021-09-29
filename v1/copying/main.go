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

func main() {
	s := "test"
	m := make(map[int]bool)
	m[1] = true
	m[3] = true
	m[5] = true
	a := Base{
		StringPtr: &s,
		Slice:     []string{"a", "b", "c"},
		BasePtr: &Base{
			Slice: []string{"1", "2", "3"},
		},
		Map: m,
	}
	fmt.Println("original:", a)
	fmt.Println("original BasePtr:", a.BasePtr)

	fmt.Println()
	fmt.Println()

	var b Base
	copier.Copy(a, &b)
	b.Slice = []string{"c", "b", "a"}
	b.Map = make(map[int]bool)
	b.Map[3] = false
	fmt.Println("b.StringPtr:", b.StringPtr)
	fmt.Println("b.BasePtr:", b.BasePtr)
	b.BasePtr = &Base{
		Slice: []string{"3", "2", "1"},
	}
	fmt.Println("a after updating b:", a)
	fmt.Println("BasePtr after updating b:", a.BasePtr)
	fmt.Println()
	fmt.Println()

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

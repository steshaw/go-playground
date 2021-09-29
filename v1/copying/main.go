package main

import (
	"log"

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
	log.Println("original:", a)
	log.Println("original SPointer:", a.SPointer)

	log.Println()
	log.Println()

	var b Base
	copier.Copy(a, &b)
	b.Slice = []string{"c", "b", "a"}
	b.Map = make(map[int]bool)
	b.Map[3] = false
	log.Println("b.Pointer:", b.Pointer)
	log.Println("b.SPointer:", b.SPointer)
	b.SPointer = &Base{
		Slice: []string{"3", "2", "1"},
	}
	log.Println("a after updating b:", a)
	log.Println("SPointer after updating b:", a.SPointer)
	log.Println()
	log.Println()

	c := a
	c.Slice = []string{"c", "b", "a"}
	c.Map[3] = false
	log.Println("c.Pointer:", c.Pointer)
	log.Println("c.SPointer:", c.SPointer)
	c.Pointer = nil
	c.SPointer.Slice = []string{"3", "2", "1"}

	log.Println("a after updating c:", a)
	log.Println("SPointer after updating c:", a.SPointer)
}

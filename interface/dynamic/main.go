package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	Str() string
}

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.Str()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'g', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	}
	return "???"
}

type Person struct {
	first, last string
}

func (p Person) Str() string {
	return fmt.Sprintf("%#v", p)
}

func main() {
	fmt.Println(ToString(23))
	fmt.Println(ToString(float32(23.15)))
	fmt.Println(ToString(63.32))
	fmt.Println(ToString(Person{"Steven", "Shaw"}))
}

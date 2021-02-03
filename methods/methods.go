package main

import "fmt"

// Square is a square.
type square struct {
	side float64
}

func (s *square) print() {
	fmt.Printf("*square.print = %+v\n", s)
}

type circle struct {
	radius float64
}

func (c circle) print() {
	fmt.Printf("circle.print = %+v\n", c)
}
func (c *circle) printPtr() {
	fmt.Printf("*circle.print = %+v\n", c)
}

type printy interface {
	print()
}

func foo(p printy) {
	fmt.Printf("foo: ")
	p.print()
}

func main() {
	c := circle{radius: 3.0}
	cp := &c
	fmt.Println("c =", c)
	fmt.Println("cp =", cp)
	c.print()
	c.printPtr()
	(&c).print()
	(&c).printPtr()
	cp.print()
	cp.printPtr()
	(*cp).print()
	(*cp).printPtr()

	s := square{side: 1.5}
	s.print()
	(*square).print(&s)
	sp := &s
	sp.print()
	(*square).print(sp)

	foo(c)
	foo(cp)
	//	foo(s)
	foo(sp)
}

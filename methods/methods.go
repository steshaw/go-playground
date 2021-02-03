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

func (c circle) myPrint() {
	fmt.Printf("circle.print via val = %+v\n", c)
}
func (c *circle) myPrintPtr() {
	fmt.Printf("*circle.print via ptr = %+v\n", c)
}

func main() {
	c := circle{radius: 3.0}
	cp := &c
	fmt.Println("c =", c)
	fmt.Println("cp =", cp)
	c.myPrint()
	c.myPrintPtr()
	(&c).myPrint()
	(&c).myPrintPtr()
	cp.myPrint()
	cp.myPrintPtr()
	(*cp).myPrint()
	(*cp).myPrintPtr()

	s := square{side: 1.5}
	s.print()
	(*square).print(&s)
	sp := &s
	sp.print()
	(*square).print(sp)
}

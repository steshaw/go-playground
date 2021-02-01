package main

import "fmt"

type circle struct {
	radius float64
}

func (self circle) myPrint() {
	fmt.Printf("print via val = %+v\n", self)
}
func (self *circle) myPrintPtr() {
	fmt.Printf("print via ptr = %+v\n", self)
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
}

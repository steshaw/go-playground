package main

import "fmt"

type person struct {
	name string // full name
	age  int    // age in years
}

func main() {
	steshaw0 := person{"Steven Shaw", 33}
	fmt.Println("steshaw0 =", steshaw0)
	fmt.Println("steshaw0.name =", steshaw0.name)
	fmt.Println("steshaw0.age =", steshaw0.age)

	steshaw1 := person{
		name: "Steven Shaw",
		age:  33,
	}
	fmt.Println("steshaw1 =", steshaw1)

	fmt.Println("steshaw0 == steshaw1 =", steshaw0 == steshaw1)
	fmt.Println("bump age")
	steshaw1.age++
	fmt.Println("steshaw0 == steshaw1 =", steshaw0 == steshaw1)
}

// Pointers are C-like.

package main

import "fmt"

func ok(i int) {
	i++ // Has no effect outside the function.
	// Probably should get a compiler-warning here about effect-less operation.
}

// []int ~= Array Int
// *int ~= Pointer Int
func mut(i *int) {
	*i++ // Increments i at caller.
}

func main() {
	var i int = 3
	fmt.Println("i =", i)
	fmt.Println("ok(i)")
	ok(i)
	fmt.Println("i =", i)
	fmt.Println("mut(&i)")
	mut(&i)
	fmt.Println("i =", i)
}

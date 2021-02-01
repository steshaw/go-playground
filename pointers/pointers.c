// Pointers are C-like.

//package main

//import "fmt"
#include <stdio.h>

/*
func ok(i int) {
	i++ // Has no effect outside the function.
	// Probably should get a compiler-warning here about effect-less operation.
}
*/
void ok(int i) {
  i++;
}

// int[] ~= Int Array (OCaml style)
// int* ~= Int Pointer (OCaml style)
// OCaml style is unfortunate as it's the reverse of normal function
// application, making it inconsistent for dependently-typed languages.
// Go here improves on C. There have been previous proposals to adopt an
// alternative syntax for C declarations.
/*
func mut(i *int) {
	*i++ // Increments i at caller.
}
*/
void mut(int *i) {
  (*i)++; // The braces are necessary here in C.
}

/*
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
*/
int main() {
  int i = 3;
  printf("i = %d\n", i);
  printf("ok(i)");
  ok(i);
  printf("i = %d\n", i);
  printf("mut(&i)");
  mut(&i);
  printf("i = %d\n", i);
}

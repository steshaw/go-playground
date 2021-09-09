package main

import "fmt"

func main() {
	a, b := 1, 2
	// 'c' here cannot be replaced by '_' otherwise you run into "No new
	// variables on left side of :=". It turns out that "redeclaring" 'b' here
	// does not actually redeclare 'b' since it is already in scope. It reuses
	// the existing 'b' and must be the same type.
	b, c := 0, 3

	// You can instead do
	b, _c := 0, 3
	_ = _c

	// Use all the variables to avoid "declared by not used"
	fmt.Println(a, b, c)
}

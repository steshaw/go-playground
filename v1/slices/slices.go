package main

import (
	"fmt"
	"strings"
)

func slices1() {
	fmt.Println("")
	fmt.Println("Slices 1")
	fmt.Println("--------")
	names := []string{
		"Steve",
		"was",
		"here",
	}
	fmt.Println("names.length =", len(names))
	fmt.Println(names)
	names[1] = "is"
	fmt.Println(names)

	ns := make([]string, 3)
	fmt.Println("ns.length =", len(ns))
	fmt.Println(ns)
	ns[1] = names[1]
	fmt.Println(ns)
	ns[2] = names[0]
	fmt.Println(ns)
	ns[0] = names[2]
	fmt.Println(ns)
	ns[0] = strings.Title(ns[0])
	fmt.Println(ns)
	ns2 := append(ns, ".")
	fmt.Println(ns)
	fmt.Println(ns2)

	fmt.Println(strings.Join(ns, ""))
	fmt.Println(strings.Join(ns2, ""))
}
func clone(strs []string) []string {
	result := make([]string, len(strs))
	copy(result, strs)
	return result
}
func slices2() {
	fmt.Println("")
	fmt.Println("Slices 2")
	fmt.Println("--------")
	strs := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
	}

	ss2 := strs[1:]
	fmt.Println("strs =", strs)
	fmt.Println("ss2 =", ss2)
	fmt.Println("--")

	ss2[2] = "4"
	fmt.Println("strs =", strs)
	fmt.Println("ss2 =", ss2)
	fmt.Println("--")

	ss3 := clone(ss2)
	fmt.Println("ss3 =", ss3)
	fmt.Println("--")

	ss2[3] = "5"
	fmt.Println("strs =", strs)
	fmt.Println("ss2 =", ss2)
	fmt.Println("ss3 =", ss3)
	fmt.Println("--")

	ss3[4] = "6"
	ss3[5] = "7"
	fmt.Println("strs =", strs)
	fmt.Println("ss2 =", ss2)
	fmt.Println("ss3 =", ss3)
	fmt.Println("--")

	if false {
		// panic: runtime error: index out of range [6] with length 6
		ss3[6] = "Oof"
	}

	ss4 := ss3[2:]
	fmt.Println("ss3 =", ss4)
	fmt.Println("ss4 =", ss4)
	fmt.Println("--")

	ss4[1] = "5"
	fmt.Println("strs =", strs)
	fmt.Println("ss2 =", ss2)
	fmt.Println("ss3 =", ss3)
	fmt.Println("ss4 =", ss4)
	fmt.Println("--")
}
func main() {
	slices1()
	slices2()
}

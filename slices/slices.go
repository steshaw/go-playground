package main

import (
	"fmt"
	"strings"
)

func main() {
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

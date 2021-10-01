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

func slices3() {
	fmt.Println("")
	fmt.Println("Slices 3")
	fmt.Println("--------")
	src := []int8{1, 2, 3}
	fmt.Printf("src=%v\n", src)

	t := func(dst []int8) {
		dst2 := dst // Alias
		// Copy all items into dst3
		dst3 := make([]int8, len(dst))
		copy(dst3, dst)

		fmt.Println("--")
		fmt.Printf("dst=%v dst2=%v\n", dst, dst2)
		num := copy(dst, src)
		fmt.Printf("copied %d items from src=%v to dst=%v\n", num, src, dst)

		{
			result := append(dst2, src...)
			fmt.Printf("append(src=%v to dst2=%v) => %v\n", src, dst2, result)
			fmt.Printf(
				"Finally src=%v, dst=%v, dst2=%v, dst3=%v, result=%v\n",
				src, dst, dst2, dst3, result,
			)
		}

		{
			result := append(dst3, src...)
			fmt.Printf("append(src=%v to dst3=%v) => %v\n", src, dst3, result)
			fmt.Printf(
				"Finally src=%v, dst=%v, dst2=%v, dst3=%v, result=%v\n",
				src, dst, dst2, dst3, result,
			)
		}
	}

	for i := 0; i < 5; i++ {
		t(make([]int8, i))
	}

	t([]int8{95})
	t([]int8{95, 96})
	t([]int8{95, 96, 97})
	t([]int8{95, 96, 97, 98})
	t([]int8{95, 96, 97, 98, 99})
}

func main() {
	slices1()
	slices2()
	slices3()
}

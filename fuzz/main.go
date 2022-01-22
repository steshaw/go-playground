package main

import "fmt"

func reverse(s string) string {
	fmt.Printf("input: %q\n", s)
	rs := []rune(s)
	fmt.Printf("runes: %q\n", rs)
	for i, j := 0, len(rs)-1; i < len(rs)/2; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := reverse(input)
	doubleRev := reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < len(rs)/2; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs), nil
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, _ := reverse(input)
	doubleRev, _ := reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

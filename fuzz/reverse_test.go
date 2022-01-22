package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{" ", " "},
		{"Hello, world", "dlrow ,olleH"},
		{"!12345", "54321!"},
	}
	for _, tc := range cases {
		rev := reverse(tc.in)
		if rev != tc.want {
			t.Errorf("reverse: %q, want %q", rev, tc.want)
		}
	}
}

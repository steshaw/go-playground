package main

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
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
		t.Run("Hm", func(t *testing.T) {
			r := require.New(t)
			rev := reverse(tc.in)
			r.Equal(rev, tc.want)
		})
	}
}

func FuzzReverse(f *testing.F) {
	cases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range cases {
		f.Add(tc) // Seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := reverse(orig)
		doubleRev := reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

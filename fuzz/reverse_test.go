package main

import (
	"fmt"
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
			rev, err := reverse(tc.in)
			r.Nil(err)
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
		fmt.Printf("string = %s\n", orig)
		rev, err := reverse(orig)
		if err == nil {
			t.Skip()
		}
		doubleRev, err := reverse(rev)
		if err == nil {
			t.Skip()
		}
		t.Run(fmt.Sprintf("orig = '%s'", orig), func(t *testing.T) {
			r := require.New(t)
			r.Equal(orig, doubleRev)
			r.True(utf8.ValidString(orig) && !utf8.ValidString(rev), "Original is valid UTF-8, but rev is not")
		})
	})
}

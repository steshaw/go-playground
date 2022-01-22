package main

import (
	"testing"

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

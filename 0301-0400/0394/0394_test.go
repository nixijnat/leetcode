package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"abc3[cd]xyz", "abccdcdcdxyz"},
		{"abccdxyz", "abccdxyz"},
	} {
		v := decodeString(c.input)
		t.Log(c, v, v == c.expect)
	}
}

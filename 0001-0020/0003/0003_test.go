package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"aab", 2},
		{"tmmzuxt", 5},
		{"cdd", 2},
	} {
		v := lengthOfLongestSubstring(c.input)
		t.Log(c, v, v == c.expect)
	}
}

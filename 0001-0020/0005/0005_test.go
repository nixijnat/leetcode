package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"", ""},
		{"a", "a"},
	} {
		t.Log(c, longestPalindrome(c.input))
	}
}

package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 string
		input2 string
		expect bool
	}{
		{"", "", true},
		{"", "a", true},
		{"a", "", false},
		{"a", ".", true},
		{"a", "*", false},
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"abccccccd", "abc*d", true},
		{"abcccccc", "abc*d", true},
		{"abcccccce", "abc*d", false},
		{"abcccccce", "abc*", false},
		{"abcabccce", "abc*e", false},
		{"abcabcccd", "abc*e", false},
		{"abababcde", "ababababcde", true},
		{"ab", ".*c", true},
	} {
		t.Log(c, isMatch(c.input1, c.input2))
	}
}

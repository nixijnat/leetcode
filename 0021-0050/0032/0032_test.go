package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 string
		expect int
	}{
		{"(()", 2},
		{"()", 2},
		{")()())", 4},
		{")", 0},
		{"(", 0},
		{"", 0},
		{"()(()", 2},
		{"(()())", 6},
	} {
		v := longestValidParentheses(c.input1)
		t.Log(c, v)
	}
}

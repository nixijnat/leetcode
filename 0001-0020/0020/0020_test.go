package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect bool
	}{
		{"{[]}", true},
		{"([)]", false},
		{"(]", false},
		{"()[]{}", true},
		{"()", true},
	} {
		v := isValid(c.input)
		t.Log(c, v, v == c.expect)
	}
}

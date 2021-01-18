package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 int
	}{
		{1},
		{2},
		{3},
	} {

		t.Log(c, generateParenthesis(c.input1))
	}
}

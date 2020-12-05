package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect bool
	}{
		{"", false},
	} {
		t.Log(c, isValid(c.input))
	}
}

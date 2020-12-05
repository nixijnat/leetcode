package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  int
		expect bool
	}{
		{-1, false},
		{0, true},
		{8, true},
		{123, false},
		{-121, false},
		{121, true},
		{14522541, true},
	} {
		t.Log(c, isPalindrome(c.input))
	}
}

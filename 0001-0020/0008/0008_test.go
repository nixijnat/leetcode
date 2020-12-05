package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"   -  42", 0},
		{"   +  42", 0},
		{"21474836460", 2147483647},
	} {
		t.Log(c, myAtoi(c.input))
	}
}

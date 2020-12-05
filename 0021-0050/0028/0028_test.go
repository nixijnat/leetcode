package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 string
		input2 string
		expect int
	}{
		{"", "", 0},
		{"", "a", -1},
		{"a", "", 0},
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
	} {
		v := strStr(c.input1, c.input2)
		t.Log(c, v, v == c.expect)
	}
}

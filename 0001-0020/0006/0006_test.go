package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		str    string
		rows   int
		expect string
	}{
		{"LEETCODEISHIRING", 1, "LEETCODEISHIRING"},
		{"LEETCODEISHIRING", 3, "LCIRETOESIIGEDHN"},
		{"LEETCODEISHIRING", 4, "LDREOEIIECIHNTSG"},
		{"", 2, ""},
		{"a", 3, "a"},
	} {
		t.Log(c, convert(c.str, c.rows))
	}
}

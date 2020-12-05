package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 string
		input2 string
		expect string
	}{
		{"11", "1", "100"},
		{"0", "0", "0"},
		{"0", "1", "1"},
		{"1010", "1011", "10101"},
	} {
		v := addBinary(c.input1, c.input2)
		t.Log(c, v, c.expect == v)
	}
}

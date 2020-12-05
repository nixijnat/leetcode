package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 int
		input2 int
		expect int
	}{
		{10, 3, 3},
		{7, -2, -3},
		{-101, -2, 50},
		{-99, 5, -19},
		{-2147483648, -1, 2147483647},
		{-2147483648, -3, 715827882},
	} {
		v := divide(c.input1, c.input2)
		t.Log(c, v, v == c.expect)
	}
}

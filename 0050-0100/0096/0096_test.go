package main

import "testing"

// 96

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 int
		expect int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 5},
		{4, 14},
	} {
		v := numTrees(c.input1)
		t.Log(c, v, v == c.expect)
	}
}

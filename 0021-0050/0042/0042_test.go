package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		expect int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{1}, 0},
		{[]int{}, 0},
	} {
		v := trap(c.input1)
		t.Log(c, v, v == c.expect)
	}
}

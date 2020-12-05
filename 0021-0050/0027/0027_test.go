package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  []int
		dup    int
		expect int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	} {
		v := removeElement(c.input, c.dup)
		t.Log(c, v, c.input[:v], v == c.expect)
	}
}

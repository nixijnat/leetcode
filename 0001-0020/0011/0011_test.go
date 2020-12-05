package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  []int
		expect int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 1}, 2},
		{[]int{4, 3, 2, 1, 4}, 16},
	} {
		v := maxArea(c.input)
		t.Log(c, v, v == c.expect)
	}
}

package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  []int
		expect int
	}{
		{[]int{1, 2, 4, 6, 4, 5, 3, 5, 7, 8}, 6},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{1, 4, 4, 4, 2, 3, 1, 4, 4, 3, 2, 4}, 10},
		{[]int{1, 4, 4, 4, 2, 3, 1, 5, 7, 6, 8, 8}, 9},
	} {
		v := findUnsortedSubarray(c.input)
		t.Log(c, v, v == c.expect)
	}
}

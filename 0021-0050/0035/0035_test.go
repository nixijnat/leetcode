package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		input2 int
		expect int
	}{
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 6, 3},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 4, 2},
		{[]int{1, 3, 5, 6}, 3, 1},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 1, 0},
		{[]int{1, 3, 5, 6}, 0, 0},
	} {
		v := searchInsert(c.input1, c.input2)
		t.Log(c, v, v == c.expect)
	}
}

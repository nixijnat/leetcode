package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		input2 int
		expect []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{5}, 6, []int{-1, -1}},
		{[]int{5}, 5, []int{0, 0}},
		{[]int{}, 5, []int{-1, -1}},
		{[]int{0, 0, 1, 1, 1, 2, 3, 4, 4, 5, 6, 7, 7, 7, 8, 8, 8, 8, 9, 9, 10}, 4, []int{7, 8}},
	} {
		v := searchRange(c.input1, c.input2)
		t.Log(c, v)
	}
}

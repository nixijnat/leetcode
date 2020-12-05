package main

import "testing"

func TestThreeSum(t *testing.T) {
	for _, cas := range []struct {
		input []int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}}, // -4 -1 0 1 2
		{[]int{0, 0, 0, 0}},          // -4 -1 0 1 2
		{[]int{-1, 0, 1, 0}},
		{[]int{-2, 0, 1, 1, 2}},
		{[]int{-2, 0, -1, -1, 2}},
	} {
		t.Log(threeSum2(cas.input))
	}
}

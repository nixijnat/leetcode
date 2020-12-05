package main

import "testing"

func TestThreeSum(t *testing.T) {
	for _, cas := range []struct {
		input  []int
		target int
		ret    int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 2, 1, -3}, 1, 0},
	} {
		v := threeSumClosest(cas.input, cas.target)
		t.Log(cas, v, v == cas.ret)
	}
}

package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		expect int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{-3, -2, -1, -4}, -1},
		{[]int{0}, 0},
		{[]int{}, 0},
	} {
		v := maxSubArray(c.input1)
		t.Log(c, v, v == c.expect)
	}
}

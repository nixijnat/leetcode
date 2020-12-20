package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  []int
		expect int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 3, 6, 5, 0, 1, 4, 7, 1}, 7},
	} {
		v := maxProfit(c.input)
		t.Log(c, v, v == c.expect)
	}
}

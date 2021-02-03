package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  []int
		expect int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	} {
		v := maxProfit(c.input)
		t.Log(c, v, v == c.expect)
	}
}

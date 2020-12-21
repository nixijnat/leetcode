package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  []int
		k      int
		expect int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 4, 2, 1, 2, 3}, 3, 4},
	} {
		v := subarraySum(c.input, c.k)
		t.Log(c, v, v == c.expect)
	}
}

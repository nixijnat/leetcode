package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		expect int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 1},
		{[]int{1, 2, 3}, 1},
		{[]int{1, 2, 3, 4}, 1},
	} {
		v := permute(c.input1)
		t.Log(c, v, len(v))
	}
}

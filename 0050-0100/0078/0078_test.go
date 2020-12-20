package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
	}{
		{[]int{1}},
		{[]int{1, 2, 3}},
		{[]int{1, 2, 3, 4}},
	} {
		v := subsets(c.input1)
		t.Log(c, len(v), v)
	}
}

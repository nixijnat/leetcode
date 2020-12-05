package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		input2 int
		expect [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{[]int{7}, []int{2, 2, 3}}},
		{[]int{2, 3, 5}, 8, [][]int{[]int{3, 5}, []int{2, 3, 3}, []int{2, 2, 2, 2}}},
	} {
		v := combinationSum(c.input1, c.input2)
		t.Log(c, v)
	}
}

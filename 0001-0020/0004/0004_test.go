package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  [][]int
		result float64
	}{
		{[][]int{[]int{}, []int{}}, 0},
		{[][]int{[]int{1, 3}, []int{2}}, 2},
		{[][]int{[]int{1, 2}, []int{3, 4}}, 2.5},
		{[][]int{[]int{1, 3}, []int{2, 4}}, 2.5},
		{[][]int{[]int{0, 0}, []int{0, 0}}, 0},
		{[][]int{[]int{2}, []int{}}, 2},
		{[][]int{[]int{}, []int{1}}, 1},
		{[][]int{[]int{1, 3, 5, 7, 9, 11}, []int{6, 8, 10, 12, 14}}, 8},
		{[][]int{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}, []int{0, 6}}, 8},
	} {
		t.Log(c.input, c.result, findMedianSortedArraysV1(c.input[0], c.input[1]))
	}
}

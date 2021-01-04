package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 [][]int
		expect int
	}{
		{[][]int{[]int{1, 3, 1},
			[]int{1, 5, 1},
			[]int{4, 2, 1}}, 7},
		{[][]int{[]int{1, 2, 3},
			[]int{4, 5, 6}}, 12},
		{[][]int{[]int{1}}, 1},
		{[][]int{}, 0},
	} {
		v := minPathSum(c.input1)
		t.Log(c, v, v == c.expect)
	}
}

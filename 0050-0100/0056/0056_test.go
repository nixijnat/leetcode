package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 [][]int
		expect [][]int
	}{
		{[][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}},
			[][]int{[]int{1, 6}, []int{8, 10}, []int{15, 18}}},
		{[][]int{[]int{1, 4}, []int{4, 5}}, [][]int{[]int{1, 5}}},
		{[][]int{[]int{1, 5}, []int{2, 3}, []int{8, 10}, []int{4, 6}}, [][]int{[]int{1, 6}, []int{8, 10}}},
	} {
		v := merge(c.input1)
		t.Log(c, v)
	}
}

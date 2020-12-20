package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		expect int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 2, 0, 1}, 3},
	} {
		v := longestConsecutive(c.input1)
		t.Log(c, v, v == c.expect)
	}
}

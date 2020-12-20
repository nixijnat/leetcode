package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		expect int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
	} {
		v := largestRectangleArea(c.input1)
		t.Log(c, v, v == c.expect)
	}
}

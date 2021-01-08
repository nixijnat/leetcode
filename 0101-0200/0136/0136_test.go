package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		expect int
	}{
		{[]int{2, 1, 1, 3, 2}, 3},
		{[]int{2}, 2},
		{[]int{}, 0},
	} {
		v := singleNumber(c.input1)
		t.Log(c, v, v == c.expect)
	}
}

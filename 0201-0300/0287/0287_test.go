package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		expect int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
	} {
		v := findDuplicate(c.input1)
		t.Log(c, v, v == c.expect)
	}
}

package q03

import (
	"leetcode/common/has"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, cas := range []struct {
		input  []int
		expect []int
	}{
		{[]int{}, []int{-1}},
		{[]int{0, 1, 0}, []int{0}},
		{[]int{2, 3, 1, 0, 2, 5, 3}, []int{2, 3}},
		{[]int{2, 3, 1, 3, 2, 5, 0}, []int{2, 3}},
		{[]int{2, 3, 1, 1}, []int{1}},
		{[]int{2, 3, 0, 1}, []int{-1}},
	} {
		res := findRepeatNumber(cas.input)
		if !has.Int(cas.expect, res) {
			t.Fatalf("input: %v, expect: %v, but get: %v", cas.input, cas.expect, res)
		}
	}
}

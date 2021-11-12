package q24

import (
	"leetcode/common/cmp"
	"leetcode/common/list"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, cas := range []struct {
		input  []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
	} {
		ln := list.MakeList(cas.input)
		res := list.ResolveList(reverseList(ln))
		if !cmp.IsEqualIntSlice(res, cas.expect) {
			t.Fatalf("input: %v, expect: %v, but get: %v", cas.input, cas.expect, res)
		}
	}
}

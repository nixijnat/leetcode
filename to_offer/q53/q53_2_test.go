package q53

import "testing"

func TestFunc2(t *testing.T) {
	for _, cas := range []struct {
		nums   []int
		expect int
	}{
		{[]int{}, 0},
		{[]int{0}, 1},
		{[]int{0, 1}, 2},
		{[]int{0, 1, 3}, 2},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 9}, 8},
	} {
		res := missingNumber(cas.nums)
		if res != cas.expect {
			t.Fatalf("input: %v, expect: %v, but get: %v", cas.nums, cas.expect, res)
		}
	}
}

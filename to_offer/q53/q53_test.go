package q53

import "testing"

func TestFunc(t *testing.T) {
	for _, cas := range []struct {
		nums   []int
		target int
		expect int
	}{
		{[]int{}, 0, 0},
		{[]int{5, 7, 7, 8, 8, 10}, 8, 2},
		{[]int{5, 7, 7, 8, 8, 10}, 6, 0},
	} {
		res := search(cas.nums, cas.target)
		if res != cas.expect {
			t.Fatalf("input: %v, %d, expect: %v, but get: %v", cas.nums, cas.target, cas.expect, res)
		}
	}
}

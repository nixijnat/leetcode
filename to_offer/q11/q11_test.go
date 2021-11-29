package q11

import "testing"

func TestFunc(t *testing.T) {
	for _, cas := range []struct {
		input  []int
		expect int
	}{
		{[]int{}, -1},
		{[]int{3, 1, 3, 3}, 1},
		{[]int{1, 3, 3}, 1},
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
		{[]int{2, 2, 2, 2, 2}, 2},
	} {
		res := minArray(cas.input)
		if res != cas.expect {
			t.Fatalf("input: %v, expect: %v, but get: %v", cas.input, cas.expect, res)
		}
	}
}

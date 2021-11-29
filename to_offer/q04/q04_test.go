package q04

import "testing"

func TestFunc(t *testing.T) {
	array2d := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	for _, cas := range []struct {
		target int
		expect bool
	}{
		{5, true},
		{20, false},
	} {
		res := findNumberIn2DArray(array2d, cas.target)
		if res != cas.expect {
			t.Fatalf("input: %v, expect: %v, but get: %v", cas.target, cas.expect, res)
		}
	}
}

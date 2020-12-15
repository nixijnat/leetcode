package main

import (
	"reflect"
	"testing"
)

func makeMatrix(l []int, n int) [][]int {
	if len(l) == 0 {
		return nil
	}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = l[i*n : i*n+n]
	}
	return res
}

func resolveMatrix(m [][]int) []int {
	res := make([]int, 0, len(m)*len(m))
	for _, v := range m {
		res = append(res, v...)
	}
	return res
}

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		n      int
		expect []int
	}{
		{[]int{1, 2, 3,
			4, 5, 6,
			7, 8, 9}, 3,
			[]int{7, 4, 1,
				8, 5, 2,
				9, 6, 3}},
		{[]int{}, 0, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{5, 1, 9, 11,
			2, 4, 8, 10,
			13, 3, 6, 7,
			15, 14, 12, 16}, 4,
			[]int{15, 13, 2, 5,
				14, 3, 4, 1,
				12, 6, 8, 9,
				16, 7, 10, 11}},
	} {
		tmp := makeMatrix(c.input1, c.n)
		rotate(tmp)
		v := resolveMatrix(tmp)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 [][]int
		expect [][]int
	}{
		{[][]int{[]int{7, 1}, []int{7, 0}, []int{4, 4}, []int{5, 2}, []int{5, 0}, []int{6, 1}},
			[][]int{[]int{5, 0}, []int{7, 0}, []int{5, 2}, []int{6, 1}, []int{4, 4}, []int{7, 1}}},
	} {
		v := reconstructQueue(c.input1)
		t.Log(c, v, reflect.DeepEqual(c.expect, v))
	}
}

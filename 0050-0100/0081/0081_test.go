package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		k      int
		expect bool
	}{
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{1, 0, 1, 1, 1}, 0, true},
		{[]int{1, 1, 1, 1, 0, 1}, 0, true},
		{[]int{3, 5, 1}, 1, true},
		{[]int{0, 0, 1, 1, 2, 0}, 2, true},
	} {
		v := search(c.input1, c.k)
		t.Log(c, v, reflect.DeepEqual(c.expect, v))
	}
}

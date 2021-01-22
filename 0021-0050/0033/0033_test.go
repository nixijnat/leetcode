package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		k      int
		expect int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{3, 1}, 1, 1},
	} {
		v := search(c.input1, c.k)
		t.Log(c, v, reflect.DeepEqual(c.expect, v))
	}
}

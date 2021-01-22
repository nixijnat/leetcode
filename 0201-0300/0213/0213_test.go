package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		expect int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 3, 2}, 3},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
	} {
		v := rob(c.input1)
		t.Log(c, v, reflect.DeepEqual(c.expect, v))
	}
}

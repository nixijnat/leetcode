package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		input2 int
		expect int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{1, 1, 1, 1, 1}, -3, 5},
		{[]int{1}, 2, 0},
	} {
		v := findTargetSumWays(c.input1, c.input2)
		t.Log(v, reflect.DeepEqual(c.expect, v))
	}
}

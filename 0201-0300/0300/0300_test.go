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
		{[]int{1, 1, 1, 1, 1}, 1},
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
	} {
		v := lengthOfLIS(c.input1)
		t.Log(c, v, reflect.DeepEqual(c.expect, v))
	}
}

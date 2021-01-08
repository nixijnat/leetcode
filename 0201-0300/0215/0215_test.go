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
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	} {
		v := findKthLargest(c.input1, c.input2)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

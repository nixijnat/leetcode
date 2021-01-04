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
		{[]int{1}, 1},
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	} {
		v := majorityElement(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

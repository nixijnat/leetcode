package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		expect []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	} {
		sortColors(c.input1)
		t.Log(c, reflect.DeepEqual(c.input1, c.expect))
	}
}

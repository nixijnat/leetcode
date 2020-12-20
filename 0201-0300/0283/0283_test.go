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
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	} {
		moveZeroes(c.input1)
		t.Log(c, reflect.DeepEqual(c.input1, c.expect))
	}
}

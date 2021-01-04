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
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	} {
		v := productExceptSelf(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

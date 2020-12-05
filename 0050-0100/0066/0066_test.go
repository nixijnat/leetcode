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
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{0}, []int{1}},
		{[]int{9}, []int{1, 0}},
	} {
		v := plusOne(c.input1)
		t.Log(c, v, reflect.DeepEqual(c.expect, v))
	}

}

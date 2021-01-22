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
		{[]int{2, 9, 7, 3, 1}, 12},
	} {
		v := rob(c.input1)
		t.Log(c, v, reflect.DeepEqual(c.expect, v))
	}
}

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
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
	} {
		v := maxProduct(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

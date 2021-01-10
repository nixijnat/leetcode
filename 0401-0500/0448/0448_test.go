package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  []int
		expect []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
	} {
		v := findDisappearedNumbers(c.input)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

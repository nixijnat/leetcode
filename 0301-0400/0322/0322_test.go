package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 int
		input2 []int
		expect int
	}{
		{11, []int{1, 2, 5}, 3},
		{3, []int{2}, -1},
		{0, []int{1}, 0},
		{1, []int{1}, 1},
		{2, []int{1}, 2},
		{25, []int{1, 5, 7}, 5},
		{6249, []int{186, 419, 83, 408}, 20},
	} {
		v := coinChange(c.input2, c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 int
		input2 [][]int
		expect bool
	}{
		{2, [][]int{{0, 1}}, true},
		{2, [][]int{{0, 1}, {1, 0}}, false},
		{1, [][]int{}, true},
		{0, [][]int{}, false},
	} {
		v := canFinish(c.input1, c.input2)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

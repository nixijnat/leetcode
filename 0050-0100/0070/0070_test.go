package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 int
		expect int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	} {
		v := climbStairs(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

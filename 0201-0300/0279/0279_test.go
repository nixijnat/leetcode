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
		{4, 1},
		{12, 3},
		{13, 2},
	} {
		v := numSquares(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

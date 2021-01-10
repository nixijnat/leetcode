package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 int
		input2 int
		expect int
	}{
		{1, 4, 2},
		{2, 5, 3},
		{2, 7, 2},
		{7, 8, 4},
	} {
		v := hammingDistance(c.input1, c.input2)
		t.Log(v, reflect.DeepEqual(c.expect, v))
	}
}

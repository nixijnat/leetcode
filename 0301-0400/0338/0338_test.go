package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 int
		expect []int
	}{
		{7, []int{0, 1, 1, 2, 1, 2, 2, 3}},
	} {
		v := countBits(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

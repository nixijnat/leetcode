package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 [][]byte
		expect int
	}{
		{[][]byte{{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'}}, 6},
	} {
		v := maximalRectangle(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

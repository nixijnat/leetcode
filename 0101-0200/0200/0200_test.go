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
		{[][]byte{
			[]byte{'1', '1', '1', '1', '0'},
			[]byte{'1', '1', '0', '1', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '0', '0', '0'}}, 1},
		{[][]byte{
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '1', '0', '0'},
			[]byte{'0', '0', '0', '1', '1'}}, 3},
	} {
		v := numIslands(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

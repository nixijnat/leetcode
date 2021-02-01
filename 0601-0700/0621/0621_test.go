package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []byte
		input2 int
		expect int
	}{
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0, 6},
		{[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2, 16},
	} {
		v := leastInterval(c.input1, c.input2)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

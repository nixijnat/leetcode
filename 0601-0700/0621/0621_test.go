package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []byte
		k      int
		expect int
	}{
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0, 6},
		{[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2, 16},
		{[]byte{'A', 'B', 'C', 'D', 'E', 'A', 'B', 'C', 'D', 'E'}, 4, 10},
	} {
		v := leastInterval(c.input1, c.k)
		t.Log(c, v, v == c.expect)
	}
}

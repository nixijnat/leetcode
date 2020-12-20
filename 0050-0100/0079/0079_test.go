package main

import "testing"

func TestFunc(t *testing.T) {

	m := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	for _, c := range []struct {
		input1 string
		expect bool
	}{
		{"ABCCED", true},
		{"SEE", true},
		{"ABCB", false},
	} {
		v := exist(m, c.input1)
		t.Log(c, v, v == c.expect)
	}

}

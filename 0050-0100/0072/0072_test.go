package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 string
		input2 string
		expect int
	}{
		{"", "", 0},
		{"", "123", 3},
		{"123", "", 3},
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	} {
		v := minDistance(c.input1, c.input2)
		t.Log(c, v, c.expect == v)
	}

}

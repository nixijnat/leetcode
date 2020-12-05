package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  int
		expect int
	}{
		{-1, -1},
		{0, 0},
		{8, 8},
		{123, 321},
		{-7668822, -2288667},
		{1534236469, 0},
		{120, 21},
	} {
		t.Log(c, reverse(c.input))
	}
}

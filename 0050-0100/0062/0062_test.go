package main

import "testing"

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 int
		input2 int
		expect int
	}{
		{3, 7, 28},
		{7, 3, 28},
		{3, 2, 3},
		{2, 3, 3},
		{23, 12, 193536720},
		{9, 8, 6435},
		{59, 5, 557845},
	} {
		v := uniquePaths(c.input1, c.input2)
		t.Log(c, v, v == c.expect)
	}
}

func TestCombination(t *testing.T) {
	for _, c := range []struct {
		input1 int
		input2 int
		expect int
	}{
		{7, 1, 7},
		{7, 2, 21},
		{7, 3, 35},
		{7, 4, 35},
		{7, 5, 21},
		{7, 6, 7},
		{7, 7, 1},
	} {
		v := combination(c.input1, c.input2)
		t.Log(c, v, v == c.expect)
	}
}

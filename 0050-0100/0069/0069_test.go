package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 int
		expect int
	}{
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{8, 2},
		{10, 3},
		{15, 3},
		{24, 4},
		{15555555, 3944},
	} {
		v := mySqrt(c.input1)
		t.Log(c, v, c.expect == v)
	}

}

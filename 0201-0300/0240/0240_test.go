package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	matrix := [][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}
	for _, c := range []struct {
		k      int
		expect bool
	}{
		{9, true},
		{11, true},
		{15, true},
		{17, true},
		{31, false},
	} {
		v := searchMatrix(matrix, c.k)
		t.Log(c, v, reflect.DeepEqual(c.expect, v))
	}
}

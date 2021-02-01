package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		expect int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
		{[]int{}, 0},
		{[]int{1}, 0},
	} {
		v := maxProfit(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

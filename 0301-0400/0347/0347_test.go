package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  []int
		k      int
		expect []int
	}{
		{[]int{1, 2, 3, 2, 1, 1}, 2, []int{1, 2}},
		{[]int{1, 2, 3, 3, 1, 1}, 2, []int{1, 3}},
		{[]int{1}, 1, []int{1}},
		{[]int{}, 1, []int{}},
	} {
		v := topKFrequent(c.input, c.k)
		sort.Ints(v)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		k      int
		expect []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{4, 3, 11}, 3, []int{11}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{9, 11}, 2, []int{11}},
		{[]int{4, -2}, 2, []int{4}},
		{makeSlice(10000, 0), 50000, func() []int {
			res := makeSlice(10000, 5000)
			res = append(res, 5000)
			return res
		}()},
	} {
		v := maxSlidingWindow(c.input1, c.k)
		//t.Log(c, v, reflect.DeepEqual(c.expect, v))
		t.Log(reflect.DeepEqual(c.expect, v))
	}
}

func makeSlice(i int, end int) []int {
	res := make([]int, 0, i*10)
	for ; i > end; i-- {
		for j := 0; j < 10; j++ {
			res = append(res, i)
		}
	}
	return res
}

package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 [][]int
		expect int
	}{
		{[][]int{[]int{0, 30}, []int{5, 10}, []int{10, 20}}, 2},
		{[][]int{[]int{7, 10}, []int{2, 4}}, 1},
		{[][]int{}, 0},
	} {
		v := minMeetingRooms(c.input1)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

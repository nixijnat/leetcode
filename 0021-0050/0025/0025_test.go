package main

import (
	"reflect"
	"testing"
)

func makeList(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	head := &ListNode{
		Val: s[0],
	}
	cur := head
	for i := 1; i < len(s); i++ {
		cur.Next = &ListNode{Val: s[i]}
		cur = cur.Next
	}
	return head
}

func resolveList(l *ListNode) []int {
	ret := make([]int, 0, 8)
	for i := l; i != nil; i = i.Next {
		ret = append(ret, i.Val)
	}
	return ret
}

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input  []int
		k      int
		expect []int
	}{
		{[]int{}, 1, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 1, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 2, []int{2, 1, 4, 3, 6, 5, 8, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3, []int{3, 2, 1, 6, 5, 4, 7, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, []int{4, 3, 2, 1, 8, 7, 6, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 5, []int{5, 4, 3, 2, 1, 6, 7, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 6, []int{6, 5, 4, 3, 2, 1, 7, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 7, []int{7, 6, 5, 4, 3, 2, 1, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8, []int{8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 9, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	} {
		v := resolveList(reverseKGroup(makeList(c.input), c.k))
		t.Log(c.expect, v, reflect.DeepEqual(c.expect, v))
	}
}

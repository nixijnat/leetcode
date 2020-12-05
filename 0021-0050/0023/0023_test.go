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
	if l == nil {
		return nil
	}
	ret := make([]int, 0, 8)
	for i := l; i != nil; i = i.Next {
		ret = append(ret, i.Val)
	}
	return ret
}

func TestFunc(t *testing.T) {
	makeLists := func(ls [][]int) []*ListNode {
		ret := make([]*ListNode, 0, len(ls))
		for _, i := range ls {
			ret = append(ret, makeList(i))
		}
		return ret
	}
	for _, cas := range []struct {
		input  [][]int
		expect []int
	}{
		{[][]int{[]int{1, 4, 5}, []int{1, 3, 4}, []int{2, 6}},
			[]int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{[]int{1, 4, 5}, []int{}, []int{2, 6}, []int{0, 3, 10, 11, 15}},
			[]int{0, 1, 2, 3, 4, 5, 6, 10, 11, 15}},
		{[][]int{}, []int{}},
		{[][]int{[]int{}}, []int{}},
		{[][]int{[]int{1, 2, 3}, []int{4, 5, 6, 7}},
			[]int{1, 2, 3, 4, 5, 6, 7}},
	} {
		ls := makeLists(cas.input)
		v := mergeKLists(ls)
		ret := resolveList(v)
		t.Log(cas, ret, reflect.DeepEqual(ret, cas.expect))
	}
}

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
	term := 15
	for i := l; i != nil; i = i.Next {
		ret = append(ret, i.Val)
		term--
		if term == 0 {
			break
		}
	}
	return ret
}

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		expect []int
	}{
		{[]int{5, 0, 1, 8, 2, 4}, []int{4, 2, 8, 1, 0, 5}},
	} {
		v := resolveList(reverseList(makeList(c.input1)))
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}

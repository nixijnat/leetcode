package main

import "testing"

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
		input1 []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{2, 1, 3}},
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{2, 1, 4, 3, 6, 5}},
	} {
		v := swapPairs(makeList(c.input1))
		t.Log(c, resolveList(v))
	}
}

package main

import "testing"

func makeList(s []int) *ListNode {
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
		input2 []int
		expect []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 3, 5, 7, 9}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6, 7, 9}},
		{[]int{1, 3, 5}, []int{2, 4, 6, 8, 10}, []int{1, 2, 3, 4, 5, 6, 8, 10}},
	} {
		v := mergeTwoLists(makeList(c.input1), makeList(c.input2))
		t.Log(c, resolveList(v))
	}
}

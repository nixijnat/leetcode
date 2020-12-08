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

func makeCycle(s []int, pos int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	head := &ListNode{
		Val: s[0],
	}
	ptrto := (*ListNode)(nil)
	if pos == 0 {
		ptrto = head
	}
	cur := head
	for i := 1; i < len(s); i++ {
		cur.Next = &ListNode{Val: s[i]}
		cur = cur.Next
		pos--
		if pos == 0 {
			ptrto = cur
		}
	}
	cur.Next = ptrto
	return head
}

func TestFunc(t *testing.T) {
	fetch := func(p *ListNode) int {
		if p == nil {
			return -1
		}
		return p.Val
	}
	for _, c := range []struct {
		input1 []int
		pos    int
		expect int
	}{
		{[]int{3, 2, 0, -4}, 1, 2},
		{[]int{3, 2}, 0, 3},
		{[]int{0}, -1, -1},
	} {
		ll := makeCycle(c.input1, c.pos)
		v := fetch(detectCycle(ll))
		t.Log(c, v, v == c.expect, resolveList(ll))
	}
}

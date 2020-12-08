package main

import (
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

func makeJoints(i1, i2 []int) (*ListNode, *ListNode) {
	var l1, l2 *ListNode
	if len(i1) != 0 {
		l1 = makeList(i1)
	}
	if len(i2) != 0 {
		joinPos := 0
		if l1 != nil {
			i, j := len(i1)-1, len(i2)-1
			for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
				if i1[i] != i2[j] {
					break
				}
				joinPos++
			}
		}
		if joinPos == 0 {
			l2 = makeList(i2)
		} else {
			l2 = makeList(i2[:len(i2)-joinPos])
			l2Last := l2
			for ; l2Last.Next != nil; l2Last = l2Last.Next {
			}
			joinPtr := l1
			for i := len(i1) - joinPos; i > 0; i-- {
				joinPtr = joinPtr.Next
			}
			l2Last.Next = joinPtr
		}
	}
	return l1, l2
}

func fetch(i *ListNode) int {
	if i == nil {
		return -1
	}
	return i.Val
}

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input1 []int
		input2 []int
		expect int
	}{
		{[]int{5, 0, 1, 8, 2, 4}, []int{4, 0, 8, 2, 4}, 8},
		{[]int{0, 9, 1, 2, 4}, []int{3, 2, 4}, 2},
		{[]int{0, 9, 1, 2}, []int{3, 2, 4}, -1},
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, []int{2}, -1},
	} {
		l1, l2 := makeJoints(c.input1, c.input2)
		// t.Log(resolveListRaw(l1), resolveListRaw(l2))
		v := fetch(getIntersectionNode(l1, l2))
		t.Log(c, v, v == c.expect)
	}
}

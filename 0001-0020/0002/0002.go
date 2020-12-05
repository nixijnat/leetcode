package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fetch := func(p *ListNode) (int, *ListNode) {
		if p == nil {
			return 0, nil
		}
		return p.Val, p.Next
	}
	var ret, cur *ListNode
	var val, v1, v2 int
	for l1 != nil || l2 != nil {
		v1, l1 = fetch(l1)
		v2, l2 = fetch(l2)
		val = v1 + v2 + val
		node := &ListNode{}
		node.Val, val = val%10, val/10
		if cur != nil {
			cur.Next = node
		} else {
			ret = node
		}
		cur = node
	}
	if val != 0 {
		cur.Next = &ListNode{
			Val: val,
		}
	}
	return ret
}

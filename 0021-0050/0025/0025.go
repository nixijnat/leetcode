package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	reverse := func(first, last *ListNode) {
		dummy := &ListNode{0, first}
		last = last.Next
		for first.Next != last {
			cur := first.Next
			first.Next = cur.Next
			cur.Next = dummy.Next
			dummy.Next = cur
		}
	}
	dummy := &ListNode{0, head}
	pre := dummy
FINISH:
	for head != nil {
		tail := head
		for i := k - 1; i > 0; i-- {
			tail = tail.Next
			if tail == nil {
				break FINISH
			}
		}
		reverse(head, tail)
		pre.Next = tail
		pre = head
		head = head.Next
	}
	return dummy.Next
}

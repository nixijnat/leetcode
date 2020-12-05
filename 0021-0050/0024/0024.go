package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	pre, p1 := dummy, head
	for p1 != nil && p1.Next != nil {
		// 交换
		p2 := p1.Next
		pre.Next = p2
		p1.Next = p2.Next
		p2.Next = p1
		// 下一个
		pre = p1
		p1 = p1.Next
	}
	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	ptr := head.Next
	head.Next = swapPairs(head.Next.Next)
	ptr.Next = head
	return ptr
}

package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}
	num := 0
	for i := head; i != nil; i = i.Next {
		num++
	}
	if num == n {
		return head.Next
	}
	pre, cur := head, head
	for i := num - n + 1; i > 1; i-- {
		pre = cur
		cur = cur.Next
	}
	pre.Next = cur.Next
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}
	var remove func(ele *ListNode, target int) int
	remove = func(ele *ListNode, target int) int {
		if ele.Next == nil {
			return 1
		}
		pos := remove(ele.Next, target)
		if pos == 0 {
			return 0
		} else if pos == target {
			ele.Next = ele.Next.Next
			return 0
		} else {
			pos++
			return pos
		}
	}
	if remove(head, n) == n {
		return head.Next
	}
	return head
}

func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}
	stack := make([]*ListNode, 0, 8)
	dummy := &ListNode{0, head} // 哑结点
	for i := dummy; i != nil; i = i.Next {
		stack = append(stack, i)
	}
	pre := stack[len(stack)-n-1]
	pre.Next = pre.Next.Next
	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}
	dummy := &ListNode{0, head}
	pre, ptr1 := dummy, head
	for i := n; i > 0; i-- {
		ptr1 = ptr1.Next
	}
	for ; ptr1 != nil; ptr1 = ptr1.Next {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}

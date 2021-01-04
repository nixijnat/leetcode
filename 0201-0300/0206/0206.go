package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := ListNode{Next: head}
	next := head.Next
	head.Next = nil
	for cur := next; cur != nil; cur = next {
		// 把cur插入头节点
		next = cur.Next
		cur.Next = dummy.Next
		dummy.Next = cur
	}
	return dummy.Next
}
func reverseList2(head *ListNode) *ListNode {
	if head != nil {
		head, _ = reverse(head)
	}
	return head
}

// 返回新的头节点和尾节点
func reverse(head *ListNode) (*ListNode, *ListNode) {
	tail := head
	if head.Next != nil {
		var last *ListNode
		head, last = reverse(head.Next)
		last.Next = tail
	}
	tail.Next = nil
	return head, tail
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nHead := reverseList(head.Next)
	// head.Next 就是新尾节点，把head 放在 新尾节点后面
	head.Next.Next = head
	head.Next = nil // 头节点变尾节点
	return nHead
}

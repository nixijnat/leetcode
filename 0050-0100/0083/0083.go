package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for next := cur.Next; next != nil; {
		if cur.Val != next.Val {
			// 连接
			cur.Next = next
			cur = next
		}
		next = next.Next
	}
	// 这一步不要忘记了
	cur.Next = nil
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		// 使用循环跳过相同节点
		// 实际上是在找当前结点的下一个节点
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

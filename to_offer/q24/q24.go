package q24

import "leetcode/common/list"

// 24. 反转链表
func reverseList(head *list.ListNode) *list.ListNode {
	if head == nil {
		return nil
	}
	dummy := list.ListNode{
		Next: head,
	}
	n := head.Next
	head.Next = nil
	for n != nil {
		tmp := dummy.Next
		dummy.Next = n
		n = n.Next
		dummy.Next.Next = tmp
	}
	return dummy.Next
}

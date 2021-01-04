package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome1(head *ListNode) bool {
	stack := make([]*ListNode, 0, 4)
	slow := head
	for fast := head; fast != nil; {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			// 这一步可能难以理解，其实就是保证
			// fast 为尾节点时，中间元素不入栈
			stack = append(stack, slow)
		}
		slow = slow.Next
	}
	// slow 与 栈对比
	for l := len(stack) - 1; l >= 0; slow, l = slow.Next, l-1 {
		if slow.Val != stack[l].Val {
			return false
		}
	}
	return true
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// slow , fast 用于找到后半部分。
	// tmp 目前用于记录 slow 的前一个指针
	slow, tmp := head, head
	for fast := head; fast != nil; {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		tmp = slow
		slow = slow.Next
	}
	// 反转后半部分，参考 0206 反转链表
	for cur, next := slow.Next, slow; cur != nil; cur = next {
		next = cur.Next
		cur.Next = tmp.Next
		tmp.Next = cur
	}
	slow.Next = nil // slow 指向尾节点
	slow = tmp.Next // slow 恢复到 后半部分的头节点
	// 双段对比
	for tmp := head; slow != nil; slow, tmp = slow.Next, tmp.Next {
		if slow.Val != tmp.Val {
			return false
		}
	}
	return true
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	// 遍历A，获得尾节点和长度
	ptr1, len1 := headA, 1
	for ; ptr1.Next != nil; ptr1, len1 = ptr1.Next, len1+1 {
	}
	// 遍历B，获得尾节点和长度
	prt2, len2 := headB, 1
	for ; prt2.Next != nil; prt2, len2 = prt2.Next, len2+1 {
	}
	// 判断是否相交
	if ptr1 != prt2 {
		return nil
	}
	// ptr1 指向长链表 ptr2 指向短链表
	// len1 记录长度差
	if len1 > len2 {
		len1 = len1 - len2
		ptr1, prt2 = headA, headB
	} else {
		len1 = len2 - len1
		ptr1, prt2 = headB, headA
	}
	// 长链表先走 len1 步
	for ; len1 > 0; ptr1, len1 = ptr1.Next, len1-1 {
	}
	// 两个链表同时走，判断指针是否相等
	for ; ptr1 != prt2; ptr1, prt2 = ptr1.Next, prt2.Next {
	}
	return ptr1
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var ptr1, ptr2 = headA, headB
	for ptr1 != ptr2 {
		// 同时到达尾节点，且 ptr1 != ptr2 则说明不相交
		if ptr1.Next == nil && ptr2.Next == nil {
			return nil
		}
		if ptr1.Next == nil {
			ptr1 = headB // 转移
		} else {
			ptr1 = ptr1.Next
		}
		if ptr2.Next == nil {
			ptr2 = headA // 转移
		} else {
			ptr2 = ptr2.Next
		}
	}
	return ptr1
}

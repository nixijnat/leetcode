package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle1(head *ListNode) *ListNode {
	m := make(map[*ListNode]int, 4)
	for i := head; i != nil; i = i.Next {
		tmp := m[i]
		if tmp > 0 {
			return i
		}
		m[i]++
	}
	return nil // 有nil指针就是结尾
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	// fast 走得快，只需要用fast判断 nil
	// 找到相遇点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	// NOTE 判断是否有环
	if fast == nil || fast.Next == nil {
		return nil
	}
	// 同时走，判断相遇
	fast = head
	for ; fast != slow; fast, slow = fast.Next, slow.Next {
	}
	return fast
}

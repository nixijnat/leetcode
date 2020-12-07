package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle1(head *ListNode) bool {
	m := make(map[*ListNode]int, 4)
	for i := head; i != nil; i = i.Next {
		tmp := m[i]
		if tmp > 0 {
			return true
		}
		m[i]++
	}
	return false // 有nil指针就是结尾
}

func hasCycle2(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	// fast 走得快，只需要用fast判断 nil
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

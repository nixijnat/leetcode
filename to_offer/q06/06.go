package q06

import "leetcode/common/list"

// 06. 从尾到头打印链表

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reversePrint(head *list.ListNode) []int {
	var res []int
	for n := head; n != nil; n = n.Next {
		res = append(res, n.Val)
	}
	reverseSlice(res)
	return res
}

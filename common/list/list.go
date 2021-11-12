package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	head := &ListNode{
		Val: s[0],
	}
	cur := head
	for i := 1; i < len(s); i++ {
		cur.Next = &ListNode{Val: s[i]}
		cur = cur.Next
	}
	return head
}

func ResolveList(l *ListNode) []int {
	if l == nil {
		return []int{}
	}
	ret := make([]int, 0, 8)
	for i := l; i != nil; i = i.Next {
		ret = append(ret, i.Val)
	}
	return ret
}

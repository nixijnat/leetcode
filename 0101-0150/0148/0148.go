package main

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList_select(head *ListNode) *ListNode {
	// 选择排序
	dummy := &ListNode{0, head}
	// cur 表示下一个节点插在它后面
	for cur := dummy; cur.Next != nil; cur = cur.Next {
		pre := cur // pre 表示 它的下一个节点是最小值
		for tmp := cur.Next; tmp.Next != nil; tmp = tmp.Next {
			if tmp.Next.Val < pre.Next.Val {
				pre = tmp
			}
		}
		// 移除 picked
		picked := pre.Next
		pre.Next = picked.Next
		// 将 picked 插入在cur后面
		picked.Next = cur.Next
		cur.Next = picked
	}
	return dummy.Next
}

func sortList_insert(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 插入排序
	dummy := &ListNode{0, head}
	cur := head.Next // cur 当前被插入节点
	head.Next = nil  // NOTE 表示已经排好了
	for next := (*ListNode)(nil); cur != nil; cur = next {
		next = cur.Next // 记录下一个被插入点
		// 从头开始找合适的位置：用 pre 好一些，避免适用两个指针
		pre := dummy
		for ; pre.Next != nil && cur.Val > pre.Next.Val; pre = pre.Next {
		}
		// 插入到 pre 后面
		cur.Next = pre.Next
		pre.Next = cur
	}
	return dummy.Next
}

func sortList_quick(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 快速排序
	// 写入数组
	arr := make([]*ListNode, 0, 8)
	for i := head; i != nil; i = i.Next {
		arr = append(arr, i)
	}
	// 排序
	sort.Slice(arr, func(i, j int) bool { return arr[i].Val < arr[j].Val })
	// 重组链表
	i := len(arr) - 1
	arr[i].Next = nil // NOTE 很重要，链表的结尾
	for i--; i >= 0; i-- {
		arr[i].Next = arr[i+1]
	}
	return arr[0]
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 计算长度
	length := 0
	for i := head; i != nil; i, length = i.Next, length+1 {
	}
	// 归并排序
	dummy := &ListNode{0, head}
	// 使用步长作外层循环，每次加倍步长
	for step := 1; step < length; step = step << 1 {
		// pre 表示当前处理段的前一个节点
		// next 表示下一个处理段的前一个节点，也就是当前处理段的最后一个节点
		pre, next := dummy, head
		// remainLen 表示剩余长度
		for remainLen := length; remainLen > 0; remainLen -= 2 * step {
			// 归并两小段的长度：用长度记录何时结束
			ln, rn := step, step
			if ln+rn > remainLen {
				rn = remainLen - ln
				if rn <= 0 { // 只有一段了不需要归并
					continue
				}
			}
			// 归并子段
			pre.Next, next = merge(pre.Next, ln, rn)
			pre = next
		}
	}
	return dummy.Next
}

// merge 参数
// l: 首节点
// ln, rn： 左、右子段的长度
// 返回值：子段的头节点和尾节点
func merge(l *ListNode, ln, rn int) (*ListNode, *ListNode) {
	// 寻找右子段的首节点
	r := l
	for i := ln; i > 0; r, i = r.Next, i-1 {
	}
	// 寻找和记录下一段的首节点，避免无法连接起来
	next := r
	for i := rn; i > 0; next, i = next.Next, i-1 {
	}
	// 归并的主体：参考 0021 合并两个升序链表
	dummy := &ListNode{0, l}
	cur := dummy
	// 通过计数进行终点判断
	for ; ln > 0 && rn > 0; cur = cur.Next {
		if l.Val <= r.Val {
			cur.Next = l
			l = l.Next
			ln--
		} else {
			cur.Next = r
			r = r.Next
			rn--
		}
	}
	if ln > 0 {
		cur.Next = l
	}
	if rn > 0 {
		cur.Next = r
		ln = rn
	}
	for ; ln > 0; cur, ln = cur.Next, ln-1 {
	}
	// cur 就是尾节点
	// 和下一段连起来
	cur.Next = next
	return dummy.Next, cur
}

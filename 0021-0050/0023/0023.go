package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (h *localHeap) print() {
	fmt.Println("----------------")
	for i, j := 1, 1; i < len(h.data); {
		end := i + j
		if end >= len(h.data) {
			end = len(h.data)
		}
		fmt.Print(strings.Repeat(" ", (len(h.data)-j)/2))
		for k := i; k < end; k++ {
			fmt.Print(h.data[k].Val, " ")
		}
		fmt.Println("")
		i = end
		j *= 2
	}
}

type localHeap struct {
	data []*ListNode
}

func (h *localHeap) pop() *ListNode {
	l := len(h.data)
	if l <= 1 {
		return nil
	}
	ret := h.data[1]
	h.data[0] = h.data[l-1]
	h.data = h.data[:l-1]
	if l == 2 {
		return ret
	}
	l = l - 1 // 调整长度
	i := 1
	for c := 2; c < l; c = i * 2 {
		next := 0
		if h.data[c].Val < h.data[0].Val {
			next = c
		}
		if c+1 < l && h.data[c+1].Val < h.data[0].Val &&
			h.data[c+1].Val < h.data[c].Val {
			next = c + 1
		}
		if next == 0 {
			break
		}
		h.data[i] = h.data[next]
		i = next
	}
	h.data[i] = h.data[0]
	return ret
}

func (h *localHeap) push(v *ListNode) {
	h.data = append(h.data, v)
	h.data[0] = v
	i := len(h.data) - 1
	for p := i / 2; p > 0 && h.data[p].Val > v.Val; p, i = p/2, p {
		h.data[i] = h.data[p]
	}
	h.data[i] = v
}

func mergeKLists1(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	h := localHeap{
		data: make([]*ListNode, 1, 8), // 1 是哑结点
	}
	// 初始化，将各链表首节点push到堆
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		h.push(lists[i])
	}
	cur := dummy
	for {
		cur.Next = h.pop()   // pop最小值
		if cur.Next == nil { // 如果为空则结束
			break
		}
		cur = cur.Next // 当前尾节点
		// 将pop节点的下一个节点尝试push进去
		if cur.Next == nil {
			continue
		}
		h.push(cur.Next)
	}
	cur.Next = nil // 将尾节点置空
	return dummy.Next
}

/// 官方库

type myHeap struct {
	data []*ListNode
}

func (m *myHeap) Len() int {
	return len(m.data)
}

func (m *myHeap) Swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *myHeap) Less(i, j int) bool {
	return m.data[i].Val < m.data[j].Val
}

func (m *myHeap) Pop() interface{} {
	l := len(m.data)
	ret := m.data[l-1]
	m.data = m.data[:l-1]
	return ret
}

func (m *myHeap) Push(x interface{}) {
	m.data = append(m.data, x.(*ListNode))
}

func mergeKLists2(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	h := &myHeap{
		data: make([]*ListNode, 0, 8),
	}
	// 初始化，将各链表首节点push到堆
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		h.data = append(h.data, lists[i])
	}
	if len(h.data) == 0 {
		return nil
	}
	heap.Init(h)
	cur := dummy
	for h.Len() > 0 { // 要用长度判断是否还有元素，不能使用自定义heap的Pop空判断
		cur.Next = heap.Pop(h).(*ListNode) // pop最小值
		cur = cur.Next                     // 当前尾节点
		// 将pop节点的下一个节点尝试push进去
		if cur.Next == nil {
			continue
		}
		heap.Push(h, cur.Next)
	}
	cur.Next = nil // 将尾节点置空
	return dummy.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
	if l1 == nil {
		cur.Next = l2
	}
	if l2 == nil {
		cur.Next = l1
	}
	return head.Next
}

// 归并一段
func merge(lists []*ListNode, low, high int) *ListNode {
	switch {
	case low == high:
		return lists[low]
	case low > high:
		return nil
	default:
		mid := (high + low) >> 1
		// 分治归并
		return mergeTwoLists(merge(lists, low, mid), merge(lists, mid+1, high))
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	// 分治归并
	return merge(lists, 0, l-1)
}

package main

import (
	"container/list"
)

type LRUCache_1 struct {
	m   map[int]*list.Element // 快速查询节点
	l   *list.List            // 双向链表 使得能够快速移除元素
	cap int
}

func newLRUCache(cap int) LRUCache_1 {
	return LRUCache_1{
		m:   make(map[int]*list.Element, cap),
		l:   list.New(),
		cap: cap,
	}
}

func Constructor(capacity int) LRUCache {
	return newLRUCache2(capacity)
}

func (c *LRUCache_1) Put(key, value int) {
	ele := c.m[key]
	// 已存在 则更新节点 并且放在链表头部
	if ele != nil {
		ele.Value = [2]int{key, value}
		c.l.MoveToFront(ele)
		return
	}
	// 不存在
	// 先校验长度，有必要则移除尾部元素
	if len(c.m) == c.cap {
		tail := c.l.Back()
		delete(c.m, tail.Value.([2]int)[0])
		c.l.Remove(tail)
	}
	// put 新元素到链表头部
	c.m[key] = c.l.PushFront([2]int{key, value})
}

func (c *LRUCache_1) Get(key int) int {
	ele := c.m[key]
	if ele == nil {
		return -1
	}
	// 访问则放在链表头部
	c.l.MoveToFront(ele)
	return ele.Value.([2]int)[1]
}

///
type myNode struct {
	val  interface{} // 存储 key 和 value
	prev *myNode
	next *myNode
}

type myList struct {
	head *myNode
	tail *myNode
}

func newList() *myList {
	// 使用一个哑结点
	dummy := &myNode{}
	dummy.next = dummy
	dummy.prev = dummy
	return &myList{
		head: dummy,
		tail: dummy,
	}
}

func (m *myList) PushFront(val interface{}) *myNode {
	n := &myNode{
		val: val,
	}
	m.insertFront(n)
	return n
}

func (m *myList) insertFront(n *myNode) {
	n.next = m.head.next
	n.prev = m.head
	m.head.next = n
	n.next.prev = n
}

func (m *myList) empty() bool {
	return m.head == m.head.next
}

func (m *myList) Back() *myNode {
	if m.empty() {
		return nil
	}
	return m.tail.prev
}

func (m *myList) MoveToFront(n *myNode) {
	if m.head.next == n {
		return
	}
	// remove
	m.Remove(n)
	// insert
	m.insertFront(n)
}

func (m *myList) Remove(n *myNode) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

type LRUCache struct {
	m   map[int]*myNode
	l   *myList
	cap int
}

func newLRUCache2(cap int) LRUCache {
	return LRUCache{
		m:   make(map[int]*myNode, cap),
		l:   newList(),
		cap: cap,
	}
}
func (c *LRUCache) Put(key, value int) {
	ele := c.m[key]
	if ele != nil {
		ele.val = [2]int{key, value}
		c.l.MoveToFront(ele)
		return
	}
	if len(c.m) == c.cap {
		tail := c.l.Back()
		delete(c.m, tail.val.([2]int)[0])
		c.l.Remove(tail)
	}
	c.m[key] = c.l.PushFront([2]int{key, value})
}

func (c *LRUCache) Get(key int) int {
	ele := c.m[key]
	if ele == nil {
		return -1
	}
	c.l.MoveToFront(ele)
	return ele.val.([2]int)[1]
}

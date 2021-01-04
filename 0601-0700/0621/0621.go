package main

import (
	"container/heap"
	"sort"
)

// 任务调度
type cons []int

func (c cons) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c cons) Len() int {
	return len(c)
}
func (c cons) Less(i, j int) bool {
	return c[i] < c[j]
}
func (c cons) Pop() interface{} {
	l := len(c)
	res := c[l-1]
	c = c[:l-1]
	return res
}
func (c cons) Push(i interface{}) {
	c = append(c, i.(int))
}
func leastInterval(tasks []byte, n int) int {
	l := len(tasks)
	if n == 0 {
		return l
	}
	m := make(map[byte]int, l/2+1)
	for i := range tasks {
		m[tasks[i]]++
	}
	list := make([]int, 0, len(m))
	for _, v := range m {
		list = append(list, v)
	}
	sort.Slice(list, func(i, j int) bool { return list[i] > list[j] })
	h := make(cons, n+1)
	heap.Init(h)
	for _, v := range list {
		h[0] += v
		heap.Fix(h, 0)
	}
	sort.Slice(h, func(i, j int) bool { return h[i] > h[j] })
	res := (h[0]-1)*(n+1) + 1
	for i := 1; i < n && h[i] == h[i-1]; i++ {
		res++
	}
	return res
}

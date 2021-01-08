package main

import (
	"container/heap"
)

func maxSlidingWindow_slow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	// 构建堆
	h := make(heap1, k)
	copy(h, nums[:k])
	heap.Init(&h)
	// 第一个结果
	res := make([]int, 0, len(nums))
	res = append(res, h[0])
	for i := k; i < len(nums); i++ {
		// 找到窗口最左端的元素，然后去掉
		// 这个时间复杂度高
		if r := find(h, nums[i-k]); r >= 0 {
			heap.Remove(&h, r)
		}
		// 添加新的值
		heap.Push(&h, nums[i])
		res = append(res, h[0])
	}
	return res
}

func find(h heap1, k int) int {
	for i := len(h) - 1; i >= 0; i-- {
		if h[i] == k {
			return i
		}
	}
	return -1
}

type heap1 []int

func (h *heap1) Len() int           { return len(*h) }
func (h *heap1) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *heap1) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heap1) Pop() interface{} {
	res := (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	return res
}
func (h *heap1) Push(v interface{}) {
	(*h) = append((*h), v.(int))
}

//---------------------------
func maxSlidingWindow_heap(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	// 构建堆
	h := make(heap2, 0, len(nums))
	for i, v := range nums[:k] {
		heap.Push(&h, [2]int{v, i})
	}
	// 第一个结果
	res := make([]int, 0, len(nums))
	res = append(res, h[0][0])
	for i := k; i < len(nums); i++ {
		// 添加新的值
		heap.Push(&h, [2]int{nums[i], i})
		// 若堆顶元素不再窗口中，则去除
		for h.Len() > 0 && h[0][1] <= i-k {
			heap.Pop(&h)
		}
		res = append(res, h[0][0])
	}
	return res
}

// 堆中存储二元组：值和索引
type heap2 [][2]int

func (h *heap2) Len() int           { return len(*h) }
func (h *heap2) Less(i, j int) bool { return (*h)[i][0] > (*h)[j][0] }
func (h *heap2) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heap2) Pop() interface{} {
	res := (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	return res
}
func (h *heap2) Push(v interface{}) {
	(*h) = append((*h), v.([2]int))
}

// 单调队列法

func maxSlidingWindow_queue(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	queue := make([]int, 0, k)
	res := make([]int, 0, k)
	for i, v := range nums {
		// 保持单调队列
		// push的元素保持单调
		for ql := len(queue); ql > 0 && nums[queue[ql-1]] < v; ql = len(queue) {
			queue = queue[:ql-1]
		}
		queue = append(queue, i)
		// 移除非窗口元素
		if queue[0] <= i-k {
			queue = queue[1:]
		}
		if i >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}
func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	prefixs := make([]int, length) // 前缀最大值
	suffixs := make([]int, length) // 后缀最大值
	for i := range nums {
		// 计算各个块的前缀最大值
		if i%k == 0 {
			prefixs[i] = nums[i]
		} else {
			prefixs[i] = max(prefixs[i-1], nums[i])
		}
		// 计算各个块的后缀最大值
		// 最后一个块的不一定长度为 k，所以 i == 0 时要注意
		if i == 0 || (length-i)%k == 0 {
			suffixs[length-1-i] = nums[length-1-i]
		} else {
			suffixs[length-1-i] = max(suffixs[length-i], nums[length-1-i])
		}
	}
	res := make([]int, 0, length)
	for i := 0; i < length-k+1; i++ {
		res = append(res, max(suffixs[i], prefixs[i+k-1]))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

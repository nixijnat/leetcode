package main

import (
	"container/heap"
	"math/rand"
	"time"
)

// 随机轴优化版
func findKthLargest1(nums []int, k int) int {
	rand.Seed(time.Now().Unix()) // 随机化种子
	low, high := 0, len(nums)-1
	k-- // 第K大的数，换算为 index 应该减一
	for low <= high {
		// 从大到小排序
		mid := partition(nums, low, high)
		switch {
		case mid == k: // 第 k个元素已经排好位置
			return nums[k]
		case mid < k:
			low = mid + 1
		default:
			high = mid - 1
		}
	}
	return 0
}

func partition(nums []int, low, high int) int {
	// 随机选择目标元素
	r := rand.Intn(high-low+1) + low
	v := nums[r] // 真实 的 low
	nums[r] = nums[low]
	for low < high {
		for ; low < high && nums[high] <= v; high-- {
		}
		nums[low] = nums[high]
		for ; low < high && nums[low] >= v; low++ {
		}
		nums[high] = nums[low]
	}
	nums[low] = v
	return low
}

// 构建堆：小根堆
type topK struct {
	data []int
}

func (t *topK) Less(i, j int) bool { return t.data[i] < t.data[j] }
func (t *topK) Swap(i, j int)      { t.data[i], t.data[j] = t.data[j], t.data[i] }
func (t *topK) Len() int           { return len(t.data) }
func (t *topK) Pop() interface{} {
	res := t.data[t.Len()-1]
	t.data = t.data[:t.Len()-1]
	return res
}
func (t *topK) Push(v interface{}) {
	t.data = append(t.data, v.(int))
}

func findKthLargest(nums []int, k int) int {
	// 初始化堆 k 大小的小根堆
	top := &topK{
		data: nums[:k],
	}
	heap.Init(top)
	for i := k; i < len(nums); i++ {
		// 与堆顶元素比较。小于则丢弃
		// 大于则丢弃堆顶元素，并入堆新元素
		if nums[i] <= top.data[0] {
			continue
		}
		heap.Pop(top)
		heap.Push(top, nums[i])
	}
	// 返回堆顶元素
	return top.data[0]
}

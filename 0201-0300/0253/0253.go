package main

import "container/heap"

func minMeetingRooms_scan(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 找最大值和最小值
	min, max := intervals[0][0], intervals[0][1]
	for i := len(intervals) - 1; i > 0; i-- {
		if min > intervals[i][0] {
			min = intervals[i][0]
		}
		if max < intervals[i][1] {
			max = intervals[i][1]
		}
	}
	res := 0
	// 扫描所有时间段，找重合最多的情况
	for ; min <= max; min++ {
		tmp := 0
		for i := len(intervals) - 1; i >= 0; i-- {
			if min >= intervals[i][0] && min < intervals[i][1] {
				tmp++
			}
		}
		if res < tmp {
			res = tmp
		}
	}
	return res
}

// 尚未开始的会议队列
type notStartQ struct {
	data [][]int
}

func (n *notStartQ) Less(i, j int) bool { return n.data[i][0] < n.data[j][0] }
func (n *notStartQ) Swap(i, j int)      { n.data[i], n.data[j] = n.data[j], n.data[i] }
func (n *notStartQ) Len() int           { return len(n.data) }
func (n *notStartQ) Pop() interface{} {
	res := n.data[n.Len()-1]
	n.data = n.data[:n.Len()-1]
	return res
}
func (n *notStartQ) Push(v interface{}) {
	n.data = append(n.data, v.([]int))
}

// 已经开始的会议队列
type startedQ struct {
	notStartQ
}

func (s *startedQ) Less(i, j int) bool { return s.data[i][1] < s.data[j][1] }

func minMeetingRooms(intervals [][]int) int {
	// 构建优先队列
	notStart := &notStartQ{
		data: intervals,
	}
	heap.Init(notStart)
	started := &startedQ{}
	res := 0
	// 开会模拟
	for notStart.Len() > 0 {
		v := heap.Pop(notStart).([]int) // 找到下一个开始的会议
		for started.Len() > 0 {         // 从已开始会议中找 可以 结束的会议
			front := started.data[0]
			// 会议的 end 和 v 的 start 比较
			// end > start 会议不能结束
			if front[1] > v[0] {
				break
			}
			heap.Pop(started) // 结束会议
		}
		// 开始会议
		heap.Push(started, v)
		// 统计会议室数量
		if res < started.Len() {
			res = started.Len()
		}
	}
	return res
}

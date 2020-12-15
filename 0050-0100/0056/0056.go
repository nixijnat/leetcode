package main

import "sort"

func merge(intervals [][]int) [][]int {
	l := len(intervals)
	if l == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, l)
	for i := 0; i < l; {
		merged := []int{intervals[i][0], intervals[i][1]}
		try := i + 1
		for ; try < l; try++ {
			if merged[1] < intervals[try][0] { // 独立
				break
			}
			if merged[1] < intervals[try][1] { // 融合
				merged[1] = intervals[try][1]
			}
			// 吞并
		}
		i = try // 下次 从已经无法融合或并的地方开始
		res = append(res, merged)
	}
	return res
}

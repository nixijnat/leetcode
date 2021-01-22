package main

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	// 个子高的放在前面
	// 如果一样高，k 数小的放在前面
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	res := make([][]int, 0, len(people))
	for _, v := range people {
		// v[0] 一定比之前的数都小，
		// 所以直接将它放在 v[1]索引 的 位置
		if v[1] >= len(res) {
			res = append(res, v)
		} else {
			// 插入
			res = res[0 : len(res)+1]
			copy(res[v[1]+1:], res[v[1]:])
			res[v[1]] = v
		}
	}
	return res
}

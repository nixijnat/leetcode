package main

import "sort"

func longestConsecutive1(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	// 先排序
	sort.Ints(nums)
	res, tmp := 1, 1
	for i := 1; i < l; i++ {
		switch nums[i] - nums[i-1] {
		case 0: // 忽略掉相同元素
		case 1: // 递增1则加1
			tmp++
			if tmp > res {
				res = tmp
			}
		default: // 否则归零重新计算
			tmp = 1
		}
	}
	return res
}

func longestConsecutive(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	m := make(map[int]bool, l)
	for _, v := range nums {
		m[v] = true
	}
	// 查询函数, step +1 / -1 指引向上/向下
	find := func(i, step int) int {
		n := 0
		for ; m[i]; i = i + step {
			n++
			delete(m, i)
		}
		return n
	}
	res := 1
	for _, v := range nums {
		if m[v] {
			continue
		}
		// 向上 / 向下寻找
		tmp := find(v, -1) + find(v+1, 1)
		if tmp > res {
			res = tmp
		}
	}
	return res
}

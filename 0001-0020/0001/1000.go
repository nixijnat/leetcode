package main

import "sort"

func twoSumSort(nums []int, target int) []int {
	type pair struct {
		i int
		v int
	}
	var vs = make([]pair, 0, len(nums))
	for i := range nums {
		vs = append(vs, pair{i, nums[i]})
	}
	sort.Slice(vs, func(i, j int) bool {
		return vs[i].v < vs[j].v
	})
	var low, high int = 0, len(nums) - 1
	for low < high {
		tmp := vs[low].v + vs[high].v
		if tmp == target {
			return []int{vs[low].i, vs[high].i}
		} else if tmp > target {
			high--
		} else {
			low++
		}
	}
	return []int{0, 0}
}

func twoSumMap(nums []int, target int) []int {
	var m map[int]int
	var useM bool
	var tmp = len(nums)
	if tmp > 20 {
		m = make(map[int]int, tmp)
		for i := range nums {
			m[nums[i]] = i
		}
		useM = true
	}
	for i := range nums {
		tmp = target - nums[i]
		if useM {
			j, ok := m[tmp]
			if ok {
				return []int{j, i}
			}
			m[nums[i]] = i
		} else {
			for j := range nums {
				if i != j && tmp == nums[j] {
					return []int{i, j}
				}
			}
		}
	}
	return nil
}

func twoSumMap2(nums []int, target int) []int {
	var m map[int]int
	var useM bool
	var tmp = len(nums)
	if tmp > 20 {
		m = make(map[int]int, tmp)
		useM = true
	}
	for i := range nums {
		tmp = target - nums[i]
		if useM {
			j, ok := m[tmp]
			if ok {
				return []int{j, i}
			}
			m[nums[i]] = i
		} else {
			for j := range nums {
				if i != j && tmp == nums[j] {
					return []int{i, j}
				}
			}
		}
	}
	return nil
}

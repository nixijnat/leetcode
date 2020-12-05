package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	if length <= 0 {
		return 0
	}
	sort.Ints(nums)
	ret, incPos := 0, 1<<31-1
	for i := 0; i < length-2; i++ {
		for j, k := i+1, length-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			tmp := sum - target
			if tmp > 0 {
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			} else {
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				tmp = -tmp
			}
			if tmp < incPos {
				ret = sum
				incPos = tmp
			}
		}
	}
	return ret
}

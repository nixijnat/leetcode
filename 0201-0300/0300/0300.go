package main

func lengthOfLIS(nums []int) int {
	buf := make([]int, len(nums))
	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		max := 0
		// 在 i+1 到 n-1 中间找比 i 大 且 buf 最大的元素
		for j := i + 1; j < len(nums); j++ {
			// 找 比 i
			if nums[j] > nums[i] && max < buf[j] {
				max = buf[j]
			}
		}
		// 加 1
		buf[i] = max + 1
		if res < buf[i] {
			res = buf[i]
		}
	}
	return res
}

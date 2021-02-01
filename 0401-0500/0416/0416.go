package main

func canPartition(nums []int) bool {
	target := 0
	for _, v := range nums {
		target += v
	}
	// 总和为奇数，则不存在
	if target&0x1 != 0 {
		return false
	}
	// 求 HALF
	return helper(nums, target>>1)
}

func helper(nums []int, target int) bool {
	if target == 0 {
		return true
	}
	// 余留负数 或 无法分解 则失败
	if target < 0 || len(nums) == 0 {
		return false
	}
	// 使用 nums[0] 参与 分解
	if helper(nums[1:], target-nums[0]) {
		return true
	}
	// 不使用 nums[0]
	if helper(nums[1:], target) {
		return true
	}
	return false
}

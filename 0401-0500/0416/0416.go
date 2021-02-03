package main

func canPartition1(nums []int) bool {
	target := 0
	for _, v := range nums {
		target += v
	}
	// 总和为奇数，则不存在
	if target&0x1 != 0 {
		return false
	}
	// 求 HALF
	return helper1(nums, target>>1)
}

func helper1(nums []int, target int) bool {
	if target == 0 {
		return true
	}
	// 余留负数 或 无法分解 则失败
	if target < 0 || len(nums) == 0 {
		return false
	}
	// 使用 nums[0] 参与 分解
	if helper1(nums[1:], target-nums[0]) {
		return true
	}
	// 不使用 nums[0]
	if helper1(nums[1:], target) {
		return true
	}
	return false
}

// 动态规划
func canPartition(nums []int) bool {
	target := 0
	for _, v := range nums {
		target += v
	}
	if target&0x1 != 0 {
		return false
	}
	target = target >> 1
	// 使用line 和 buf 两个数组存储可组成的数字集合
	// 避免单个数组在计算时产生干扰
	line := make([]bool, target+1)
	buf := make([]bool, target+1)
	for _, v := range nums {
		if v > target { // 过大的数字
			continue
		}
		for i := range line {
			if !line[i] { // 未访问 或 不可达的数字
				continue
			}
			tmp := v + i
			// 不在目标范围内
			if tmp < 0 || tmp > target {
				continue
			}
			buf[tmp] = true // 可以组成的数字
		}
		buf[v] = true // 单独v，不和其他数字组合
		if buf[target] {
			return true
		}
		copy(line, buf) // 要bopy 而不是 交换
	}
	return false
}

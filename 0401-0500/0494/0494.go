package main

// 原生方法
// 效率低下，重复计算
func findTargetSumWays1(nums []int, S int) int {
	return find(nums, S, 0)
}

func find(nums []int, S, cur int) int {
	l := len(nums)
	if l == 0 {
		if S == cur {
			return 1
		} else {
			return 0
		}
	}
	return find(nums[1:], S, cur+nums[0]) + find(nums[1:], S, cur-nums[0])
}

// DP
func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// S 大于总和，则不可能
	if S > sum || -S > sum {
		return 0
	}
	max := 2 * sum
	buf := make([]int, max+1)
	buf[sum+nums[0]] = 1
	buf[sum-nums[0]] += 1 // 必须是 +=， nums[0]==0时，0为2
	for i := 1; i < len(nums); i++ {
		tmp := make([]int, 2*sum+1)
		for j := 0; j <= max; j++ {
			// buf[j] 为0，表示之前不存在这种情况
			// 也避免判断 j-nums[i] 和 j+nums[i] 是否在合法范围
			if buf[j] > 0 {
				tmp[j-nums[i]] += buf[j]
				tmp[j+nums[i]] += buf[j]
			}
		}
		buf = tmp
	}
	return buf[S+sum]
}

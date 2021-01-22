package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	last1, last2 := 0, 0
	// 计算不要最后一间房时的最大收益
	for i := 0; i < len(nums)-1; i++ {
		tmp := max(nums[i]+last1, last2)
		last1 = last2
		last2 = tmp
	}
	res1 := last2
	// 计算不要第一间房的最大收益
	last1, last2 = 0, 0
	for i := 1; i < len(nums); i++ {
		tmp := max(nums[i]+last1, last2)
		last1 = last2
		last2 = tmp
	}
	// 取两者的较大值
	return max(res1, last2)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

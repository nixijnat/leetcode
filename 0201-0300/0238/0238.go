package main

func productExceptSelf1(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	before := make([]int, n) // 记录 前缀乘积
	after := make([]int, n)  // 记录 后缀乘积
	before[0], after[n-1] = 1, 1
	for i := 1; i < n; i++ {
		before[i] = before[i-1] * nums[i-1]
		after[n-1-i] = after[n-i] * nums[n-i]
	}
	for i := 0; i < n; i++ {
		after[i] *= before[i] // 前缀后缀乘积
	}
	return after
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	res := make([]int, n) // 先记录 后缀乘积
	res[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		res[i] = res[i+1] * nums[i+1]
	}
	prefix := 1 // 前缀乘积
	for i := 1; i < n; i++ {
		prefix *= nums[i-1]
		res[i] *= prefix // 前缀后缀乘积
	}
	return res
}

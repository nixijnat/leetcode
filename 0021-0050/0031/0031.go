package main

func nextPermutation(nums []int) {
	l := len(nums)
	if l == 0 {
		return
	}
	i1 := l - 2
	// 第一次扫描：找下降臂的高一位：比拐点小一点的数
	for ; i1 >= 0 && nums[i1] >= nums[i1+1]; i1-- {
	}
	if i1 >= 0 {
		// 能够找到高一位则继续在下降臂找替换点：这个数比高一位大一点
		i2 := l - 1
		for ; i2 > i1 && nums[i2] <= nums[i1]; i2-- {
		}
		nums[i1], nums[i2] = nums[i2], nums[i1]
	}
	// 翻转下降臂
	for low, high := i1+1, l-1; low < high; low, high = low+1, high-1 {
		nums[low], nums[high] = nums[high], nums[low]
	}
}

package main

// 0046
func permute1(nums []int) [][]int {
	// 全排列数量是固定值，使用这个固定值作为终止条件
	// 计算总量 n * (n-1) * (n-2) * ... * 2 * 1
	n := 1
	for i := len(nums); i > 1; i-- {
		n *= i
	}
	res := make([][]int, 0, n)
	for i := n; i > 0; i-- {
		nums = next(nums)
		res = append(res, nums)
	}
	return res
}

// 下一个更大的数
func next(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	high := len(res) - 2
	for ; high >= 0 && res[high] >= res[high+1]; high-- {
	}
	if high >= 0 {
		low := len(res) - 1
		for ; res[low] <= res[high]; low-- {
		}
		res[high], res[low] = res[low], res[high]
	}
	// 翻转
	for l, r := high+1, len(res)-1; l < r; {
		res[r], res[l] = res[l], res[r]
		l++
		r--
	}
	return res
}

// 构造法
func permute(nums []int) [][]int {
	length := len(nums)
	res := make([][]int, 0, length)
	res = append(res, []int{nums[0]})
	for i := 1; i < length; i++ {
		curNum := len(res)
		for j := 1; j <= i; j++ {
			// 重复使用 res[:curNum]
			// 在 第 j 个位置插入 nums[j]
			for _, v := range res[:curNum] {
				tmp := make([]int, j, length)
				copy(tmp, v[:j])
				tmp = append(tmp, nums[i])
				tmp = append(tmp, v[j:]...)
				res = append(res, tmp)
			}
		}
		// 直接在 res[:curNum] 上修改，头部插入 nums[]
		// 避免重复分配
		for j := 0; j < curNum; j++ {
			res[j] = append(res[j][:1], res[j]...)
			res[j][0] = nums[i]
		}
	}
	return res
}

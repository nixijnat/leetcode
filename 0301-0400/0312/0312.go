package main

// map 版本
func maxCoins1(nums []int) int {
	// 扩展数组，避免负索引
	nums2 := make([]int, 0, len(nums)+2)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, nums...)
	nums2 = append(nums2, 1)
	// map[[2]int] int 作为 buf，记忆化搜索
	return find1(nums2, 0, len(nums2)-1, make(map[[2]int]int, len(nums2)*len(nums2)))
}
func find1(nums []int, low, high int, buf map[[2]int]int) int {
	if low == high {
		return 0
	}
	key := [2]int{low, high}
	if v, ok := buf[key]; ok {
		return v
	}
	res := 0
	// 计算最大值
	for i := low + 1; i < high; i++ {
		tmp := nums[low]*nums[high]*nums[i] +
			find1(nums, low, i, buf) + find1(nums, i, high, buf)
		if res < tmp {
			res = tmp
		}
	}
	// 存储
	buf[key] = res
	return res
}

// 优化版本
func maxCoins2(nums []int) int {
	// 扩展数组，避免负索引
	nums2 := make([]int, 0, len(nums)+2)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, nums...)
	nums2 = append(nums2, 1)
	// [][]int 作为 buf，记忆化搜索
	buf := make([][]int, len(nums2))
	for i := range buf {
		buf[i] = make([]int, len(nums2))
		for j := range buf[i] {
			buf[i][j] = -1
		}
	}
	return find(nums2, 0, len(nums2)-1, buf)
}
func find(nums []int, low, high int, buf [][]int) int {
	if low == high {
		return 0
	}
	if buf[low][high] != -1 {
		return buf[low][high]
	}
	res := 0
	// 计算最大值
	for i := low + 1; i < high; i++ {
		tmp := nums[low]*nums[high]*nums[i] +
			find(nums, low, i, buf) + find(nums, i, high, buf)
		if res < tmp {
			res = tmp
		}
	}
	// 存储
	buf[low][high] = res
	return res
}

// 自下而上构造法
func maxCoins(nums []int) int {
	// 扩展数组，避免负索引
	nums2 := make([]int, 0, len(nums)+2)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, nums...)
	nums2 = append(nums2, 1)
	// map[[2]int] int 作为 buf，记忆化搜索
	buf := make([][]int, len(nums2))
	for i := range buf {
		buf[i] = make([]int, len(nums2))
	}

	// high 为上界 low 为 下界
	// 假设只有一个数字，即从 -1 0 1 开始， 首先计算 buf[0][2]
	// 然后增加一个数字，low 应该从 high-2 开始降低，便于计算包含新数字 buf[low][high]
	// 若 low 从 0 开始增加，则很多 buf[k][high] 是未计算过的
	// 因此 k 也应该从 high-1 下降
	for high := 2; high < len(nums2); high++ {
		for low := high - 2; low >= 0; low-- {
			// k 为分界点，变化范围 low < k < high
			for k := high - 1; k > low; k-- {
				tmp := nums2[low]*nums2[high]*nums2[k] + buf[low][k] + buf[k][high]
				if buf[low][high] < tmp {
					buf[low][high] = tmp
				}
			}
		}
	}
	return buf[0][len(nums2)-1]
}

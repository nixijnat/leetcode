package main

func findDuplicate1(nums []int) int {
	for _, v := range nums {
		if v < 0 { // 把 v 修正为正数
			v = -v
		}
		if nums[v] < 0 { // nums[v] 已经为负数则 v重复
			return v
		}
		nums[v] = -nums[v] // 标记负数
	}
	return 0
}

func findDuplicate(nums []int) int {
	l := len(nums)
	res := 0
	for low, high := 0, l-1; low <= high; {
		mid := low + (high-low)>>1
		cnt := 0
		// 统计小于当前i的数量次数
		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			// counts[i] <= i 表示低位
			low = mid + 1
		} else {
			// counts[i] > i 表示高位，并 i 可能是结果
			res = mid
			high = mid - 1
		}
	}
	return res
}

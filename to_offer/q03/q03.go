package q03

// 03. 数组中重复的数字

func findRepeatNumberMap(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, n := range nums {
		if m[n] {
			return n
		}
		m[n] = true
	}
	return -1
}

// flag
func findRepeatNumber(nums []int) int {
	zero := false
	for _, n := range nums {
		if n == 0 {
			if zero {
				return 0
			}
			zero = true
		}
		if n < 0 {
			n = -n
		}
		if nums[n] < 0 {
			return n
		}
		nums[n] = -nums[n]
	}
	return -1
}

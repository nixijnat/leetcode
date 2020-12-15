package main

func canJump_slow(nums []int) bool {
	n := len(nums)
	buf := make([]bool, n)
	buf[n-1] = true
	for i := n - 2; i >= 0; i-- {
		for j := 1; j <= nums[i]; j++ {
			if buf[i+j] {
				buf[i] = true
				break
			}
		}
	}
	return buf[0]
}
func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	border := nums[0]
	// border 为势力范围边界
	for i := 0; i <= border; i++ {
		tmp := i + nums[i]
		if tmp > border { // 扩展势力范围
			border = tmp
		}
		if border >= n-1 { // 覆盖了终点
			return true
		}
	}
	return false
}

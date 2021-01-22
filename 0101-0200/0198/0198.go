package main

// 1982
func rob(nums []int) int {
	// last1 是前前一个房屋的收益
	// last1 是前一个房屋的收益
	last1, last2 := 0, 0
	for _, v := range nums {
		tmp := max(last1+v, last2)
		last1 = last2
		last2 = tmp
	}
	return last2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

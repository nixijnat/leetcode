package main

import "math"

// 错误解法
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buf := make([]int, len(prices))
	return helper(prices, buf, 0)
}

func helper(prices, buf []int, low int) int {
	if low >= len(prices) {
		return 0
	}
	if buf[low] != 0 {
		return buf[low]
	}
	res := math.MinInt64
	for j := low + 1; j < len(prices); j++ {
		tmp := prices[j] - prices[low] + helper(prices, buf, j+2)
		if tmp > res {
			res = tmp
		}
	}
	buf[low] = res
	return res
}

package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	// 初始状态
	holdOn, sell, frozen := -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		// 状态转移
		holdOn, sell, frozen = max(holdOn, frozen-prices[i]),
			holdOn+prices[i], max(sell, frozen)
	}
	return max(holdOn, max(sell, frozen))
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

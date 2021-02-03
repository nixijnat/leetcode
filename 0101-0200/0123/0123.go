package main

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	buy1, sell1, buy2, sell2 := -prices[0], 0, -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1, sell1, buy2, sell2 = max(-prices[i], buy1),
			max(sell1, buy1+prices[i]),
			max(buy2, sell1-prices[i]),
			max(sell2, buy2+prices[i])
	}
	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

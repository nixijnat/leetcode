package main

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	last, ret := prices[0], 0
	for _, v := range prices {
		tmp := v - last
		if tmp > 0 { // 升值时的收益
			ret += tmp
		}
		last = v
	}
	return ret
}

package main

// 低效的代码
func maxProfit_low(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	min, max, ret := prices[0], prices[0], 0
	for i := 1; i < l; i++ {
		// 有更小的值时，更新最大值和最小值
		if prices[i] < min {
			min = prices[i]
			max = prices[i]
			continue
		}
		// 更新最大值
		if prices[i] > max {
			max = prices[i]
		}
		// 求差
		if tmp := max - min; ret < tmp {
			ret = tmp
		}
	}
	return ret
}

// 优化1
func maxProfit_opt1(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	min, max, ret := prices[0], prices[0], 0
	for i := 1; i < l; i++ {
		// 有更小的值时，更新最大值和最小值
		if prices[i] < min {
			// 更新max和min前求一次
			if tmp := max - min; ret < tmp {
				ret = tmp
			}
			min = prices[i]
			max = prices[i]
			continue
		}
		// 更新最大值
		if prices[i] > max {
			max = prices[i]
		}
	}
	// 最终求一次
	if tmp := max - min; ret < tmp {
		ret = tmp
	}
	return ret
}

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	min, ret := prices[0], 0
	for _, v := range prices {
		// 有更小的值时
		if v < min {
			min = v
			continue
		}
		// 求当天收益，并求最大收益
		if tmp := v - min; ret < tmp {
			ret = tmp
		}
	}
	return ret
}

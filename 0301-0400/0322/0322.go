package main

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	// 将结果存储
	buf := make([]int, amount+1)
	return helper(coins, buf, amount)
}

func helper(coins, buf []int, amount int) int {
	if amount == 0 {
		return 0
	}
	// 已有结果
	if buf[amount] != 0 {
		return buf[amount]
	}
	// res 为一个较大的数
	// 但是不能为 amount，怕 amount 个 1
	res := amount + 1
	for _, v := range coins {
		if amount < v {
			continue
		}
		// 子问题
		tmp := helper(coins, buf, amount-v)
		if tmp != -1 && tmp < res {
			res = tmp
		}
	}
	if res == amount+1 {
		res = -1
	} else {
		res++
	}
	buf[amount] = res
	return res
}

func coinChange_2(coins []int, amount int) int {
	buf := make([]int, amount+1)
	// 从 1 开始 构造
	for i := 1; i <= amount; i++ {
		buf[i] = i + 1 // 预设一个较大的数
		for _, v := range coins {
			if i-v < 0 {
				continue
			}
			if buf[i-v] != -1 && buf[i-v] < buf[i] {
				buf[i] = buf[i-v]
			}
		}
		// 预设的数没变 说明不存在构造的方式
		if buf[i] == i+1 {
			buf[i] = -1
		} else {
			buf[i]++
		}
	}
	// 构造的最终结果
	return buf[amount]
}

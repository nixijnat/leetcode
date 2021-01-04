package main

func maxProduct1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	positive, nagetive := 0, 0 // 0 表示没有负值/正值
	res := nums[0]
	for _, v := range nums {
		switch {
		case v > 0:
			if positive == 0 {
				positive = v // 初始化正值
			} else {
				positive *= v // 变大
			}
			nagetive *= v // 变小
		case v < 0:
			tmp := nagetive
			if positive != 0 {
				nagetive = positive * v // 翻转为负值
			} else {
				nagetive = v // 初始化负值
			}
			positive = tmp * v // 翻转为正值
		default:
			positive, nagetive = 0, 0 // 归零
			if res < 0 {              // 与已有结果判断
				res = 0
			}
			continue
		}
		if positive != 0 && positive > res {
			res = positive
		}
		// 判断 res和负值 也是必要的，全负值时需要
		if nagetive != 0 && nagetive > res {
			res = nagetive
		}
	}
	return res
}

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	min, max := 1, 1
	res := nums[0]
	for _, v := range nums {
		min1 := min * v
		max1 := max * v
		min = minf(min1, minf(max1, v))
		max = maxf(min1, maxf(max1, v))
		res = maxf(max, res)
	}
	return res
}

func minf(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func maxf(a, b int) int {
	if a < b {
		return b
	}
	return a
}

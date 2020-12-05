package main

func loopSub(dividend int, divisor int) int {
	candicates := make([]int, 0, 8) // 缓存减数
	candicates = append(candicates, divisor)
	index, count := 0, 0
	for dividend >= divisor {
		// 找合适的减数
		for ; index >= 0 && dividend < candicates[index]; index-- {
		}
		// 减 和 加
		dividend -= candicates[index]
		count += 1 << index
		// 增加一个 double 减数：只有上升期增加
		if len(candicates)-1 == index {
			candicates = append(candicates, candicates[index]+candicates[index])
			index++
		}
	}
	return count
}

func divide(dividend int, divisor int) int {
	pos := true
	count := 0
	if dividend == 0 || divisor == 1 {
		count = dividend
		goto END
	}
	if divisor == -1 {
		count = -dividend
		goto END
	}
	// 符号处理
	if dividend < 0 {
		dividend = -dividend
		pos = !pos
	}
	if divisor < 0 {
		divisor = -divisor
		pos = !pos
	}
	// 辗转相减
	count = loopSub(dividend, divisor)
	if !pos {
		count = -count
	}
END:
	// 结果处理
	const max = 1 << 31
	if count < -max || count > max-1 {
		return max - 1
	}
	return count
}

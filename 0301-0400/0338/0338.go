package main

func countBits_direct(num int) []int {
	counter := func(num int) int {
		count := 0
		for ; num != 0; num >>= 1 { // 右移
			if num&0x1 != 0 { // 检验最低位是否为 1
				count++
			}
		}
		return count
	}
	res := make([]int, 0, num+1)
	for i := 0; i <= num; i++ {
		res = append(res, counter(i))
	}
	return res
}
func countBits_optim(num int) []int {
	counter := func(num int) int {
		count := 0
		for ; num != 0; num &= num - 1 { // i & (i-1) 去除最右的 1
			count++
		}
		return count
	}
	res := make([]int, 0, num+1)
	for i := 0; i <= num; i++ {
		res = append(res, counter(i))
	}
	return res
}

func countBits_topbit(num int) []int {
	res := make([]int, 1, num+1)
	step, nextStep := 0, 1 // 当前步长，与下一个步长
	for i := 1; i <= num; i++ {
		if i == nextStep { // 更新步长
			step = nextStep
			nextStep <<= 1
		}
		res = append(res, res[i-step]+1)
	}
	return res
}

func countBits(num int) []int {
	res := make([]int, 1, num+1)
	for i := 1; i <= num; i++ {
		res = append(res, res[i&(i-1)]+1)
	}
	return res
}

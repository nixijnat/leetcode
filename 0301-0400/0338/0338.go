package main

func countBits1(num int) []int {
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
func countBits(num int) []int {
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

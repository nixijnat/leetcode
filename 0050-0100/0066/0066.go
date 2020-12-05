package main

func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return []int{1}
	}
	carry := 1
	for i := l - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] < 10 { // 没有进位直接返回
			return digits
		}
		carry = digits[i] / 10
		digits[i] %= 10
	}
	// 一定有进位，则把进位放到最高位
	digits = append(digits, carry)
	copy(digits[1:], digits)
	digits[0] = carry
	return digits
}

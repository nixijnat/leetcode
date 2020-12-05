package main

import "strconv"

func addBinary1(a string, b string) string {
	l1, l2 := len(a), len(b)
	switch {
	case l1 == 0:
		return b
	case l2 == 0:
		return a
	}
	// 通过 fetch 规避 长短不一的问题
	fetch := func(s string, i int) byte {
		if i < 0 {
			return 0
		}
		return s[i] - '0'
	}
	carry := byte(0)
	res := make([]byte, 0, 10)
	// 倒着遍历
	for l1, l2 = l1-1, l2-1; l1 >= 0 || l2 >= 0; l1, l2 = l1-1, l2-1 {
		carry = fetch(a, l1) + fetch(b, l2) + carry
		tmp := carry & 0x1         // 取余 % 2
		carry = carry >> 1         // 除法 / 2
		res = append(res, tmp+'0') // append 更方便扩容
	}
	// 不要遗忘最后一个
	if carry != 0 {
		res = append(res, (carry&0x1)+'0')
	}
	// 将 结果翻转
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func addBinary(a string, b string) string {
	a1, _ := strconv.ParseInt(a, 2, 64)
	b1, _ := strconv.ParseInt(b, 2, 64)
	return strconv.FormatInt(a1+b1, 2)
}

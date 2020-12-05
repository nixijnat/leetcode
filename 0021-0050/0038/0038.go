package main

import "strconv"

func countAndSay1(n int) string {
	if n == 1 {
		return "1"
	}
	// 可以用库代码 strconv.Itoa
	parse := func(i int) []byte {
		ret := make([]byte, 0, 4)
		for ; i != 0; i = i / 10 {
			ret = append(ret, byte(i%10)+'0')
		}
		return ret
	}
	// 缓冲区，避免反复分配空间
	ret := make([]byte, 1, 10)
	tmp := make([]byte, 0, 10)
	ret[0] = '1'
	for ; n > 1; n-- {
		l := len(ret)
		char := ret[0] // 记录 基准字符
		count := 1     // 次数
		for i := 1; i < l; i++ {
			if char == ret[i] { // 相同则增加
				count++
			} else { // 不同则say
				tmp = append(tmp, parse(count)...)
				tmp = append(tmp, char)
				char = ret[i]
				count = 1
			}
		}
		// 最后一个字符
		tmp = append(tmp, parse(count)...)
		tmp = append(tmp, char)
		ret, tmp = tmp, ret[:0] // 反转
	}
	return string(ret)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	// 可以用库代码 strconv.Itoa
	// 不额外申请空间
	appendCount := func(buf []byte, i int) []byte {
		for ; i != 0; i = i / 10 {
			buf = append(buf, byte(i%10)+'0')
		}
		return buf
	}
	// 缓冲区，避免反复分配空间
	ret := make([]byte, 1, 10)
	tmp := make([]byte, 0, 10)
	ret[0] = '1'
	for ; n > 1; n-- {
		l := len(ret)
		count := 0 // 次数
		for i := 0; i < l; i++ {
			count++
			if i == l-1 || ret[i] != ret[i+1] { // 不记录起始点
				tmp = appendCount(tmp, count) // 直接append 到 缓冲区
				tmp = append(tmp, ret[i])
				count = 0
			}
		}
		ret, tmp = tmp, ret[:0] // 反转
	}
	return string(ret)
}

func countAndSay3(n int) string {
	cur := "1"
	for ; n > 1; n-- {
		l := len(cur)
		p1, p2 := 0, 1 // 双指针
		tmp := ""
		for ; p2 <= l; p2++ {
			if p2 == l || cur[p1] != cur[p2] {
				// 字符串拼接：潜在反复分配空间风险
				tmp += strconv.Itoa(p2-p1) + cur[p1:p1+1]
				p1 = p2
			}
		}
		cur = tmp
	}
	return cur
}

package main

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 || length == 1 {
		return length
	} else if length == 2 {
		if s[0] == s[1] {
			return 1
		} else {
			return 2
		}
	}
	m := make([]int, 128) // 字符：索引
	ret := 0
	for head, tail := 0, 1; tail < length; tail++ {
		next := m[s[tail]]
		if next > head || (next == head && s[tail] == s[head]) { // 起始元素特殊处理
			tmp := tail - head
			if tmp > ret {
				ret = tmp
			}
			head = next + 1
			if ret >= length-head { // 后续长度已经无法更大
				break
			}
		}
		m[s[tail]] = tail
	}
	return ret
}

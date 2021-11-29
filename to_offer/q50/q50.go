package q50

// 50. 第一个只出现一次的字符
func firstUniqChar(s string) byte {
	const charNum = 26
	var dict [charNum]int
	for i := range s {
		temp := int(s[i] - 'a')
		switch {
		case dict[temp] == 0:
			dict[temp] = i + 1
		case dict[temp] > 0:
			dict[temp] = -dict[temp]
		}
	}
	min := len(s)
	for _, c := range dict {
		c--
		if c >= 0 && c < min {
			min = c
		}
	}
	if min == len(s) {
		return ' '
	}
	return s[min]
}

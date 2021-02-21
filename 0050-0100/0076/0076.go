package main

// 不在乎次数
func minWindow_nocount(s, t string) string {
	if s == "" {
		return ""
	}
	// 不在乎次数，则不统计次数
	freq := make(map[byte]struct{}, len(t))
	for i := range t {
		freq[t[i]] = struct{}{}
	}
	sl, tl := len(s), len(freq)
	ret := ""
	matched := make(map[byte]int, tl)
	for lp, rp := 0, 0; ; {
		// 右指针移动，直到找到第一个符合的
		for ; rp < sl && len(matched) != tl; rp++ {
			if _, ok := freq[s[rp]]; ok {
				matched[s[rp]]++
			}
		}
		if len(matched) != tl { // 只比较字符类型数量
			break
		}
		if ret == "" || len(ret) > rp-lp {
			ret = s[lp:rp]
		}
		// 左指针移动，找所有符合条件的结果，直到让子序列不符合要求
		// exit 用于标记退出，因为左指针要跳过 无关字符
		for exit := false; lp < rp; lp++ {
			val, ok := matched[s[lp]]
			if !ok {
				continue
			}
			if exit {
				break
			}
			if val > 1 { // 频次较多，仍然符合要求
				matched[s[lp]] = val - 1
				continue
			}
			delete(matched, s[lp]) // 移除字符，让子序列不符合要求
			exit = true
		}
	}
	return ret
}

// 在乎次数
func minWindow(s, t string) string {
	if s == "" {
		return ""
	}
	freq := make(map[byte]int, len(t)) // 频次
	for i := range t {
		freq[t[i]]++
	}
	sl, fl := len(s), len(freq)
	ret := ""                         // 结果
	matched := make(map[byte]int, fl) // 已匹配的情况
	unmatchCnt := fl                  // 未匹配的字母数量
	for lp, rp := 0, 0; ; {
		// 右指针移动，直到找到第一个符合的
		for ; rp < sl && unmatchCnt != 0; rp++ {
			count, ok := freq[s[rp]] // 忽略掉不匹配的元素
			if !ok {
				continue
			}
			matched[s[rp]]++             // 增加频次
			if matched[s[rp]] == count { // 频次到达要求
				unmatchCnt--
			}
		}
		if unmatchCnt != 0 { // 没找到就直接退出
			break
		}
		// 左指针移动，找所有符合条件的结果，直到让子序列不符合要求
		// exit 用于标记退出，因为左指针要跳过 无关字符
		for exit := false; lp < rp; lp++ {
			count, ok := freq[s[lp]] // 忽略掉不匹配的元素
			if !ok {
				continue
			}
			if exit { // 退出
				break
			}
			// 找结果
			if ret == "" || len(ret) > rp-lp {
				ret = s[lp:rp]
			}
			matched[s[lp]]--            // 降低频次
			if matched[s[lp]] < count { // 不再符合要求
				exit = true
				unmatchCnt++
			}
			// 降低频次后仍然符合要求
		}
	}
	return ret
}
func minWindow_new(s, t string) string {
	if s == "" || len(s) < len(t) {
		return ""
	}
	// 统计字符频次
	var freq [128]int
	for i := range t {
		freq[t[i]]++
	}
	// 记录未匹配的字符数
	var unmatchCnt = len(t)
	var count [128]int
	var res string
	for l, r := 0, 0; r < len(s); r++ {
		// 右指针扩张寻找目标字串
		val := s[r]
		// 忽略无关字符
		if freq[val] == 0 {
			continue
		}
		// 更新未匹配次数
		if count[val] < freq[val] {
			unmatchCnt--
		}
		count[val]++ // 记录实际字符频次
		if unmatchCnt != 0 {
			continue
		}
		// 左指针收缩，直到不满足要求
		// 注意：佐治
		for ; ; l++ {
			// 更新结果
			if res == "" || len(res) > len(s[l:r+1]) {
				res = s[l : r+1]
			}
			val := s[l]
			// 忽略无关字符
			if freq[val] == 0 {
				continue
			}
			count[val]--
			// 频次不满足要求，则退出
			if count[val] < freq[val] {
				unmatchCnt++
				// 注意，退出前将左指针指向 相关字符
				for l++; l < r && freq[s[l]] == 0; l++ {
				}
				break
			}
		}
	}
	return res
}

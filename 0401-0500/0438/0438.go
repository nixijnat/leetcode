package main

func findAnagrams_slow(s, p string) []int {
	// 统计目标频次
	freq := [26]byte{}
	for i := range p {
		freq[p[i]-'a']++
	}
	res := make([]int, 0, 4)
	windowLen := len(p) - 1
	// 滑动窗口
	for left, right := 0, windowLen; right < len(s); right = left + windowLen {
		var cnt [26]byte
		// 统计窗口内部频次
		for ; right >= left; right-- {
			val := s[right] - 'a'
			if freq[val] == 0 { // 无关字符，提前中断
				break
			}
			cnt[val]++
		}
		// 是否提前中断
		if right >= left {
			left = right + 1
			continue
		}
		// 是否字母异位词
		if cnt == freq {
			res = append(res, left)
		}
		left++
	}
	return res
}

func findAnagrams_normal(s, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	// 统计目标频次
	freq := [26]int{}
	for i := range p {
		freq[p[i]-'a']++
	}
	res := make([]int, 0, 4)
	var cnt [26]int
	// 滑动窗口
	for left, right := 0, 0; right < len(s); {
		// 找字母异位词
		for ; right < len(s) && right-left < len(p); right++ {
			val := s[right] - 'a'
			cnt[val]++
			// 如果 cnt 频次比 freq 高，则 left 需要收缩
			// 直到 cnt 频次和 freq 一样
			// 这里包含了 s[right] 为 无关字符的情况
			for ; cnt[val] > freq[val]; left++ {
				cnt[s[left]-'a']--
			}
		}
		// 如果 right-left < len(p) 说明没有找到 字母异位词
		if right-left < len(p) {
			continue
		}
		res = append(res, left)
		// left 向右收缩一次
		cnt[s[left]-'a']--
		left++
	}
	return res
}
func findAnagrams(s, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	// 统计目标频次
	freq := [26]int{}
	for i := range p {
		freq[p[i]-'a']++
	}
	res := make([]int, 0, 4)
	var cnt [26]int
	// 滑动窗口
	for left, right := 0, 0; right < len(s); right++ {
		// 找字母异位词
		val := s[right] - 'a'
		cnt[val]++
		// 如果 cnt 频次比 freq 高，则 left 需要收缩
		// 直到 cnt 频次和 freq 一样
		// 这里包含了 s[right] 为 无关字符的情况
		for ; cnt[val] > freq[val]; left++ {
			cnt[s[left]-'a']--
		}
		// 如果 right-left < len(p) 说明没有找到 字母异位词
		if right-left < len(p)-1 { // 应该是 len(p)-1
			continue
		}
		res = append(res, left)
		// left 向右收缩一次
		cnt[s[left]-'a']--
		left++
	}
	return res
}

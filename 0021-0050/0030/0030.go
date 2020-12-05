package main

import "fmt"

func findSubstring1(s string, words []string) []int {
	wsl := len(words) // 字符串长度
	if s == "" || wsl == 0 {
		return nil
	}
	wl := len(words[0]) // 候选字符串长度
	mgl := wl * wsl     // 字符串聚合长度
	sl := len(s)        // 母字符串长度
	if mgl > sl {
		return nil
	}
	// 两个哈希表：一个用于运行中判断，一个备份
	m1, mbak := make(map[string]int, wsl), make(map[string]int, wsl)
	for i := range words {
		mbak[words[i]]++
	}
	ret := make([]int, 0, 4)
	for i := 0; i <= sl-mgl; i++ { // 余留长度不足则停止，是 <= 而不是 <
		for k, v := range mbak { // 恢复次数
			m1[k] = v
		}
		count := 0
		// 分区段判断子串是否存在
		for index := i; count < wsl; index, count = index+wl, count+1 {
			v := s[index : index+wl]
			c := m1[v]
			fmt.Println(i, v, c)
			if c == 0 {
				break
			}
			m1[v] = c - 1
		}
		if count == wsl {
			ret = append(ret, i)
		}
	}
	return ret
}

func findSubstring2(s string, words []string) []int {
	wsl := len(words) // 字符串长度
	if s == "" || wsl == 0 {
		return nil
	}
	wl := len(words[0]) // 候选字符串长度
	mgl := wl * wsl     // 字符串聚合长度
	sl := len(s)        // 母字符串长度
	if mgl > sl {
		return nil
	}
	// 两个哈希表：一个用于运行中判断，一个备份
	m1, mbak := make(map[string]int, wsl), make(map[string]int, wsl)
	for i := range words {
		mbak[words[i]]++
	}
	ret := make([]int, 0, 4)
	for i := 0; i < wl; i++ {
		// 余留长度不足则停止，是 <= 而不是 <
		for start, count := i, 0; start <= sl-mgl; count = 0 {
			for k, v := range mbak { // 恢复次数
				m1[k] = v
			}
			// 分区段判断子串是否存在
			typ := 0 // 找到了
			for index := start + count*wl; count < wsl; index, count = index+wl, count+1 {
				v := s[index : index+wl]
				c, ok := m1[v]
				if !ok {
					typ = 2 // 不存在
					break
				}
				if c == 0 {
					typ = 1 // 存在但已经用完
					break
				}
				m1[v] = c - 1
			}
			if typ == 0 { // 跳 wl 格
				ret = append(ret, start)
				start += wl
			} else if typ == 1 { // 存在已经用完  跳 wl 格
				start += wl
			} else { // 不存在 直接跳过不存在的格子
				start += (count + 1) * wl
			}
		}
	}
	return ret
}

func findSubstring(s string, words []string) []int {
	wsl := len(words) // 字符串长度
	if s == "" || wsl == 0 {
		return nil
	}
	wl := len(words[0]) // 候选字符串长度
	mgl := wl * wsl     // 字符串聚合长度
	sl := len(s)        // 母字符串长度
	if mgl > sl {
		return nil
	}
	// 两个哈希表：一个用于运行中判断，一个备份
	mbak := make(map[string]int, wsl)
	for i := range words {
		mbak[words[i]]++
	}
	ret := make([]int, 0, 4)
	for i := 0; i < wl; i++ {
		// 余留长度不足则停止，是 <= 而不是 <
		for start := i; start <= sl-mgl; {
			m1 := make(map[string]int, wsl)
			// 分区段判断子串是否存在
			index := start + (wsl-1)*wl
			for ; index >= start; index = index - wl {
				v := s[index : index+wl]
				if m1[v]++; m1[v] > mbak[v] {
					break
				}
			}
			if index < start { // 跳 wl 格
				ret = append(ret, start)
				start += wl
			} else { // 当前index 不符合
				start = index + wl
			}
		}
	}
	return ret
}

package main

import "strings"

// 超时版本
func wordBreak_overtime(s string, wordDict []string) bool {
	// 终点，
	if s == "" {
		return true
	}
	for _, v := range wordDict {
		// 若头部匹配，则去掉头部
		if !strings.HasPrefix(s, v) {
			continue
		}
		// 检查字串
		if wordBreak(s[len(v):], wordDict) {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	// buf 记忆化搜索
	return helper(s, 0, wordDict, make([]uint8, len(s)))
}

func helper(s string, low int, wordDict []string, buf []uint8) bool {
	// 完美匹配
	if s[low:] == "" {
		return true
	}
	// 已经计算过则取结果
	if buf[low] != 0 {
		// 1 表示存在，2 表示不存在，0表示未计算
		return buf[low] == 1
	}
	for _, v := range wordDict {
		// 若头部匹配，则去掉头部
		// 一定是 s[low:]
		if !strings.HasPrefix(s[low:], v) {
			continue
		}
		// 检查字串
		if helper(s, low+len(v), wordDict, buf) {
			buf[low] = 1 // 存在
			return true
		}
	}
	buf[low] = 2 // 不存在
	return false
}

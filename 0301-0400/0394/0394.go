package main

import (
	"bytes"
	"strconv"
	"strings"
)

func decodeString1(s string) string {
	length := len(s)
	var res = bytes.NewBuffer(make([]byte, 0, length))
	// 扫描字符串
	for i := 0; i < length; {
		if s[i] < '0' || s[i] > '9' {
			// 普通字符串直接写入到缓存
			res.WriteByte(s[i])
			i++
			continue
		}
		// 扫描到次数，则发现重复模式
		end := i + strings.Index(s[i:], "[") // 次数的终点，需要 + i
		k, _ := strconv.Atoi(s[i:end])       // 次数
		start := end + 1                     // 模式的起点
		c := 1                               // [ 的个数，需要匹配 ]
		// 扫描到 k 所对应的重复模式
		for end = start; end < length; end++ {
			switch s[end] {
			case '[':
				c++
			case ']':
				c--
			}
			if c == 0 { // 模式匹配完成
				break
			}
		}
		// 递归得到内层模式的结果
		// 可优化：如果字符串不存在 [ 则直接返回
		tmp := decodeString(s[start:end])
		res.WriteString(strings.Repeat(tmp, k)) // 重复k次
		i = end + 1                             // 新起点
	}
	return res.String()
}
func decodeString(s string) string {
	// 没有 模式 则返回
	if !strings.Contains(s, "[") {
		return s
	}
	length := len(s)
	s1 := make([]int, 0, length/2)
	s2 := make([]string, 0, length/2)
	k := 1 // k 初始为1
	cur := bytes.NewBuffer(make([]byte, 0, length))
	for i := 0; i < length; i++ {
		v := s[i]
		if v >= '0' && v <= '9' {
			// 扫描到次数，则把 k 和 cur 压入栈
			// 并更新 k 和 cur
			s1 = append(s1, k)
			s2 = append(s2, cur.String())
			cur.Reset()
			end := strings.Index(s[i:], "[") + i // 需要 + i
			k, _ = strconv.Atoi(s[i:end])
			i = end
		} else if v == ']' {
			// 扫描 ] 则 生成重复串
			tmp := strings.Repeat(cur.String(), k)
			cur.Reset()
			// 取出栈顶 k 和 cur
			sl := len(s1) - 1
			k = s1[sl]
			s1 = s1[:sl]
			// 更新 cur，把 生成串追加到 cur后面
			cur.WriteString(s2[sl])
			s2 = s2[:sl]
			cur.WriteString(tmp)
		} else {
			// 普通字符串
			cur.WriteByte(v)
		}
	}
	return cur.String()
}

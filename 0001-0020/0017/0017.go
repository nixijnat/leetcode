package main

import "container/list"

func letterCombinations1(digits string) []string {
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	ret := make([]string, 0, 8)
	for i := range digits {
		dict := m[digits[i]]
		if dict == "" {
			continue
		}
		if i == 0 {
			// 初始化
			for j := range dict {
				ret = append(ret, string(dict[j]))
			}
		} else {
			dictLen := len(dict)
			retLen := len(ret)
			// 先用 ret[:retLen] 为原始值 求 一部分
			for j := 1; j < dictLen; j++ {
				for _, str := range ret[:retLen] {
					ret = append(ret, str+string(dict[j]))
				}
			}
			// 对 ret[:retLen] 修正
			for j, str := range ret[:retLen] {
				ret[j] = str + string(dict[0])
			}
		}
	}
	return ret
}

func letterCombinations(digits string) []string {
	m := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	l := list.New()
	l.PushBack("")
	for i := 0; i < len(digits); i++ {
		tmp := m[digits[i]-'2']
		for j := l.Len(); j > 0; j-- { // 计数的方式确定一轮循环结束
			headPtr := l.Front()
			str := headPtr.Value.(string)
			headPtr.Value = str + string(tmp[0])
			l.MoveToBack(headPtr) // move 避免 生成新的node
			for j := 1; j < len(tmp); j++ {
				l.PushBack(str + string(tmp[j]))
			}
		}
	}
	ret := make([]string, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next() {
		ret = append(ret, i.Value.(string))
	}
	return ret
}

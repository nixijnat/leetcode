package main

import "sort"

func groupAnagrams1(strs []string) [][]string {
	n := len(strs)
	if n == 0 {
		return [][]string{}
	}
	// 哈希归位
	tmp := make(map[string][]string, n/2)
	for _, v := range strs {
		bs := []byte(v)
		// 对字符串排序
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		tmp[string(bs)] = append(tmp[string(bs)], v)
	}
	res := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		res = append(res, v)
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	if n == 0 {
		return [][]string{}
	}
	// 哈希归位
	tmp := make(map[[26]byte][]string, n/2)
	for _, v := range strs {
		cnt := count(v) // 统计字母个数
		tmp[cnt] = append(tmp[cnt], v)
	}
	res := make([][]string, 0, len(tmp))
	for _, v := range tmp {
		res = append(res, v)
	}
	return res
}

func count(str string) (res [26]byte) {
	for _, v := range str {
		res[v-'a']++
	}
	return
}

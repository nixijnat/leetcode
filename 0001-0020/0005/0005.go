package main

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 || l == 1 {
		return s
	}
	find := func(low, high int, ret *string) {
		for ; low >= 0 && high < l; low, high = low-1, high+1 {
			if s[low] != s[high] {
				break
			}
			if len(*ret) < high-low+1 {
				*ret = s[low : high+1]
			}
		}
	}
	ret := s[:1]
	// 若当前找到的长度一半已经比剩下长度更长，则停止
	for cur := 1; cur < l && len(ret)/2 < l-cur; cur++ {
		// 奇数
		find(cur-1, cur+1, &ret)
		find(cur-1, cur, &ret)
	}
	return ret
}

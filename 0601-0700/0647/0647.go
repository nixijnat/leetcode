package main

// 6471
func countSubstrings(s string) int {
	res := 0
	for i := range s {
		res++
		// 奇数中心
		low, high := i-1, i+1
		for low >= 0 && high <= len(s)-1 && s[low] == s[high] {
			res++
			low--
			high++
		}
		// 偶数中心
		if i < len(s)-1 && s[i] == s[i+1] {
			res++
			low, high := i-1, i+2
			for low >= 0 && high <= len(s)-1 && s[low] == s[high] {
				res++
				low--
				high++
			}
		}
	}
	return res
}

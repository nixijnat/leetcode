package q58

// 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	if n <= 0 || len(s) <= n {
		return s
	}
	return s[n:] + s[:n]
}

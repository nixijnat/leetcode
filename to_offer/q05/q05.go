package q05

// 05. 替换空格

func replaceSpace(s string) string {
	n := make([]byte, 0, len(s)*2)
	for i := range s {
		if s[i] == ' ' {
			n = append(n, '%', '2', '0')
		} else {
			n = append(n, s[i])
		}
	}
	return string(n)
}

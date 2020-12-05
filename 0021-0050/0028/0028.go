package main

func strStr(haystack string, needle string) int {
	if needle == "" || haystack == needle {
		return 0
	}
	l1, l2 := len(haystack), len(needle)
	if l2 >= l1 {
		return -1
	}

	for i := 0; i < l1; i++ {
		if l1-i < l2 {
			return -1
		}
		j := 0
		for ; j < l2; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == l2 {
			return i
		}
	}
	return -1
}

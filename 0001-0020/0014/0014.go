package main

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	switch length {
	case 0:
		return ""
	case 1:
		return strs[0]
	}
	getPrefix := func(str1, str2 string) string {
		index := -1
		for i := 0; i < len(str1) && i < len(str2); i++ {
			if str1[i] != str2[i] {
				break
			}
			index = i
		}
		return str1[:index+1]
	}
	ret := getPrefix(strs[0], strs[1])
	for i := 2; i < length; i++ {
		if ret == "" || strs[i] == "" {
			return ""
		}
		ret = getPrefix(ret, strs[i])
	}
	return ret
}

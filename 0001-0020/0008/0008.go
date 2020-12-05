package main

import "strconv"

func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	const maxInt = 1<<31 - 1
	var flag, num int64
	for i := range s {
		switch {
		case s[i] == ' ':
			if flag != 0 {
				return int(num)
			}
		case s[i] >= '0' && s[i] <= '9':
			if flag >= 0 {
				num = num*10 + int64(s[i]-'0')
				flag = 1
			} else {
				num = num*10 - int64(s[i]-'0')
			}
			if num > maxInt {
				return maxInt
			} else if num < -maxInt-1 {
				return -maxInt - 1
			}
		case s[i] == '-':
			if flag != 0 {
				return int(num)
			}
			flag = -1
		case s[i] == '+':
			if flag != 0 {
				return int(num)
			}
			flag = 1
		default:
			return int(num)
		}
	}
	return int(num)
	strconv.Atoi(s string)
}

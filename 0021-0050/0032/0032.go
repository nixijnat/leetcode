package main

func longestValidParentheses(s string) int {
	l := len(s)
	if l <= 1 {
		return 0
	}
	stack := make([]byte, 0, l/2) // 判断括号合法
	record := make([]int, l)      // 记录长度
	max := 0
	for i := 0; i < l; i++ {
		if s[i] == '(' { // 左括号就入栈
			stack = append(stack, '(')
			continue
		}
		sl := len(stack)
		if sl == 0 { // 非法括号 无意义
			continue
		}
		stack = stack[:sl-1] // 出栈
		// 尝试合并包裹的序列
		record[i] = 2 + record[i-1]
		// 尝试合并上一组序列
		if j := i - record[i]; j >= 0 {
			record[i] += record[j]
		}
		if record[i] > max {
			max = record[i]
		}
	}
	return max
}

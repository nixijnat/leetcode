package main

func isValid(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	} else if length&1 == 1 {
		return false
	}
	stack := make([]byte, 0, length)
	for i := range s {
		var expect byte
		switch s[i] {
		case '(', '{', '[':
			stack = append(stack, s[i])
		case ')':
			expect = '('
		case '}', ']':
			expect = s[i] - 2
		default:
			return false
		}
		if expect == 0 {
			continue
		}
		length = len(stack)
		if length == 0 || stack[length-1] != expect {
			return false
		}
		stack = stack[:length-1]
	}
	return len(stack) == 0
}

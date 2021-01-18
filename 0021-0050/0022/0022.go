package main

func generateParenthesis(n int) []string {
	res := make([]string, 0, 1)
	generate("(", n-1, n, &res)
	return res
}

// 生成函数
func generate(prefix string, ln, rn int, res *[]string) {
	// 非法
	if ln < 0 || rn < 0 || ln > rn {
		return
	}
	// 终止：合法
	if ln == 0 && rn == 0 {
		(*res) = append((*res), prefix)
		return
	}
	generate(prefix+"(", ln-1, rn, res)
	generate(prefix+")", ln, rn-1, res)
}

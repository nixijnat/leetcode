package main

// 0856
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	} else if len(matrix[0]) == 0 {
		return 0
	}
	// 当前行的柱子高度
	line := make([]int, len(matrix[0]))
	stk := make([]int, 0, len(matrix[0]))
	res := 0
	for _, v := range matrix {
		// 更新当前行柱子的高度
		for i, v1 := range v {
			if v1 == '0' { // 若为 0 则归零
				line[i] = 0
			} else { // 否则增加高度
				line[i]++
			}
		}
		// 以下与 84 题一样
		stk = stk[:0]
		left := make([]int, len(matrix[0]))
		for j, v1 := range line {
			for ; len(stk) > 0 && line[stk[len(stk)-1]] >= v1; stk = stk[:len(stk)-1] {
			}
			if len(stk) == 0 {
				left[j] = -1
			} else {
				left[j] = stk[len(stk)-1]
			}
			stk = append(stk, j)
		}
		stk = stk[:0]
		right := make([]int, len(matrix[0]))
		for j := len(line) - 1; j >= 0; j-- {
			for ; len(stk) > 0 && line[stk[len(stk)-1]] >= line[j]; stk = stk[:len(stk)-1] {
			}
			if len(stk) == 0 {
				right[j] = len(line)
			} else {
				right[j] = stk[len(stk)-1]
			}
			stk = append(stk, j)
		}
		for j, v1 := range line {
			tmp := (right[j] - left[j] - 1) * v1
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}

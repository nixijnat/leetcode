package main

func maximalSquare1(matrix [][]byte) int {
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
			// 取较短的边
			min := (right[j] - left[j] - 1)
			if min > v1 {
				min = v1
			}
			tmp := min * min
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}

/////
func maximalSquare(matrix [][]byte) int {
	// 不用使用矩阵存储 dp
	line := make([]int, len(matrix[0]))
	res := 0
	for i := range matrix {
		// 左上角的元素
		leftUp := 0
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				line[j] = 0
				continue
			}
			// 存储下一个元素的左上角元素
			nextLeftUp := line[j]
			if j == 0 {
				// 第一列
				line[j] = 1
			} else {
				// line[j] 为 up，line[j-1] 为 left
				line[j] = min(leftUp, min(line[j], line[j-1])) + 1
			}
			// 存储下一个元素的左上角元素
			leftUp = nextLeftUp
			if line[j] > res {
				res = line[j]
			}
		}
	}
	return res * res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

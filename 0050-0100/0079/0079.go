package main

func exist(board [][]byte, word string) bool {
	height := len(board)
	if height == 0 {
		return false
	}
	width := len(board[0])
	if width == 0 {
		return false
	}
	matrix := make([][]bool, len(board))
	for i := range board {
		matrix[i] = make([]bool, len(board[i]))
	}
	var find func(string, int, int) bool
	find = func(w string, i, j int) bool {
		// 边界情况
		if i < 0 || i >= height || j < 0 || j >= width {
			return false
		}
		// 已经走过的点 -> 终止
		// 字符匹配失败 -> 终止
		if matrix[i][j] || w[0] != board[i][j] {
			return false
		}
		// 成功
		if len(w) == 1 {
			return true
		}
		// 标记走过
		matrix[i][j] = true
		// 四向奔走继续匹配
		res := find(w[1:], i-1, j) || find(w[1:], i+1, j) ||
			find(w[1:], i, j-1) || find(w[1:], i, j+1)
		if !res {
			// 匹配失败：去除标记
			matrix[i][j] = false
		}
		return res
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// 探测起始点
			res := find(word, i, j)
			if res {
				return res
			}
		}
	}
	return false
}

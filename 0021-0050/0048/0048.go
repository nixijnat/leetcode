package main

func rotate1(matrix [][]int) {
	// n 是最大索引 不是长度
	n := len(matrix) - 1
	// 只需要遍历一半元素
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i <= n; i++ {
		for j := n / 2; j >= 0; j-- {
			matrix[i][j], matrix[i][n-j] = matrix[i][n-j], matrix[i][j]
		}
	}
}

func rotate(matrix [][]int) {
	// n 是最大索引 不是长度
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = tmp
		}
	}
}

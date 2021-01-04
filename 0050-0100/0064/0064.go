package main

func minPathSum1(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	// 从 0, 0 开始找最小路径
	return find(grid, 0, 0, m, n, 0, int(1<<31-1))
}

// i, j 开始找的点
// m, n 边界
// cur 当前已走的路
// min 截至当前已经探明的最短路径
func find(grid [][]int, i, j, m, n int, cur, min int) int {
	if i == m || j == n {
		return min
	}
	cur += grid[i][j] // 到 i, j 的最短路径
	if cur >= min {   // 已经比min大则停止
		return min
	}
	if i == m-1 && j == n-1 { // 终点
		return cur
	}
	// 向下走
	min = find(grid, i+1, j, m, n, cur, min)
	// 向右走
	return find(grid, i, j+1, m, n, cur, min)
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	res := make([]int, n)
	// 第一行的最小路径和
	res[0] = grid[0][0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] + grid[0][i]
	}
	// 第二行的最小路径和
	for i := 1; i < m; i++ {
		res[0] += grid[i][0]
		for j := 1; j < n; j++ {
			res[j] = grid[i][j] + min(res[j], res[j-1])
		}
	}
	return res[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

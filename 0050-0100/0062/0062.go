package main

func uniquePaths2(m int, n int) int {
	col := make([]int, m)
	for i := range col {
		col[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			col[j] += col[j-1]
		}
	}
	return col[m-1]
}

// 排列组合
func uniquePaths3(m int, n int) int {
	denominator := 1 // 分母
	mumerator := 1   // 分子
	total := m + n - 2
	for i := 1; i <= m-1; i++ {
		mumerator *= total - i + 1 // 大数会溢出
		denominator *= i           // 大数会溢出
	}
	return mumerator / denominator
}

func uniquePaths(m int, n int) int {
	small := m
	if small > n { // 选小的，避免溢出和大量计算
		small = n
	}
	ret := 1
	total := m + n - 2
	for i := 1; i <= small-1; i++ {
		// NOTE: ret *= (total - i + 1) / i 是不行的，导致乘法最后进行
		ret = ret * (total - i + 1) / i
	}
	return ret
}

func combination(n, k int) int {
	small := n - k // 补
	if small > k { // 取小的
		small = k
	}
	res := 1
	for i := 0; i < small; i++ {
		res = res * (n - i) / (i + 1)
	}
	return res
}

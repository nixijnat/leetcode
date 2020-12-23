package main

func numTrees1(n int) int {
	if n <= 2 {
		return n
	}
	res := 0
	for i := 1; i <= n; i++ {
		l, r := 1, 1
		// 左子树的种类数
		if i-1 > 0 {
			l = numTrees(i - 1)
		}
		// 右子树的种类数
		if n-i > 0 {
			r = numTrees(n - i)
		}
		// 乘积
		res += l * r
	}
	return res
}

func numTrees2(n int) int {
	if n <= 2 {
		return n
	}
	// 缓存计算结果
	buf := make([]int, n+1)
	// 特殊的 0 赋值为 1，方便乘法
	buf[0], buf[1], buf[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		// 使用公式计算
		for j := 0; j < i; j++ {
			buf[i] += buf[i-1-j] * buf[j]
		}
	}
	return buf[n]
}

func numTrees(n int) int {
	if n <= 2 {
		return n
	}
	// 缓存计算结果
	buf := make([]int, n+1)
	// 特殊的 0 赋值为 1，方便乘法
	buf[0], buf[1], buf[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		// 对称性加速计算
		half := i >> 1
		for j := 0; j < half; j++ {
			buf[i] += buf[i-1-j] * buf[j] * 2
		}
		// 奇数情况下，需要将中间数的结果附加上
		if i&0x1 != 0 {
			buf[i] += buf[half] * buf[half]
		}
	}
	return buf[n]
}

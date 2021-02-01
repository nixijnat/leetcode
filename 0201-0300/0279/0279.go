package main

import "math"

// 2798
func numSquares(n int) int {
	// 找到最大的平方数
	// limit := findMax(n)
	limit := int(math.Sqrt(float64(n)))
	buf := make([]int, n+1)
	for i := 1; i <= n; i++ {
		buf[i] = i + 1
		for j := 1; j <= limit; j++ {
			tmp := i - j*j
			if tmp >= 0 && buf[tmp] != -1 && buf[tmp] < buf[i] {
				buf[i] = buf[tmp]
			}
		}
		if buf[i] == i+1 {
			buf[i] = -1
		} else {
			buf[i]++
		}
	}
	return buf[n]
}

// 二分法找平方数
func findMax(n int) int {
	low, high := 0, n
	res := 0
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid <= n {
			res = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}
